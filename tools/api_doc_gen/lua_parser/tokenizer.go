package lua_parser

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenKind int

const (
	tokenIdentifier tokenKind = iota
	tokenKeyword
	tokenNumber
	tokenString
	tokenSymbol
	tokenEOF
)

type Token struct {
	Kind   tokenKind
	Text   string
	Line   int
	Column int
}

func (token Token) IsKeyword(value string) bool {
	return token.Kind == tokenKeyword && token.Text == value
}

func (token Token) IsSymbol(value string) bool {
	return token.Kind == tokenSymbol && token.Text == value
}

var luaKeywords = map[string]bool{
	"and": true, "break": true, "do": true, "else": true, "elseif": true,
	"end": true, "false": true, "for": true, "function": true, "if": true,
	"in": true, "local": true, "nil": true, "not": true, "or": true,
	"repeat": true, "return": true, "then": true, "true": true,
	"until": true, "while": true,
}

type scanner struct {
	source []rune
	index  int
	line   int
	column int
}

func Tokenize(source string) ([]Token, error) {
	state := &scanner{
		source: []rune(source),
		line:   1,
		column: 1,
	}
	tokens := []Token{}
	for {
		token, ok, err := state.nextToken()
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		tokens = append(tokens, token)
		if token.Kind == tokenEOF {
			break
		}
	}
	return tokens, nil
}

func (state *scanner) nextToken() (Token, bool, error) {
	if err := state.skipWhitespaceAndComments(); err != nil {
		return Token{}, false, err
	}
	if state.atEnd() {
		return Token{Kind: tokenEOF, Text: "<eof>", Line: state.line, Column: state.column}, true, nil
	}

	line, column := state.line, state.column
	r := state.peek(0)

	if isIdentifierStart(r) {
		text := state.readWhile(func(value rune) bool {
			return isIdentifierPart(value)
		})
		kind := tokenIdentifier
		if luaKeywords[text] {
			kind = tokenKeyword
		}
		return Token{Kind: kind, Text: text, Line: line, Column: column}, true, nil
	}

	if unicode.IsDigit(r) {
		text := state.readNumber()
		return Token{Kind: tokenNumber, Text: text, Line: line, Column: column}, true, nil
	}

	if r == '"' || r == '\'' {
		text, err := state.readQuotedString(r)
		if err != nil {
			return Token{}, false, err
		}
		return Token{Kind: tokenString, Text: text, Line: line, Column: column}, true, nil
	}

	if r == '[' {
		if longText, ok, err := state.readLongBracketLiteral(); err != nil {
			return Token{}, false, err
		} else if ok {
			return Token{Kind: tokenString, Text: longText, Line: line, Column: column}, true, nil
		}
	}

	for _, candidate := range []string{"...", "..", "==", "~=", "<=", ">="} {
		if state.matchString(candidate) {
			state.advanceN(len(candidate))
			return Token{Kind: tokenSymbol, Text: candidate, Line: line, Column: column}, true, nil
		}
	}

	state.advance()
	return Token{Kind: tokenSymbol, Text: string(r), Line: line, Column: column}, true, nil
}

func (state *scanner) skipWhitespaceAndComments() error {
	for !state.atEnd() {
		if unicode.IsSpace(state.peek(0)) {
			state.advance()
			continue
		}
		if state.matchString("--") {
			state.advanceN(2)
			if _, ok, err := state.readLongBracketLiteral(); err != nil {
				return err
			} else if ok {
				continue
			}
			for !state.atEnd() && state.peek(0) != '\n' {
				state.advance()
			}
			continue
		}
		break
	}
	return nil
}

func (state *scanner) readQuotedString(quote rune) (string, error) {
	var builder strings.Builder
	builder.WriteRune(state.advance())
	for !state.atEnd() {
		current := state.advance()
		builder.WriteRune(current)
		if current == '\\' && !state.atEnd() {
			builder.WriteRune(state.advance())
			continue
		}
		if current == quote {
			return builder.String(), nil
		}
	}
	return "", fmt.Errorf("unterminated string at line %d", state.line)
}

func (state *scanner) readLongBracketLiteral() (string, bool, error) {
	if state.peek(0) != '[' {
		return "", false, nil
	}
	startIndex := state.index
	startLine := state.line
	startColumn := state.column
	equalsCount := 0
	state.advance()
	for !state.atEnd() && state.peek(0) == '=' {
		equalsCount++
		state.advance()
	}
	if state.atEnd() || state.peek(0) != '[' {
		state.index = startIndex
		state.line = startLine
		state.column = startColumn
		return "", false, nil
	}
	state.advance()

	var builder strings.Builder
	builder.WriteString("[")
	builder.WriteString(strings.Repeat("=", equalsCount))
	builder.WriteString("[")
	closing := "]" + strings.Repeat("=", equalsCount) + "]"
	for !state.atEnd() {
		if state.matchString(closing) {
			builder.WriteString(closing)
			state.advanceN(len(closing))
			return builder.String(), true, nil
		}
		builder.WriteRune(state.advance())
	}
	return "", false, fmt.Errorf("unterminated long bracket literal at line %d", startLine)
}

func (state *scanner) readWhile(predicate func(rune) bool) string {
	var builder strings.Builder
	for !state.atEnd() && predicate(state.peek(0)) {
		builder.WriteRune(state.advance())
	}
	return builder.String()
}

func (state *scanner) readNumber() string {
	var builder strings.Builder
	for !state.atEnd() {
		current := state.peek(0)
		if unicode.IsDigit(current) || unicode.IsLetter(current) || current == '.' || current == '_' {
			builder.WriteRune(state.advance())
			continue
		}
		break
	}
	return builder.String()
}

func (state *scanner) matchString(value string) bool {
	if state.index+len([]rune(value)) > len(state.source) {
		return false
	}
	for offset, expected := range []rune(value) {
		if state.peek(offset) != expected {
			return false
		}
	}
	return true
}

func (state *scanner) advanceN(count int) {
	for i := 0; i < count; i++ {
		state.advance()
	}
}

func (state *scanner) advance() rune {
	current := state.source[state.index]
	state.index++
	if current == '\n' {
		state.line++
		state.column = 1
	} else {
		state.column++
	}
	return current
}

func (state *scanner) peek(offset int) rune {
	if state.index+offset >= len(state.source) {
		return 0
	}
	return state.source[state.index+offset]
}

func (state *scanner) atEnd() bool {
	return state.index >= len(state.source)
}

func isIdentifierStart(value rune) bool {
	return value == '_' || unicode.IsLetter(value)
}

func isIdentifierPart(value rune) bool {
	return isIdentifierStart(value) || unicode.IsDigit(value)
}
