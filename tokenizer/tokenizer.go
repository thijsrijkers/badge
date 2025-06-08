package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	TokenLet TokenType = iota
	TokenIdent
	TokenNumber
	TokenEqual
	TokenPlus
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	runes := []rune(input)
	i := 0

	for i < len(runes) {
		ch := runes[i]

		switch {
		case unicode.IsSpace(ch):
			i++

		case unicode.IsLetter(ch):
			start := i
			for i < len(runes) && (unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i])) {
				i++
			}
			word := string(runes[start:i])
			if word == "let" {
				tokens = append(tokens, Token{Type: TokenLet, Value: word})
			} else {
				tokens = append(tokens, Token{Type: TokenIdent, Value: word})
			}

		case unicode.IsDigit(ch):
			start := i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{Type: TokenNumber, Value: string(runes[start:i])})

		case ch == '=':
			tokens = append(tokens, Token{Type: TokenEqual, Value: string(ch)})
			i++

		case ch == '+':
			tokens = append(tokens, Token{Type: TokenPlus, Value: string(ch)})
			i++

		default:
			return nil, fmt.Errorf("unexpected character: '%c'", ch)
		}
	}

	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens, nil
}
