package controller

import (
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/host/v3"
)

var state *driverreg.State

const defaultTimeout = 15

func GPIOInit() error {
	var err error
	state, err = host.Init()
	if err != nil {
		return err
	}
	return nil
}

func GPIOOut(pin int, enable bool) error {

	return nil
}
