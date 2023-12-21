package atom

import (
	"errors"
	"mycobot"
	"mycobot/cmds"
)

type AtomPinMode byte

const (
	AtomPinWorking    AtomPinMode = 0
	AtomPinStopped    AtomPinMode = 1
	AtomPinPullUpMode AtomPinMode = 2
)

func SetColor(rgb []byte) error {
	if len(rgb) != 3 {
		return errors.New("need 3 x byte RGB value to set color")
	}
	msg := cmds.NewMessage(cmds.SetColor, rgb)
	if err := mycobot.SendMessage(msg.Data); err != nil {
		return err
	}
	return nil
}

func SetPinMode(pin int8, mode AtomPinMode) error {
	if pin < 1 || pin > 5 {
		return errors.New("pin value must be between 1 and 5")
	}
	if mode > 2 {
		return errors.New("mode must be between 0 and 2")
	}
	vals := []byte{byte(pin), byte(mode)}
	msg := cmds.NewMessage(cmds.SetPinMode, vals)
	if err := mycobot.SendMessage(msg.Data); err != nil {
		return err
	}
	return nil
}

func SetDigitalOutput(pin int8, mode AtomPinMode) error {
	if pin < 1 || pin > 5 {
		return errors.New("pin value must be between 1 and 5")
	}
	if mode > 1 {
		return errors.New("mode must be either 0 or 1")
	}
	vals := []byte{byte(pin), byte(mode)}
	msg := cmds.NewMessage(cmds.SetDigitalOutput, vals)
	if err := mycobot.SendMessage(msg.Data); err != nil {
		return err
	}
	return nil

}

func GetDigitalInput() {

}

func SetPWMMode() {

}

func SetPWMOutput() {

}
