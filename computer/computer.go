package computer

import (
	"github.com/jaakidup/interfaces/display"
	"github.com/jaakidup/interfaces/keyboard"
)

// NewComputer ...
func NewComputer(make string, model string, display display.Display, keyboard keyboard.Keyboard) *Computer {
	return &Computer{
		Make:     make,
		Model:    model,
		Display:  display,
		Keyboard: keyboard,
	}
}

// Computer ...
type Computer struct {
	Make     string
	Model    string
	Display  display.Display
	Keyboard keyboard.Keyboard
}

// Test ...
func (c *Computer) Test() {
	c.Display.Show("Start Test procedure")
	testkeys := c.Keyboard.PressKey("This is coming from the keyboard")
	c.Display.Show(testkeys)

	c.Display.Show("End Test procedure")

}
