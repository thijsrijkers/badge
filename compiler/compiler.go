package compiler

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"wind/elf"
	"wind/tokenizer"
)

func CompileLine(line string) error {
	tokens, err := tokenizer.Tokenize(line)
	if err != nil {
		return err
	}

	if len(tokens) < 6 || tokens[0].Type != tokenizer.TokenLet ||
		tokens[1].Type != tokenizer.TokenIdent || tokens[1].Value != "i" ||
		tokens[2].Type != tokenizer.TokenEqual ||
		tokens[3].Type != tokenizer.TokenNumber ||
		tokens[4].Type != tokenizer.TokenPlus ||
		tokens[5].Type != tokenizer.TokenNumber {
		return errors.New("invalid syntax or token pattern")
	}

	left, err := strconv.ParseUint(tokens[3].Value, 10, 64)
	if err != nil {
		return errors.New("invalid left operand")
	}
	right, err := strconv.ParseUint(tokens[5].Value, 10, 64)
	if err != nil {
		return errors.New("invalid right operand")
	}
	if right > 0xFFFFFFFF {
		return errors.New("right operand too large for immediate")
	}

	code := []byte{
		0x48, 0xB8, // mov rax, imm64
	}
	code = append(code, intToBytes64(left)...)
	code = append(code, 0x48, 0x05) // add rax, imm32
	code = append(code, intToBytes32(uint32(right))...)
	code = append(code, 0x48, 0x89, 0xC7)                         // mov rdi, rax
	code = append(code, 0x48, 0xC7, 0xC0, 0x3C, 0x00, 0x00, 0x00) // mov rax, 60 (exit)
	code = append(code, 0x0F, 0x05)                               // syscall

	final := append(elf.ELFHeader, elf.ProgHeaderWithSize(uint64(len(code)))...)
	final = append(final, code...)

	err = os.WriteFile("out", final, 0755)
	if err != nil {
		return fmt.Errorf("failed to write ELF file: %w", err)
	}

	cmd := exec.Command("./out")
	err = cmd.Run()
	if err != nil {
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

// helpers
func intToBytes64(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func intToBytes32(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}
