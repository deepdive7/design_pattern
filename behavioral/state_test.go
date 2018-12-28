package behavioral

import "testing"

func TestState(t *testing.T) {
	light := &Light{State: &OnLightState{}}
	light.PressSwitch()
	light.PressSwitch()
	light.PressSwitch()
	light.PressSwitch()
}
