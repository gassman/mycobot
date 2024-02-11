package mycobot

import (
	"errors"
	"log"
	"mycobot/controller"
)

type Cobot struct {
	Device *controller.Connection
}

var MyDevice Cobot = Cobot{}

func InitCobot() {
	ctrl, err := controller.NewConnection(115200, 1)
	if err != nil {
		log.Fatalf("error initializing device, reason %v\n", err)
	}
	MyDevice.Device = ctrl
}

func SendMessage(data []byte) error {
	if MyDevice.Device == nil {
		return errors.New("cobot needs to initialized before sending commands")
	}
	err := (*MyDevice.Device).WriteData(data)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMessage() ([]byte, error) {
	return []byte{}, nil
}

func GetAngles() ([]float32, error) {

	return []float32{}, nil
}

func GetRadians() ([]float32, error) {

	return []float32{}, nil
}

func GetCoords() ([]float32, error) {

	return []float32{}, nil
}

func SendAngles(angles []float32) error {

	return nil
}

func SendRadians(coords []float32) error {

	return nil
}

func SendCoords(coords []float32) error {

	return nil
}

func SyncSendAngles(angles []float32, speed int, timeout int) error {

	return nil
}

func SyncSendCoords(coords []float32, speed int, timeout int) error {

	return nil
}

func Wait() {

}

func Close() {

}

func Open() {

}

func GetAccelertaion() (int, error) {
	return 1, nil
}

func SetAcceleration(accel int) error {

	return nil
}

func GetVersion() {

}

func PowerOn() {

}

func PowerOff() {

}

func IsPowerOn() (bool, error) {
	return true, nil
}

func IsControllerConnected() (bool, error) {
	return true, nil
}

func IsInPosition() (bool, error) {
	return true, nil
}

func IsMoving() (bool, error) {
	return true, nil
}

func ReleaseAllServos() {

}

//
// JOG Mode Functions
//

func JogAngle() {

}

func JogCoord() {

}

func JogStop() {

}

func JogSetEncoder() {

}

func JogGetEncoder() {

}

func JogGetEncoders() {

}

func JogPause() {

}

func JogResume() {

}

func JogStopAll() {

}

func IsPaused() {

}

// Opeartional status and settings
func GetEncoder() {

}

func SetEncoder() {

}

func GetEncoders() {

}

func SetEncoders() {

}

func GetSpeed() {

}

func SetSpeed() {

}

func GetJointMinAngle() {

}

func GetJointMaxAngle() {

}

//
// Servo Control
//

func IsServoEnabled() {

}

func AreAllServosEnabled() {

}

func SetServoData() {

}

func GetServoData() {

}

func SetServoCallibration() {

}

func ReleaseServo() {

}

func FocusServo() {

}

//
// AtomIO
//
