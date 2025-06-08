package expr

import (
	"errors"
	"fmt"
	"strconv"
	"wind/tokenizer"
)

type Expr struct {
	Left     uint64
	Op       string
	Right    uint64
	IsBinary bool
}

// ParseExpr supports numbers and variables as operands, plus a single operator for binary expr
func ParseExpr(tokens []tokenizer.Token, vars map[string]uint64) (Expr, error) {
	if len(tokens) == 1 {
		// Single token: number or variable
		if tokens[0].Type == tokenizer.TokenNumber {
			val, err := strconv.ParseUint(tokens[0].Value, 10, 64)
			if err != nil {
				return Expr{}, err
			}
			return Expr{Left: val, IsBinary: false}, nil
		} else if tokens[0].Type == tokenizer.TokenIdent {
			val, ok := vars[tokens[0].Value]
			if !ok {
				return Expr{}, fmt.Errorf("undefined variable: %s", tokens[0].Value)
			}
			return Expr{Left: val, IsBinary: false}, nil
		}
	} else if len(tokens) == 3 {
		// Binary expression: operand operator operand
		leftVal, err := operandValue(tokens[0], vars)
		if err != nil {
			return Expr{}, err
		}
		rightVal, err := operandValue(tokens[2], vars)
		if err != nil {
			return Expr{}, err
		}
		op := tokens[1].Value
		return Expr{Left: leftVal, Right: rightVal, Op: op, IsBinary: true}, nil
	}
	return Expr{}, errors.New("unsupported expression format")
}

func operandValue(token tokenizer.Token, vars map[string]uint64) (uint64, error) {
	if token.Type == tokenizer.TokenNumber {
		return strconv.ParseUint(token.Value, 10, 64)
	} else if token.Type == tokenizer.TokenIdent {
		val, ok := vars[token.Value]
		if !ok {
			return 0, fmt.Errorf("undefined variable: %s", token.Value)
		}
		return val, nil
	}
	return 0, errors.New("invalid operand")
}

func GenerateASM(expr Expr) []string {
	if !expr.IsBinary {
		return []string{
			fmt.Sprintf("    mov rax, %d", expr.Left),
		}
	}

	var asmOp string
	switch expr.Op {
	case "+":
		asmOp = "add"
	case "-":
		asmOp = "sub"
	case "*":
		asmOp = "imul"
	default:
		return []string{"; unsupported operator"}
	}

	return []string{
		fmt.Sprintf("    mov rax, %d", expr.Left),
		fmt.Sprintf("    %s rax, %d", asmOp, expr.Right),
	}
}
