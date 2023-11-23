package cobot

import (
	log "github.com/sirupsen/logrus"
	"go.bug.st/serial"
)

type Controller struct {
	Ports      []string
	Port       string
	Mode       *serial.Mode
	Connection serial.Port
	PortState  PortStatus
}

type PortStatus int

const (
	PortClosed PortStatus = 0
	PortOpen   PortStatus = 1
)

func Initialize() *Controller {
	log.SetFormatter(&log.JSONFormatter{})
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Errorf("error getting serial port list, reason: %v\n", err)
		return nil
	}
	if len(ports) == 0 {
		log.Errorln("No serial ports found!no serial ports found")
		return nil
	}
	for _, port := range ports {
		log.Debugf("Found port: %v\n", port)
	}
	ctrl := new(Controller)
	ctrl.Ports = ports
	ctrl.PortState = PortClosed

	return ctrl
}

func (c *Controller) Connect() {
	if stringInSlice(c.Port, c.Ports) && c.Mode != nil {
		conn, err := serial.Open(c.Port, c.Mode)
		if err != nil {
			log.Errorf("error opening serial port, reason: %v\n", err)
		}
		c.Connection = conn
		c.PortState = PortOpen
	} else {
		log.Errorln("error connecting serial port, mode and port are not set.")
	}
}

func (c *Controller) Disconnect() {
	_ = c.Connection.Close()
	c.PortState = PortClosed
}

func (c *Controller) WriteData(data []byte) {
	if c.PortState != PortOpen {
		log.Errorln("error, serial port not opened or connected")
		return
	}
	txSize := len(data)
	dataSize, err := c.Connection.Write(data)
	if err != nil {
		log.Errorf("error writing data to serial port, reason %v\n", err)
		return
	}
	if txSize != dataSize {
		log.Errorln("error, less data written than supplied")
	}
}

func (c *Controller) ReadData() []byte {
	if c.PortState != PortOpen {
		log.Errorln("error, serial port not opened or connected")
		return []byte{}
	}
	var data []byte
	_, err := c.Connection.Read(data)
	if err != nil {
		log.Errorf("error reading data from serial port, reason: %v\n", err)
	}
	return data
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
