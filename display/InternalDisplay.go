package display

import "fmt"

// InternalDisplay ...
type InternalDisplay struct {
	Make               string
	Model              string
	CurrentlyDisplayed string
}

// NewInternalDisplay ...
func NewInternalDisplay(make, model string) *InternalDisplay {
	return &InternalDisplay{
		Make:  make,
		Model: model,
	}
}

// Show ...
func (sd *InternalDisplay) Show(text string) {
	sd.CurrentlyDisplayed = text
	fmt.Println("Display", text)
}

// Read ...
func (sd *InternalDisplay) Read() string {
	return sd.CurrentlyDisplayed
}
