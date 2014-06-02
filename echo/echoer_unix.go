// Copied without shame from https://code.google.com/p/gopass/
// and github.com/cloudfoundry/cli

// +build darwin freebsd linux netbsd openbsd

package echo

import (
	"fmt"
	"syscall"
)

const (
	sttyArg0 = "/bin/stty"
)

var (
	ws           syscall.WaitStatus = 0
	sttyArgvEOn  []string           = []string{"stty", "echo"}
	sttyArgvEOff []string           = []string{"stty", "-echo"}
)

func echoOff(fd []uintptr) (int, error) {
	pid, err := syscall.ForkExec(sttyArg0, sttyArgvEOff, &syscall.ProcAttr{Files: fd})

	if err != nil {
		return 0, fmt.Errorf("failed turning off console echo:\n\t%s", err)
	}

	return pid, nil
}

func echoOn(fd []uintptr) error {
	pid, err := syscall.ForkExec(sttyArg0, sttyArgvEOn, &syscall.ProcAttr{Files: fd})

	if err == nil {
		syscall.Wait4(pid, &ws, 0, nil)
	}

	return err
}
