package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"badge/compiler"
	"badge/elf"
)

func main() {
	// Read Badge source code from test.badge file
	source, err := ioutil.ReadFile("test.badge")
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	// Compile the entire source (multiple lines)
	err = compiler.CompileLines(string(source))
	if err != nil {
		log.Fatalf("Compile error: %v", err)
	}
	fmt.Println("Compilation successful")

	// Run the compiled ELF executable and get its exit code
	exitCode, err := elf.RunElfAndGetExitCode("./out")
	if err != nil {
		log.Fatalf("Run error: %v", err)
	}

	fmt.Printf("Program exited with code: %d\n", exitCode)
}
