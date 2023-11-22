package cmds

const (
	// Servo control
	IsServoEnable       Instruction = 0x50
	IsAllServoEnable    Instruction = 0x51
	SetServoData        Instruction = 0x52
	GetServoData        Instruction = 0x53
	SetServoCalibration Instruction = 0x54
	JointBrake          Instruction = 0x55
	ReleaseServo        Instruction = 0x56
	FocusServo          Instruction = 0x57

	// Motor status read
	GetServoSpeed      Instruction = 0xe1
	GetServoCurrents   Instruction = 0xe2
	GetServoVoltages   Instruction = 0xe3
	GetServoStatus     Instruction = 0xe4
	GetServoTemps      Instruction = 0xe5
	GetServoLastpdi    Instruction = 0xe6
	ServoRestore       Instruction = 0xe7
	SetErrorDetectMode Instruction = 0xe8
	GetErrorDetectMode Instruction = 0xe9
)
