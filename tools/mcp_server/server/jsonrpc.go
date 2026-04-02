package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

type jsonRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type jsonRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  interface{}     `json:"result,omitempty"`
	Error   interface{}     `json:"error,omitempty"`
}

type App struct {
	docsRoot    string
	feedingRoot string
	queueMu     sync.Mutex   // protects queue read-modify-write operations
	store       *Store
	storeErr    error
	storeOnce   sync.Once
	storeDone   chan struct{}
	storeMu     sync.RWMutex
	validator   *Validator
}

func NewApp(docsRoot string) (*App, error) {
	app := &App{
		docsRoot:    docsRoot,
		feedingRoot: inferFeedingRoot(docsRoot),
		validator:   NewValidator(),
		storeDone:   make(chan struct{}),
	}
	app.ensureStoreLoaded()
	return app, nil
}

func (a *App) SetFeedingRoot(path string) {
	if strings.TrimSpace(path) == "" {
		return
	}
	a.feedingRoot = path
}

func (a *App) ServeStdio(ctx context.Context, input io.Reader, output io.Writer) error {
	r := bufio.NewReader(input)
	w := bufio.NewWriter(output)
	mode := "unknown"
	log.Printf("mcp stdio server started")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		payload, currentMode, err := readStdioPayload(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if mode == "unknown" {
			mode = currentMode
			log.Printf("mcp stdio detected mode=%s", mode)
		}
		resp := a.handlePayload(payload)
		if len(resp.ID) == 0 {
			continue
		}
		if mode == "framed" {
			if err := writeFrame(w, resp); err != nil {
				return err
			}
		} else {
			if err := writeLineJSON(w, resp); err != nil {
				return err
			}
		}
		if err := w.Flush(); err != nil {
			return err
		}
		log.Printf("mcp stdio response sent")
	}
}

func (a *App) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 10<<20))
	if err != nil {
		http.Error(w, "read request", http.StatusBadRequest)
		return
	}
	resp := a.handlePayload(body)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (a *App) handlePayload(payload []byte) jsonRPCResponse {
	var req jsonRPCRequest
	if err := json.Unmarshal(payload, &req); err != nil {
		return jsonRPCResponse{JSONRPC: "2.0", Error: map[string]any{"code": -32700, "message": "parse error"}}
	}
	log.Printf("mcp request method=%s", req.Method)
	if req.JSONRPC == "" {
		req.JSONRPC = "2.0"
	}
	result, rpcErr := a.dispatch(req.Method, req.Params)
	if rpcErr != nil {
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: map[string]any{"code": rpcCodeFromError(rpcErr), "message": rpcErr.ErrorMessage, "data": rpcErr}}
	}
	return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Result: result}
}

func (a *App) dispatch(method string, params json.RawMessage) (interface{}, *model.APIError) {
	switch method {
	case "initialize":
		return map[string]any{
			"protocolVersion": "2025-03-26",
			"serverInfo":      map[string]any{"name": "war-api-mcp-server", "version": "1.0.0"},
			"capabilities": map[string]any{
				"tools":     map[string]any{},
				"resources": map[string]any{},
			},
		}, nil
	case "initialized", "notifications/initialized":
		return map[string]any{}, nil
	case "tools/list":
		return map[string]any{"tools": ToolRegistry()}, nil
	case "tools/call":
		var req struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid tools/call payload"}
		}
		out, apiErr := a.HandleToolCall(req.Name, req.Arguments)
		if apiErr != nil {
			return map[string]any{"content": []map[string]string{{"type": "text", "text": toJSON(apiErr)}}, "isError": true}, nil
		}
		return map[string]any{"content": []map[string]string{{"type": "text", "text": toJSON(out)}}}, nil
	case "resources/list":
		return map[string]any{"resources": ResourceRegistry()}, nil
	case "resources/read":
		var req struct {
			URI string `json:"uri"`
		}
		if err := json.Unmarshal(params, &req); err != nil || strings.TrimSpace(req.URI) == "" {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "uri is required"}
		}
		res, apiErr := a.HandleResourceRead(req.URI)
		if apiErr != nil {
			return nil, apiErr
		}
		return map[string]any{"contents": []map[string]any{{"uri": req.URI, "mimeType": "application/json", "text": toJSON(res)}}}, nil
	case "prompts/list":
		return map[string]any{"prompts": []map[string]string{{"name": "warapi_symbol_research", "description": "Canonical symbol research workflow"}}}, nil
	case "prompts/get":
		return map[string]any{"description": "WAR symbol research", "messages": []map[string]any{{"role": "user", "content": map[string]string{"type": "text", "text": "Use lookup_symbol, explain_symbol_usage, explain_confidence, and get_related_symbols in that order."}}}}, nil
	case "feeding/ingest":
		var req schema.IngestObservationRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/ingest payload"}
		}
		if len(req.Observation) == 0 {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "observation is required"}
		}
		return a.ingestObservation(req), nil
	case "feeding/ingest_batch":
		var req schema.IngestObservationBatchRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/ingest_batch payload"}
		}
		if req.Limit < 0 {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "limit must be >= 0"}
		}
		return a.ingestObservationBatch(req), nil
	case "feeding/list_pending":
		var req schema.ListPendingObservationsRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/list_pending payload"}
		}
		return a.listPendingObservations(req), nil
	case "feeding/review":
		var req schema.ReviewObservationRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/review payload"}
		}
		if strings.TrimSpace(req.ObservationID) == "" {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "observation_id is required"}
		}
		if req.Verdict != "accept" && req.Verdict != "reject" {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "verdict must be 'accept' or 'reject'"}
		}
		return a.reviewObservation(req), nil
	case "feeding/promote":
		var req schema.PromoteObservationRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/promote payload"}
		}
		if strings.TrimSpace(req.ObservationID) == "" {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "observation_id is required"}
		}
		return a.promoteObservation(req), nil
	case "feeding/list_rejected":
		var req schema.ListRejectedObservationsRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/list_rejected payload"}
		}
		return a.listRejectedObservations(req), nil
	case "feeding/regenerate":
		var req schema.RegenerateRequest
		if err := json.Unmarshal(params, &req); err != nil {
			return nil, &model.APIError{ErrorCode: "invalid_params", ErrorMessage: "invalid feeding/regenerate payload"}
		}
		return a.regenerateFromPromotedKnowledge(req), nil
	default:
		return nil, &model.APIError{ErrorCode: "method_not_found", ErrorMessage: "method not found"}
	}
}

func (a *App) ensureStoreLoaded() {
	a.storeOnce.Do(func() {
		go func() {
			store, err := NewStore(a.docsRoot)
			a.storeMu.Lock()
			a.store = store
			a.storeErr = err
			a.storeMu.Unlock()
			close(a.storeDone)
			if err != nil {
				log.Printf("mcp store load failed: %v", err)
				return
			}
			log.Printf("mcp store loaded")
		}()
	})
}

func (a *App) storeReady() (*Store, *model.APIError) {
	a.ensureStoreLoaded()
	<-a.storeDone

	a.storeMu.RLock()
	defer a.storeMu.RUnlock()
	if a.storeErr != nil {
		return nil, &model.APIError{ErrorCode: "server_unavailable", ErrorMessage: "failed to load documentation index", Details: map[string]any{"cause": a.storeErr.Error()}}
	}
	if a.store == nil {
		return nil, &model.APIError{ErrorCode: "server_unavailable", ErrorMessage: "documentation index unavailable"}
	}
	return a.store, nil
}

func readStdioPayload(r *bufio.Reader) ([]byte, string, error) {
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && strings.TrimSpace(line) != "" {
				return []byte(strings.TrimSpace(line)), "line", nil
			}
			return nil, "", err
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		if strings.HasPrefix(strings.ToLower(trimmed), "content-length:") {
			payload, ferr := readFrameFromFirstHeader(trimmed, r)
			return payload, "framed", ferr
		}

		return []byte(trimmed), "line", nil
	}
}

func readFrameFromFirstHeader(firstHeader string, r *bufio.Reader) ([]byte, error) {
	length := 0
	headers := []string{firstHeader}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		trimmed := strings.TrimRight(line, "\r\n")
		if trimmed == "" {
			break
		}
		headers = append(headers, trimmed)
	}
	for _, h := range headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 && strings.EqualFold(strings.TrimSpace(parts[0]), "content-length") {
			n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, err
			}
			length = n
		}
	}
	if length <= 0 {
		return nil, fmt.Errorf("missing content-length")
	}
	payload := make([]byte, length)
	_, err := io.ReadFull(r, payload)
	return payload, err
}

func readFrame(r *bufio.Reader) ([]byte, error) {
	length := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 && strings.EqualFold(strings.TrimSpace(parts[0]), "content-length") {
			n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, err
			}
			length = n
		}
	}
	if length <= 0 {
		return nil, fmt.Errorf("missing content-length")
	}
	payload := make([]byte, length)
	_, err := io.ReadFull(r, payload)
	return payload, err
}

func writeLineJSON(w io.Writer, resp jsonRPCResponse) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(append(b, '\n'))
	return err
}

func rpcCodeFromError(apiErr *model.APIError) int {
	if apiErr == nil {
		return -32000
	}
	switch apiErr.ErrorCode {
	case "method_not_found":
		return -32601
	case "invalid_params":
		return -32602
	default:
		return -32000
	}
}

func init() {
	log.SetOutput(os.Stderr)
}

func writeFrame(w io.Writer, resp jsonRPCResponse) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	var frame bytes.Buffer
	frame.WriteString("Content-Length: ")
	frame.WriteString(strconv.Itoa(len(b)))
	frame.WriteString("\r\n\r\n")
	frame.Write(b)
	_, err = w.Write(frame.Bytes())
	return err
}

func toJSON(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
