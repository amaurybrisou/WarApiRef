package deep_analysis

import (
	"regexp"
	"strings"
)

// ArgumentInference analyzes function arguments and parameters
type ArgumentInference struct {
	FunctionArgs map[string]*ArgumentAnalysis
	CallPatterns map[string][]CallPattern // function -> patterns of how it's called
}

// ArgumentAnalysis holds inferred information about a function argument
type ArgumentAnalysis struct {
	FunctionPath   string
	ParameterName  string
	ParameterIndex int
	Usage          []string     // patterns of how this parameter is used
	TypeSignals    []TypeSignal // signals about its type
	IsOptional     bool
	IsVariadic     bool
	DefaultValue   string
	Evidence       []ArgumentEvidence
}

// TypeSignal represents an observation about a type
type TypeSignal struct {
	Signal     string // "indexed", "called", "compared_to_number", etc
	Confidence int    // 0-100
	Details    string
}

// ArgumentEvidence documents inferred argument properties
type ArgumentEvidence struct {
	Type       string // "call_site", "usage", "comment", "pattern"
	Details    string
	Confidence int // 0-100
}

// CallPattern represents a specific way a function was called
type CallPattern struct {
	Arguments     []string // argument values passed
	ArgumentTypes []string // inferred types of arguments
	Context       string   // what called it
	Frequency     int      // how many times this pattern seen
	Confidence    int      // 0-100 confidence this is typical
}

// NewArgumentInference creates argument inference engine
func NewArgumentInference() *ArgumentInference {
	return &ArgumentInference{
		FunctionArgs: make(map[string]*ArgumentAnalysis),
		CallPatterns: make(map[string][]CallPattern),
	}
}

// AnalyzeParameters extracts parameter information from function signature
func (ai *ArgumentInference) AnalyzeParameters(functionPath, source string) []string {
	// Extract function signature
	sigRegex := regexp.MustCompile(`(?:local\s+)?function\s+\w+\s*\(([^)]*)\)`)
	match := sigRegex.FindStringSubmatch(source)
	if len(match) < 2 {
		return []string{}
	}

	paramStr := match[1]
	if paramStr == "" {
		return []string{}
	}

	// Handle variadic (...) parameters
	if strings.Contains(paramStr, "...") {
		// Replace ... with a placeholder
		analysis := &ArgumentAnalysis{
			FunctionPath:   functionPath,
			ParameterName:  "...",
			ParameterIndex: 0,
			IsVariadic:     true,
		}
		ai.FunctionArgs[functionPath+".variadic"] = analysis
		return []string{"..."}
	}

	// Split by commas, handling nested parentheses
	params := strings.Split(paramStr, ",")
	paramNames := []string{}
	for i, p := range params {
		p = strings.TrimSpace(p)
		if p != "" {
			// Extract parameter name (might have default value)
			nameMatch := regexp.MustCompile(`(\w+)`).FindString(p)
			if nameMatch != "" {
				paramNames = append(paramNames, nameMatch)

				// Check for default value
				var defaultVal string
				if strings.Contains(p, "=") {
					parts := strings.Split(p, "=")
					if len(parts) > 1 {
						defaultVal = strings.TrimSpace(parts[1])
					}
				}

				analysis := &ArgumentAnalysis{
					FunctionPath:   functionPath,
					ParameterName:  nameMatch,
					ParameterIndex: i,
					DefaultValue:   defaultVal,
					IsOptional:     defaultVal != "",
				}
				ai.FunctionArgs[functionPath+"@"+nameMatch] = analysis
			}
		}
	}

	return paramNames
}

// AnalyzeParameterUsage examines how parameters are used in function body
func (ai *ArgumentInference) AnalyzeParameterUsage(functionPath, source string, paramNames []string) {
	lines := strings.Split(source, "\n")

	for _, param := range paramNames {
		key := functionPath + "@" + param
		if _, ok := ai.FunctionArgs[key]; !ok {
			continue
		}

		for _, line := range lines {
			if !strings.Contains(line, param) {
				continue
			}

			// Pattern: param.field (accessed like object)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\.(\w+)`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "field_access")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "has_fields",
					Confidence: 75,
					Details:    "Accessed like table/object with fields",
				})
			}

			// Pattern: param[...]  (indexed)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\[`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "indexing")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "indexable",
					Confidence: 80,
					Details:    "Used with bracket indexing []",
				})
			}

			// Pattern: param(...) (called as function)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\s*\(`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "called")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "callable",
					Confidence: 85,
					Details:    "Invoked as function with ()",
				})
			}

			// Pattern: param + num (math operation)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\s*[\+\-\*/%]`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "arithmetic")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "numeric",
					Confidence: 85,
					Details:    "Used in arithmetic operations",
				})
			}

			// Pattern: param:method(...) (method call)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `:`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "method_call")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "has_methods",
					Confidence: 80,
					Details:    "Methods called with : syntax (object pattern)",
				})
			}

			// Pattern: if param then (boolean condition)
			if regexp.MustCompile(`if\s+` + regexp.QuoteMeta(param) + `\s+then`).MatchString(line) {
				ai.FunctionArgs[key].Usage = append(ai.FunctionArgs[key].Usage, "boolean_condition")
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "truthy_check",
					Confidence: 70,
					Details:    "Checked as boolean in if statement",
				})
			}

			// Pattern: param == string/number (type comparison)
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\s*==\s*"[^"]*"`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "compared_to_string",
					Confidence: 75,
					Details:    "Compared for equality with string",
				})
			}
			if regexp.MustCompile(regexp.QuoteMeta(param) + `\s*==\s*\d+`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "compared_to_number",
					Confidence: 80,
					Details:    "Compared for equality with number",
				})
			}

			// Pattern: for k, v in ipairs(param) (array iteration)
			if regexp.MustCompile(`ipairs\s*\(\s*` + regexp.QuoteMeta(param) + `\s*\)`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "array_like",
					Confidence: 85,
					Details:    "Iterated with ipairs() (array pattern)",
				})
			}

			// Pattern: for k, v in pairs(param) (table iteration)
			if regexp.MustCompile(`pairs\s*\(\s*` + regexp.QuoteMeta(param) + `\s*\)`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "table_like",
					Confidence: 85,
					Details:    "Iterated with pairs() (table/dict pattern)",
				})
			}

			// Pattern: #param (length operator, array/table)
			if regexp.MustCompile(`#` + regexp.QuoteMeta(param)).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "has_length",
					Confidence: 90,
					Details:    "Length operator # used (table/string/array)",
				})
			}

			// Pattern: tostring(param) (can convert to string)
			if regexp.MustCompile(`tostring\s*\(\s*` + regexp.QuoteMeta(param) + `\s*\)`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "convertible_to_string",
					Confidence: 60,
					Details:    "Converted with tostring()",
				})
			}

			// Pattern: tonumber(param)
			if regexp.MustCompile(`tonumber\s*\(\s*` + regexp.QuoteMeta(param) + `\s*\)`).MatchString(line) {
				ai.FunctionArgs[key].TypeSignals = append(ai.FunctionArgs[key].TypeSignals, TypeSignal{
					Signal:     "string_to_number",
					Confidence: 75,
					Details:    "Parsed with tonumber() (likely string)",
				})
			}
		}
	}
}

// AnalyzeCallSites examines how a function is called to understand argument patterns
func (ai *ArgumentInference) AnalyzeCallSites(functionPath, source string, paramNames []string) {
	lines := strings.Split(source, "\n")

	// Extract function name from path
	parts := strings.Split(functionPath, ".")
	funcName := parts[len(parts)-1]

	callCount := 0
	patterns := []CallPattern{}

	for _, line := range lines {
		// Find calls to this function
		callRegex := regexp.MustCompile(regexp.QuoteMeta(funcName) + `\s*\(([^)]*)\)`)
		matches := callRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) < 2 {
				continue
			}

			callCount++
			argsStr := match[1]
			if argsStr == "" {
				continue
			}

			// Parse arguments
			args := strings.Split(argsStr, ",")
			pattern := CallPattern{
				Arguments:     []string{},
				ArgumentTypes: []string{},
				Frequency:     1,
				Confidence:    75,
			}

			for i, arg := range args {
				arg = strings.TrimSpace(arg)
				pattern.Arguments = append(pattern.Arguments, arg)

				// Infer type of this argument
				argType := inferArgumentType(arg)
				pattern.ArgumentTypes = append(pattern.ArgumentTypes, argType)

				// Record this type for the corresponding parameter
				if i < len(paramNames) {
					key := functionPath + "@" + paramNames[i]
					if analysis, ok := ai.FunctionArgs[key]; ok {
						analysis.Evidence = append(analysis.Evidence, ArgumentEvidence{
							Type:       "call_site",
							Details:    "Called with: " + arg + " (type: " + argType + ")",
							Confidence: 70,
						})
					}
				}
			}

			patterns = append(patterns, pattern)
		}
	}

	if len(patterns) > 0 {
		ai.CallPatterns[functionPath] = patterns
	}
}

// InferArgumentRole determines what role an argument plays in the function
func (ai *ArgumentInference) InferArgumentRole(functionPath, paramName string) string {
	key := functionPath + "@" + paramName
	analysis, ok := ai.FunctionArgs[key]
	if !ok {
		return "unknown"
	}

	// Analyze signal patterns to determine role
	signals := make(map[string]int)
	for _, sig := range analysis.TypeSignals {
		signals[sig.Signal]++
	}

	// Check for specific patterns
	if _, hasIndexing := signals["indexable"]; hasIndexing {
		return "collection"
	}
	if _, hasFields := signals["has_fields"]; hasFields {
		return "object"
	}
	if _, isCallable := signals["callable"]; isCallable {
		return "callback"
	}
	if _, isNumeric := signals["numeric"]; isNumeric {
		return "number"
	}
	if _, isTruthy := signals["truthy_check"]; isTruthy {
		return "flag"
	}
	if _, hasLength := signals["has_length"]; hasLength {
		return "collection"
	}

	// If mostly used in array/table patterns
	if (signals["array_like"] + signals["table_like"]) > 1 {
		return "collection"
	}

	return "value"
}

// inferArgumentType determines the likely type of an argument value
func inferArgumentType(argValue string) string {
	argValue = strings.TrimSpace(argValue)

	// Nil
	if argValue == "nil" {
		return "nil"
	}

	// Boolean literals
	if argValue == "true" || argValue == "false" {
		return "boolean"
	}

	// String literals
	if (strings.HasPrefix(argValue, "\"") && strings.HasSuffix(argValue, "\"")) ||
		(strings.HasPrefix(argValue, "'") && strings.HasSuffix(argValue, "'")) {
		return "string"
	}

	// Number literals
	if regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(argValue) {
		return "number"
	}

	// Table literals
	if strings.HasPrefix(argValue, "{") && strings.HasSuffix(argValue, "}") {
		return "table"
	}

	// Function calls (result of function call)
	if strings.Contains(argValue, "(") && strings.Contains(argValue, ")") {
		return "function_result"
	}

	// Variable reference
	if regexp.MustCompile(`^\w+$`).MatchString(argValue) {
		return "variable"
	}

	// Field access (table.field)
	if regexp.MustCompile(`\w+\.\w+`).MatchString(argValue) {
		return "field_value"
	}

	// Unknown
	return "unknown"
}
