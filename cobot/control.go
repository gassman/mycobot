package cobot

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

type Controller struct {
	Port     string
	Command  []byte
	Response []byte
}

func Initialize() *Controller {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	ctrl := new(Controller)

	return nil
}

func Connect() {

}

func Disconnect() {

}
