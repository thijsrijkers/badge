package tokenizer

import (
	"errors"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenLet TokenType = iota
	TokenIdent
	TokenEqual
	TokenNumber
	TokenOperator
	TokenUnknown
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(line string) ([]Token, error) {
	var tokens []Token
	words := strings.Fields(line)
	for _, word := range words {
		switch {
		case word == "let":
			tokens = append(tokens, Token{Type: TokenLet, Value: word})
		case word == "=":
			tokens = append(tokens, Token{Type: TokenEqual, Value: word})
		case isNumber(word):
			tokens = append(tokens, Token{Type: TokenNumber, Value: word})
		case isOperator(word):
			tokens = append(tokens, Token{Type: TokenOperator, Value: word})
		case isIdent(word):
			tokens = append(tokens, Token{Type: TokenIdent, Value: word})
		default:
			return nil, errors.New("unknown token: " + word)
		}
	}
	return tokens, nil
}

func isNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*"
}

func isIdent(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, r := range s {
		if i == 0 && !unicode.IsLetter(r) && r != '_' {
			return false
		}
		if i > 0 && !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}
