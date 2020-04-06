package display

// SmallDisplay ...
type SmallDisplay struct {
	Make               string
	Model              string
	CurrentlyDisplayed string
}

// NewSmallDisplay ...
func NewSmallDisplay(make, model string) *SmallDisplay {
	return &SmallDisplay{
		Make:  make,
		Model: model,
	}
}

// Show ...
func (sd *SmallDisplay) Show(text string) {
	sd.CurrentlyDisplayed = text
}

// Read ...
func (sd *SmallDisplay) Read() string {
	return sd.CurrentlyDisplayed
}
