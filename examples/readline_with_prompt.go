package main

import "github.com/tjarratt/wall_street"

func main() {
	reader := wall_street.NewReadline()
	reader.MaskUserInput = true

	secret := reader.Readline("Would you like to share a secret?")
	if secret == "secret" {
		println("OH MY!")
	}
}
