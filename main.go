package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"wind/elf"
	"wind/compiler"
)

func main() {
	// Read Wind source code from test.zl file
	source, err := ioutil.ReadFile("test.wind")
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	line := string(source)

	// Compile the Wind code line
	err = compiler.CompileLine(line)
	if err != nil {
		log.Fatalf("Compile error: %v", err)
	}

	// Run the compiled ELF executable and get its exit code
	exitCode, err := elf.RunElfAndGetExitCode("./out")
	if err != nil {
		log.Fatalf("Run error: %v", err)
	}

	fmt.Printf("Program exited with code: %d\n", exitCode)
}
