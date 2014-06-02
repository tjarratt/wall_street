// +build windows

// Copied without shame from github.com/cloudfoundry/cli
package echo

import (
	"os"
	"syscall"
)

// see SetConsoleMode documentation for bit flags
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms686033(v=vs.85).aspx
const ENABLE_ECHO_INPUT = 0x0004

var previousMode *uint32

func echoOn(fileDescriptors []uintptr) error {
	defer func() {
		previousMode = nil
	}()

	stdinHandle := syscall.Handle(os.Stdin.Fd())
	_, err := setConsoleMode(stdinHandle, previousMode)
	return err
}

func echoOff(fileDescriptors []uintptr) (int, error) {
	stdinHandle := syscall.Handle(os.Stdin.Fd())

	// attempt to preserve previous console mode, even if echoOff called multiple times
	if previousMode == nil {
		err := syscall.GetConsoleMode(stdinHandle, &previousMode)
		if err != nil {
			return err
		}
	}

	newMode := (previousMode &^ ENABLE_ECHO_INPUT)
	return setConsoleMode(stdinHandle, newMode)
}

func setConsoleMode(console syscall.Handle, mode uint32) (result int, err error) {
	dll := syscall.MustLoadDLL("kernel32")
	proc := dll.MustFindProc("SetConsoleMode")
	r, _, err := proc.Call(uintptr(console), uintptr(mode))

	return r, err
}
