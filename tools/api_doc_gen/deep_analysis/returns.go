package deep_analysis

import (
	"regexp"
	"strings"
)

// ReturnInference analyzes functions to infer return types
type ReturnInference struct {
	FunctionReturns map[string]*ReturnAnalysis
	AssignmentTypes map[string][]string // variable -> types assigned to it
	CallReturnTypes map[string][]string // function -> inferred return types
}

// ReturnAnalysis holds inferred information about return types
type ReturnAnalysis struct {
	FunctionPath       string
	ExplicitReturns    []string // explicit return types found in comments or annotations
	InferredTypes      []string // types inferred from return statements
	LiteralReturns     []string // literal values returned
	AssignmentContexts []string // what variables these return values are assigned to
	ComparisonContexts []string // what they're compared against
	Evidence           []InferenceEvidence
}

// InferenceEvidence documents how a type was inferred
type InferenceEvidence struct {
	Type       string // "literal", "assignment_target", "comparison", "call_result", "table_expected"
	Details    string
	Confidence int // 0-100
	LineNum    int
}

// NewReturnInference creates return type inference engine
func NewReturnInference() *ReturnInference {
	return &ReturnInference{
		FunctionReturns: make(map[string]*ReturnAnalysis),
		AssignmentTypes: make(map[string][]string),
		CallReturnTypes: make(map[string][]string),
	}
}

// AnalyzeReturns examines function source for return type signals
func (ri *ReturnInference) AnalyzeReturns(functionPath, source string) *ReturnAnalysis {
	ra := &ReturnAnalysis{
		FunctionPath:       functionPath,
		ExplicitReturns:    []string{},
		InferredTypes:      []string{},
		LiteralReturns:     []string{},
		AssignmentContexts: []string{},
		ComparisonContexts: []string{},
		Evidence:           []InferenceEvidence{},
	}

	lines := strings.Split(source, "\n")
	for i, line := range lines {
		lineNum := i + 1

		// Extract explicit type annotations (e.g., --[[ @returns boolean ]])
		if strings.Contains(line, "@returns") || strings.Contains(line, "@return") {
			typeMatch := regexp.MustCompile(`@returns?\s+(\w+)`).FindStringSubmatch(line)
			if len(typeMatch) > 1 {
				ra.ExplicitReturns = append(ra.ExplicitReturns, typeMatch[1])
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "explicit_annotation",
					Details:    "Found in function comment",
					Confidence: 100,
					LineNum:    lineNum,
				})
			}
		}

		// Extract return statements and their literal values
		returnMatch := regexp.MustCompile(`return\s+(.+?)(?:\s*--|$)`).FindStringSubmatch(line)
		if len(returnMatch) > 1 {
			returnValue := strings.TrimSpace(returnMatch[1])
			ra.LiteralReturns = append(ra.LiteralReturns, returnValue)

			// Infer type from literal
			inferredType := inferTypeFromLiteral(returnValue)
			if inferredType != "" {
				ra.InferredTypes = append(ra.InferredTypes, inferredType)
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "literal",
					Details:    "Literal value: " + returnValue,
					Confidence: 75,
					LineNum:    lineNum,
				})
			}

			// Check if it's a function call result
			if strings.Contains(returnValue, "(") && strings.Contains(returnValue, ")") {
				callMatch := regexp.MustCompile(`(\w+(?:\.\w+)*)\(`).FindStringSubmatch(returnValue)
				if len(callMatch) > 1 {
					ra.Evidence = append(ra.Evidence, InferenceEvidence{
						Type:       "call_result",
						Details:    "Returns result of: " + callMatch[1],
						Confidence: 50,
						LineNum:    lineNum,
					})
				}
			}
		}
	}

	ri.FunctionReturns[functionPath] = ra
	return ra
}

// AnalyzeCallSites examines where a function is called to infer return types
func (ri *ReturnInference) AnalyzeCallSites(functionPath, source string) {
	lines := strings.Split(source, "\n")
	for i, line := range lines {
		lineNum := i + 1

		// Pattern: variable = function()
		assignMatch := regexp.MustCompile(`(\w+)\s*=\s*(\w+(?:\.\w+)*)\(`).FindStringSubmatch(line)
		if len(assignMatch) > 2 {
			varName := assignMatch[1]
			calledFunc := assignMatch[2]
			if strings.HasSuffix(functionPath, calledFunc) {
				// This is a call site of our function
				ri.AssignmentTypes[varName] = append(ri.AssignmentTypes[varName], "unknown")

				// Infer type from what the variable is used for
				inferredType := ri.inferTypeFromUsage(varName, source, lineNum)
				if inferredType != "" {
					if ra, ok := ri.FunctionReturns[functionPath]; ok {
						ra.AssignmentContexts = append(ra.AssignmentContexts, varName+" (type: "+inferredType+")")
						ra.InferredTypes = appendUnique(ra.InferredTypes, inferredType)
						ra.Evidence = append(ra.Evidence, InferenceEvidence{
							Type:       "assignment_target",
							Details:    "Assigned to " + varName + " used as " + inferredType,
							Confidence: 60,
							LineNum:    lineNum,
						})
					}
				}
			}
		}

		// Pattern: if variable then...
		ifMatch := regexp.MustCompile(`if\s+(\w+)\s+then`).FindStringSubmatch(line)
		if len(ifMatch) > 1 {
			varName := ifMatch[1]
			if ra, ok := ri.FunctionReturns[functionPath]; ok {
				ra.ComparisonContexts = append(ra.ComparisonContexts, "boolean_check: "+varName)
				ra.InferredTypes = appendUnique(ra.InferredTypes, "boolean")
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "comparison",
					Details:    "Used in boolean context (if statement)",
					Confidence: 80,
					LineNum:    lineNum,
				})
			}
		}

		// Pattern: for k, v in ipairs(variable)
		iterMatch := regexp.MustCompile(`for\s+\w+(?:,\s*\w+)?\s+in\s+ipairs\((\w+)\)`).FindStringSubmatch(line)
		if len(iterMatch) > 1 {
			_ = iterMatch[1] // varName being iterated
			if ra, ok := ri.FunctionReturns[functionPath]; ok {
				ra.AssignmentContexts = append(ra.AssignmentContexts, "iterated with ipairs (array)")
				ra.InferredTypes = appendUnique(ra.InferredTypes, "array")
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "comparison",
					Details:    "Iterated with ipairs (array/table pattern)",
					Confidence: 85,
					LineNum:    lineNum,
				})
			}
		}

		// Pattern: for k, v in pairs(variable)
		pairsMatch := regexp.MustCompile(`for\s+\w+(?:,\s*\w+)?\s+in\s+pairs\((\w+)\)`).FindStringSubmatch(line)
		if len(pairsMatch) > 1 {
			_ = pairsMatch[1] // varName being iterated
			if ra, ok := ri.FunctionReturns[functionPath]; ok {
				ra.AssignmentContexts = append(ra.AssignmentContexts, "iterated with pairs (table)")
				ra.InferredTypes = appendUnique(ra.InferredTypes, "table")
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "comparison",
					Details:    "Iterated with pairs (table/dictionary pattern)",
					Confidence: 85,
					LineNum:    lineNum,
				})
			}
		}
	}
}

// AnalyzeFieldAccess examines how returned values are used to infer their structure
func (ri *ReturnInference) AnalyzeFieldAccess(functionPath, source string) {
	if ra, ok := ri.FunctionReturns[functionPath]; ok {
		// Look for patterns like: result.field, result[key], result:method()
		fieldAccessRegex := regexp.MustCompile(`(\w+)\.(\w+)`)
		matches := fieldAccessRegex.FindAllStringSubmatch(source, -1)
		for _, match := range matches {
			if len(match) > 2 {
				ra.AssignmentContexts = append(ra.AssignmentContexts, "field_accessed: "+match[2])
				if !strings.Contains(strings.Join(ra.InferredTypes, ","), "table") {
					ra.InferredTypes = appendUnique(ra.InferredTypes, "table_with_field_"+match[2])
				}
				ra.Evidence = append(ra.Evidence, InferenceEvidence{
					Type:       "table_expected",
					Details:    "Field ." + match[2] + " accessed on return value",
					Confidence: 70,
					LineNum:    0,
				})
			}
		}
	}
}

// inferTypeFromLiteral determines type from literal value
func inferTypeFromLiteral(literal string) string {
	literal = strings.TrimSpace(literal)

	// Boolean literals
	if literal == "true" || literal == "false" {
		return "boolean"
	}

	// Number literals
	if regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(literal) {
		if strings.Contains(literal, ".") {
			return "number"
		}
		return "integer"
	}

	// String literals
	if (strings.HasPrefix(literal, "\"") && strings.HasSuffix(literal, "\"")) ||
		(strings.HasPrefix(literal, "'") && strings.HasSuffix(literal, "'")) {
		return "string"
	}

	// Nil
	if literal == "nil" {
		return "nil"
	}

	// Table literals
	if strings.HasPrefix(literal, "{") && strings.HasSuffix(literal, "}") {
		return "table"
	}

	// Array-like tables
	if regexp.MustCompile(`\{\s*\d+\s*,`).MatchString(literal) {
		return "array"
	}

	return ""
}

// inferTypeFromUsage analyzes how a variable is used to determine its type
func (ri *ReturnInference) inferTypeFromUsage(varName, source string, startLine int) string {
	pattern := regexp.MustCompile(`\b` + varName + `\b`)
	matches := pattern.FindAllStringIndex(source, -1)
	if len(matches) == 0 {
		return ""
	}

	// If compared with true/false, it's boolean
	if regexp.MustCompile(varName + `\s*==\s*(true|false)`).MatchString(source) {
		return "boolean"
	}

	// If used in arithmetic, it's number
	if regexp.MustCompile(varName + `\s*[\+\-\*/]`).MatchString(source) {
		return "number"
	}

	// If used with string operations
	if regexp.MustCompile(varName + `\.upper\(\)|` + varName + `\.lower\(\)|` + varName + `\.sub\(`).MatchString(source) {
		return "string"
	}

	// If indexed like array/table
	if regexp.MustCompile(varName + `\[\w+\]`).MatchString(source) {
		return "table"
	}

	return ""
}

// BestGuess returns the most likely return type for a function
func (ra *ReturnAnalysis) BestGuess() (string, int) {
	if len(ra.ExplicitReturns) > 0 {
		return ra.ExplicitReturns[0], 95
	}

	// Count occurrences of each inferred type
	typeCounts := make(map[string]int)
	for _, t := range ra.InferredTypes {
		typeCounts[t]++
	}

	if len(typeCounts) == 0 {
		return "unknown", 0
	}

	// Find most common type
	bestType := ""
	bestCount := 0
	for t, count := range typeCounts {
		if count > bestCount {
			bestCount = count
			bestType = t
		}
	}

	// Confidence based on how many pieces of evidence
	confidence := (bestCount * 25) // up to 75% for multiple matches
	if confidence > 75 {
		confidence = 75
	}

	return bestType, confidence
}

// appendUnique appends value if not already in slice
func appendUnique(slice []string, value string) []string {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}
