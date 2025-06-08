package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"wind/compiler"
	"wind/elf"
)

func main() {
	// Read Wind source code from test.wind file
	source, err := ioutil.ReadFile("test.wind")
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
