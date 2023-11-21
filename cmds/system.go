package cmds

const (
	// Basic
	Header Instruction = 0xfe
	Footer Instruction = 0xfa

	// System Status
	RobotVersion    Instruction = 0x01
	SoftwareVersion Instruction = 0x02
	GetRobotId      Instruction = 0x03
	SetRobotId      Instruction = 0x04

	GetErrorInfo   Instruction = 0x07
	ClearErrorInfo Instruction = 0x08
	GetAtomVersion Instruction = 0x09

	// Overall Status
	PowerOn               Instruction = 0x10
	PowerOff              Instruction = 0x11
	IsPowerOn             Instruction = 0x12
	ReleaseAllServos      Instruction = 0x13
	IsControllerConnected Instruction = 0x14
	ReadNextError         Instruction = 0x15
	SetFreshMode          Instruction = 0x16
	GetFreshMode          Instruction = 0x17
	SetFreeMode           Instruction = 0x1A
	IsFreeMode            Instruction = 0x1B
	CobotxGetAngle        Instruction = 0x1C
	CobotxIsGoZero        Instruction = 0x1D
	FocusAllServos        Instruction = 0x18
	GoZero                Instruction = 0x19

	// Mode and operation
	GetAngles         Instruction = 0x20
	SendAngle         Instruction = 0x21
	SendAngles        Instruction = 0x22
	GetCoOrds         Instruction = 0x23
	SendCoOrd         Instruction = 0x24
	SendCoOrds        Instruction = 0x25
	Pause             Instruction = 0x26
	IsPaused          Instruction = 0x27
	Resume            Instruction = 0x28
	Stop              Instruction = 0x29
	IsInPosition      Instruction = 0x2A
	IsMoving          Instruction = 0x2B
	GetAngle          Instruction = 0x2C
	GetCoOrd          Instruction = 0x2D
	SendAnglesAuto    Instruction = 0x2E
	GetSolutionAngles Instruction = 0x2E
	SetSolutionAngles Instruction = 0x2F
)
