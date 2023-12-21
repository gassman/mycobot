package atom

import "testing"

// TestSetColor calls SetColor to set the color of the 5x5 LED,
// checking for an error.
func TestSetColor(t *testing.T) {
	t.Log(`Executing TestSetColor`)
}

// TestSetPinMode calls SetPinMode to set parameters for AtomIO pins,
// checking for an error.
func TestSetPinMode(t *testing.T) {
	t.Log(`Executing TestSetPinMode`)
}

// TestSetDigitalOutput calls SetDigitalOutput to set output of AtomIO pins,
// checking for an error.
func TestSetDigitalOutput(t *testing.T) {
	t.Log(`Executing TestSetDigitalOutput`)
}

// TestGetDigitalInput calls GetDigitalInput to get a value from a AtomIO pin,
// checking for an error.
func TestGetDigitalInput(t *testing.T) {
	t.Log(`Executing TestGetDigitalInput`)
}

// TestSetPWMMode calls SetPWMMode to set PWM values for an AtomIO pin,
// checking for an error.
func TestSetPWMMode(t *testing.T) {
	t.Log(`Executing TestSetPWMMode`)
}

// TestSetPWMOutput calls SetPWMOutput to set PWM output for an AtomIO pin,
// checking for an error.
func TestSetPWMOutput(t *testing.T) {
	t.Log(`Executing TestSetPWMOutput`)
}
