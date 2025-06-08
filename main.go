package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"zul/elf"
	"zul/compiler"
)

func main() {
	// Read Zul source code from test.zl file
	source, err := ioutil.ReadFile("test.zl")
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	line := string(source)

	// Compile the Zul code line
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
