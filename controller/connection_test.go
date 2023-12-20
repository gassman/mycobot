package controller

import (
	"testing"
)

var conn *Connection
var err error

// TestNewConnection calls NewConnection with baud and stop bits
// Returned structure shoud contain a list of available ports
func TestNewConnection(t *testing.T) {
	conn, err = NewConnection(9600, 1)
	if conn == nil || err != nil {
		t.Fatalf(`Error creating a new conection , %v`, err)
	}
}

// TestConnect calls Connect with a selected port,
// checking for an error.
func TestConnect(t *testing.T) {
	if conn == nil {
		t.Fatalf(`Looks like a previous test failed, giving up`)
	}
	err = conn.Connect("/dev/ttyS0")
	if err != nil {
		t.Fatalf(`Error connecting to device, %v, check permissions`, err)
	}
}

// TestWriteData calls WriteData to send message bytes to device connected to serial port,
// checking for an error.
func TestWriteData(t *testing.T) {
	if conn == nil {
		t.Fatalf(`Looks like a previous test failed, giving up`)
	}
	msg := []byte("hello")
	err = conn.WriteData(msg)
	if err != nil {
		t.Fatalf(`Error sending bytes to device, %v`, err)
	}
}

// TestReadData calls ReadData to receive message bytes from device connected to serial port,
// checking for an error.
func TestReadData(t *testing.T) {
	if conn == nil {
		t.Fatalf(`Looks like a previous test failed, giving up`)
	}
	_, err = conn.ReadData()
	if err != nil {
		t.Fatalf(`Error reading bytes from device, %v`, err)
	}
}

// TestDisconnect calls Disconnect to terminate the serial port conection,
// checking for an error.
func TestDisconnect(t *testing.T) {
	if conn == nil {
		t.Fatalf(`Looks like a previous test failed, giving up`)
	}
	err = conn.Disconnect()
	if err != nil {
		t.Fatalf(`Error disconnecting from device, %v`, err)
	}
}
