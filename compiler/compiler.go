package compiler

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"wind/expr"
	"wind/tokenizer"
)

func CompileLines(source string) error {
	lines := strings.Split(source, "\n")

	asmLines := []string{
		"section .text",
		"global _start",
		"_start:",
	}

	// Store variable values to resolve references
	vars := make(map[string]uint64)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens, err := tokenizer.Tokenize(line)
		if err != nil {
			return err
		}

		if len(tokens) < 4 || tokens[0].Type != tokenizer.TokenLet || tokens[1].Type != tokenizer.TokenIdent || tokens[2].Type != tokenizer.TokenEqual {
			return errors.New("invalid syntax")
		}

		varName := tokens[1].Value

		exprTokens := tokens[3:]
		exprParsed, err := expr.ParseExpr(exprTokens, vars)
		if err != nil {
			return err
		}

		asmCode := expr.GenerateASM(exprParsed)
		asmLines = append(asmLines, asmCode...)

		// Calculate and store variable value
		val := exprParsed.Left
		if exprParsed.IsBinary {
			switch exprParsed.Op {
			case "+":
				val = exprParsed.Left + exprParsed.Right
			case "-":
				val = exprParsed.Left - exprParsed.Right
			case "*":
				val = exprParsed.Left * exprParsed.Right
			default:
				return fmt.Errorf("unsupported operator %s", exprParsed.Op)
			}
		}
		vars[varName] = val
	}

	// Exit syscall using rax value from last expression
	asmLines = append(asmLines,
		"    mov rdi, rax",
		"    mov rax, 60",
		"    syscall",
	)

	asmCode := strings.Join(asmLines, "\n")

	err := os.WriteFile("out.asm", []byte(asmCode), 0644)
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

	return nil
}
