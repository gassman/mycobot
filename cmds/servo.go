package cmds

const (
	// Servo control
	IsServoEnable       = 0x50
	IsAllServoEnable    = 0x51
	SetServoData        = 0x52
	GetServoData        = 0x53
	SetServoCalibration = 0x54
	JointBrake          = 0x55
	ReleaseServo        = 0x56
	FocusServo          = 0x57

	// Motor status read
	GetServoSpeed      = 0xe1
	GetServoCurrents   = 0xe2
	GetServoVoltages   = 0xe3
	GetServoStatus     = 0xe4
	GetServoTemps      = 0xe5
	GetServoLastpdi    = 0xe6
	ServoRestore       = 0xe7
	SetErrorDetectMode = 0xe8
	GetErrorDetectMode = 0xe9
)
