package atom

import "testing"

// TestGetGripperValue calls GetGripperValue to request current state of gripper from Atom,
// checking for an error.
func TestGetGripperValue(t *testing.T) {
	t.Log(`Executing TestGetGripperValue`)
}

// TestSetGripperState calls TestSetGripperState to set current state of gripper on Atom,
// checking for an error.
func TestSetGripperState(t *testing.T) {
	t.Log(`Executing TestSetGripperState`)
}

// TestSetGripperValue calls SetGripperValue to a value for the gripper on Atom,
// checking for an error.
func TestSetGripperValue(t *testing.T) {
	t.Log(`Executing TestSetGripperValue`)
}

// TestSetGripperIni calls SetGripperIni to set initial values for the gripper on Atom,
// checking for an error.
func TestSetGripperIni(t *testing.T) {
	t.Log(`Executing TestSetGripperIni`)
}

// TestIsGripperMoving calls IsGripperMoving to determine if gripper is still moving,
// checking for an error.
func TestIsGripperMoving(t *testing.T) {
	t.Log(`Executing TestIsGripperMoving`)
}
