package controller

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"go.bug.st/serial"
)

type Controller interface {
	Connect(port string) error
	Disconnect() error
	WriteData(data []byte) error
	ReadData() ([]byte, error)
}

type PortStatus int

const (
	PortClosed PortStatus = 0
	PortOpen   PortStatus = 1
)

type Connection struct {
	ConnectedPort serial.Port
	PortMode      *serial.Mode
	PortState     PortStatus
	PortList      []string
}

func NewConnection(baud int, stopBits int) (*Connection, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	if len(ports) == 0 {
		return nil, errors.New("no serial ports found")
	}
	for _, port := range ports {
		log.Debugf("Found port: %v\n", port)
	}
	conn := new(Connection)
	conn.PortMode = new(serial.Mode)
	conn.PortMode.BaudRate = baud
	conn.PortMode.DataBits = 8
	conn.PortMode.Parity = serial.NoParity
	conn.PortMode.StopBits = serial.StopBits(stopBits)
	conn.PortState = PortClosed
	return conn, nil
}

func (c *Connection) Connect(port string) error {
	if stringInSlice(port, c.PortList) && c.PortMode != nil {
		conn, err := serial.Open(port, c.PortMode)
		if err != nil {
			return err
		}
		c.ConnectedPort = conn
		c.PortState = PortOpen
	} else {
		return errors.New("error connecting mode or port are not set")
	}
	return nil
}

func (c *Connection) Disconnect() error {
	c.PortState = PortClosed
	if err := c.ConnectedPort.Close(); err != nil {
		return err
	}
	return nil
}

func (c *Connection) WriteData(data []byte) error {
	if c.PortState != PortOpen {
		return errors.New("serial port not opened or connected")
	}
	txSize := len(data)
	dataSize, err := c.ConnectedPort.Write(data)
	if err != nil {
		return err
	}
	if txSize != dataSize {
		return errors.New("less data written than supplied")
	}
	return nil
}

func (c *Connection) ReadData() ([]byte, error) {
	if c.PortState != PortOpen {
		return []byte{}, errors.New("serial port not opened or connected")
	}
	var data []byte
	_, err := c.ConnectedPort.Read(data)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
