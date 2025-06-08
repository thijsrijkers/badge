package compiler

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"wind/tokenizer"
)

type Expr struct {
	Left     uint64
	Op       string
	Right    uint64
	IsBinary bool
}

func CompileLine(line string) error {
	tokens, err := tokenizer.Tokenize(line)
	if err != nil {
		return err
	}

	if len(tokens) < 4 || tokens[0].Type != tokenizer.TokenLet || tokens[1].Type != tokenizer.TokenIdent || tokens[2].Type != tokenizer.TokenEqual {
		return errors.New("invalid syntax")
	}

	exprTokens := tokens[3:]
	expr, err := parseExpr(exprTokens)
	if err != nil {
		return err
	}

	asmCode := generateASM(expr)

	err = os.WriteFile("out.asm", []byte(asmCode), 0644)
	if err != nil {
		return fmt.Errorf("failed to write asm file: %w", err)
	}

	cmdAssemble := exec.Command("nasm", "-f", "elf64", "out.asm", "-o", "out.o")
	if out, err := cmdAssemble.CombinedOutput(); err != nil {
		return fmt.Errorf("assembly failed: %w\n%s", err, out)
	}

	cmdLink := exec.Command("ld", "out.o", "-o", "out")
	if out, err := cmdLink.CombinedOutput(); err != nil {
		return fmt.Errorf("linking failed: %w\n%s", err, out)
	}

	cmdRun := exec.Command("./out")
	if err := cmdRun.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			ws := exitErr.Sys().(interface{ ExitStatus() int })
			fmt.Printf("Program exited with code: %d\n", ws.ExitStatus())
			return nil
		}
		return fmt.Errorf("execution error: %w", err)
	}

	fmt.Println("Program exited with code 0")
	return nil
}

func parseExpr(tokens []tokenizer.Token) (Expr, error) {
	if len(tokens) == 1 && tokens[0].Type == tokenizer.TokenNumber {
		val, err := strconv.ParseUint(tokens[0].Value, 10, 64)
		if err != nil {
			return Expr{}, err
		}
		return Expr{Left: val, IsBinary: false}, nil
	}

	if len(tokens) == 3 &&
		tokens[0].Type == tokenizer.TokenNumber &&
		tokens[1].Type == tokenizer.TokenOperator &&
		tokens[2].Type == tokenizer.TokenNumber {

		left, err := strconv.ParseUint(tokens[0].Value, 10, 64)
		if err != nil {
			return Expr{}, err
		}
		right, err := strconv.ParseUint(tokens[2].Value, 10, 64)
		if err != nil {
			return Expr{}, err
		}
		op := tokens[1].Value
		return Expr{Left: left, Right: right, Op: op, IsBinary: true}, nil
	}

	return Expr{}, errors.New("unsupported expression")
}

func generateASM(expr Expr) string {
	if !expr.IsBinary {
		return fmt.Sprintf(`section .text
global _start

_start:
    mov rax, %d
    mov rdi, rax
    mov rax, 60
    syscall
`, expr.Left)
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
		return "; unsupported operator\n"
	}

	return fmt.Sprintf(`section .text
global _start

_start:
    mov rax, %d
    %s rax, %d
    mov rdi, rax
    mov rax, 60
    syscall
`, expr.Left, asmOp, expr.Right)
}
