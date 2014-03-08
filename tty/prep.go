package tty

import (
	"io"
	"os"
)

type topType struct {}

// see rltty.c:595
// assumes NO_TTY_DRIVER is not defined
func PrepTermMode() {
	// FIXME: capture sigint around this function // signals.c:538

	// this could also be rl_instream, which indicates that maybe,
	// just maybe, we should be injecting stdin / out as a dependency

	// var tio topType
	tty := os.Stdin
	_, err := getTTYSettings(tty)
	if err != nil { // FIXME: check for ENOTTY or EINVAL or maybe ENOTSUP ??
		// rl_echoing_p = 1 // ??
		return
	}

}

func DePrepTermMode() {
}

type ttySettings struct { }

func getTTYSettings(pipe io.Reader) (settings ttySettings, err error) {
	return
}
