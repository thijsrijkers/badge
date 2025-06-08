package elf

import (
	"fmt"
	"syscall"
)

func RunElfAndGetExitCode(path string) (int, error) {
    pid, err := syscall.ForkExec(path, []string{path}, &syscall.ProcAttr{
        Files: []uintptr{uintptr(syscall.Stdin), uintptr(syscall.Stdout), uintptr(syscall.Stderr)},
    })
    if err != nil {
        return -1, err
    }

    var ws syscall.WaitStatus
    _, err = syscall.Wait4(pid, &ws, 0, nil)
    if err != nil {
        return -1, err
    }

    if ws.Exited() {
        return ws.ExitStatus(), nil
    }

    return -1, fmt.Errorf("process did not exit normally")
}

 