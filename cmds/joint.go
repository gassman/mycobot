package cmds

const (
	// RUNNING STATUS AND SETTINGS
	GetSpeed         Instruction = 0x40
	SetSpeed         Instruction = 0x41
	GetFeedOverride  Instruction = 0x42
	GetAcceleration  Instruction = 0x44
	SetAcceleration  Instruction = 0x45
	GetJointMinAngle Instruction = 0x4a
	GetJointMaxAngle Instruction = 0x4b
	SetJointMin      Instruction = 0x4c
	SetJointMax      Instruction = 0x4d
)
