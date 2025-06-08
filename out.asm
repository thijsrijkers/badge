section .text
global _start
_start:
    mov rax, 5
    mov rax, 5
    add rax, 3
    mov rdi, rax
    mov rax, 60
    syscall