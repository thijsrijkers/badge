# Wind

Wind is a lightweight scripting language that compiles directly to x86_64 machine code. It is designed for simplicity and performance, allowing scripts to execute as native binaries without any runtime or virtual machine.

## Overview

Wind bridges the gap between scripting and systems programming by providing a minimalistic language that maps directly to raw machine instructions. It is well-suited for educational purposes, experimentation with compiler design, and low-level systems programming.

## Goals

- Compile scripts directly to x86_64 machine code
- Eliminate external runtimes and dependencies
- Provide a minimal, readable syntax for low-level operations
- Offer a scripting experience with the performance of native execution

## Implementation

Wind is implemented in Go and processes `.zl` source files. The compiler reads ZLang code, parses it, and emits raw machine instructions targeting the x86_64 architecture. The output is a standalone executable file that can be run directly on Linux systems.

## Features

- Register-level arithmetic and data movement
- Direct system call access (e.g., write, exit)
- Minimal and predictable execution model
- Executable output without linking or intermediate formats

## Testing Wind Scripts

To test Wind scripts, create a source file with the `.zl` extension (e.g., `test.zl`) in the project directory. For example:

```bash
let i = 40 + 2
```

Then, run the main.go program which reads the Wind source file, compiles it into a native executable, runs the executable, and prints the program’s exit code:

```bash
go run main.go
```

You should see output like:

```bash
Program exited with code: 42
```

## Checking the Result on Different Platforms
After running the compiled Wind script, you can check the program's exit code directly from the shell:

- MacOS, Linux and WSL (Windows Subsystem for Linux):

```bash
echo $?
```

The number printed is the exit code returned by the Wind program (in this example, 42). This exit code confirms that your Wind script compiled and executed correctly as a native binary.


## Use Cases

- Educational tool for learning machine code and compiler development
- Rapid prototyping of low-level behaviors and utilities
- Building small, standalone command-line tools
- Exploration of syscall interfaces and instruction encoding

## License

This project is open source and available under the MIT License.
