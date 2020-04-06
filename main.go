package main

import (
	"github.com/jaakidup/interfaces/computer"
	"github.com/jaakidup/interfaces/display"
	"github.com/jaakidup/interfaces/keyboard"
)

var apple computer.Computer

func main() {

	apple := computer.Computer{
		Make:     "Apple",
		Model:    "Macbook",
		Display:  display.NewInternalDisplay("LG", "TrueTone"),
		Keyboard: keyboard.NewInternalKeyboard("Apple", "Magick Keyboard"),
	}
	apple.Test()
}
