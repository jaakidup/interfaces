package keyboard

// InternalKeyboard ...
type InternalKeyboard struct {
	Make  string
	Model string
}

// NewInternalKeyboard ...
func NewInternalKeyboard(make, model string) *InternalKeyboard {
	return &InternalKeyboard{
		Make:  make,
		Model: model,
	}
}

// PressKey ...
func (ik *InternalKeyboard) PressKey(keycode string) string {
	return keycode
}
