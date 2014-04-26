package main

import (
	"github.com/tjarratt/wall_street"
)

func main() {
	reader := wall_street.NewReadline()
	answer := reader.Readline("Would you like to play a game? ")
	if answer == "yes" {
		println("cool!")
	} else {
		println("oh... okay? goodbye!")
	}
}
