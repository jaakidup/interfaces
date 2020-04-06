package keyboard

// Keyboard ...
type Keyboard interface {
	PressKey(keycode string) string
}
