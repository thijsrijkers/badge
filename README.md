# Zul (lang)

Zul is a lightweight scripting language that compiles directly to x86_64 machine code. It is designed for simplicity and performance, allowing scripts to execute as native binaries without any runtime or virtual machine.

## Overview

ZLang bridges the gap between scripting and systems programming by providing a minimalistic language that maps directly to raw machine instructions. It is well-suited for educational purposes, experimentation with compiler design, and low-level systems programming.

## Goals

- Compile scripts directly to x86_64 machine code
- Eliminate external runtimes and dependencies
- Provide a minimal, readable syntax for low-level operations
- Offer a scripting experience with the performance of native execution

## Implementation

ZLang is implemented in Go and processes `.zl` source files. The compiler reads ZLang code, parses it, and emits raw machine instructions targeting the x86_64 architecture. The output is a standalone executable file that can be run directly on Linux systems.

## Features

- Register-level arithmetic and data movement
- Direct system call access (e.g., write, exit)
- Minimal and predictable execution model
- Executable output without linking or intermediate formats

## Use Cases

- Educational tool for learning machine code and compiler development
- Rapid prototyping of low-level behaviors and utilities
- Building small, standalone command-line tools
- Exploration of syscall interfaces and instruction encoding

## License

This project is open source and available under the MIT License.
