package display

// Display ...
type Display interface {
	Show(text string)
	Read() string
}
