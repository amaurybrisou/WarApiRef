package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"roraddons/tools/mcp_server/model"
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
	store     *Store
	validator *Validator
}

func NewApp(docsRoot string) (*App, error) {
	store, err := NewStore(docsRoot)
	if err != nil {
		return nil, err
	}
	return &App{store: store, validator: NewValidator()}, nil
}

func (a *App) ServeStdio(ctx context.Context, input io.Reader, output io.Writer) error {
	r := bufio.NewReader(input)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		payload, err := readFrame(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		resp := a.handlePayload(payload)
		if len(resp.ID) == 0 {
			continue
		}
		if err := writeFrame(output, resp); err != nil {
			return err
		}
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
	if req.JSONRPC == "" {
		req.JSONRPC = "2.0"
	}
	result, rpcErr := a.dispatch(req.Method, req.Params)
	if rpcErr != nil {
		return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Error: map[string]any{"code": -32000, "message": rpcErr.ErrorMessage, "data": rpcErr}}
	}
	return jsonRPCResponse{JSONRPC: "2.0", ID: req.ID, Result: result}
}

func (a *App) dispatch(method string, params json.RawMessage) (interface{}, *model.APIError) {
	switch method {
	case "initialize":
		return map[string]any{
			"protocolVersion": "2024-11-05",
			"serverInfo":      map[string]any{"name": "war-api-mcp", "version": "2.0.0"},
			"capabilities": map[string]any{
				"tools":     map[string]any{"listChanged": false},
				"resources": map[string]any{"subscribe": false, "listChanged": false},
				"prompts":   map[string]any{"listChanged": false},
			},
		}, nil
	case "notifications/initialized":
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
	default:
		return nil, &model.APIError{ErrorCode: "method_not_found", ErrorMessage: "method not found"}
	}
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
