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
	SetFreeMode           Instruction = 0x1a
	IsFreeMode            Instruction = 0x1b
	CobotxGetAngle        Instruction = 0x1c
	CobotxIsGoZero        Instruction = 0x1d
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
	IsInPosition      Instruction = 0x2a
	IsMoving          Instruction = 0x2b
	GetAngle          Instruction = 0x2c
	GetCoOrd          Instruction = 0x2d
	SendAnglesAuto    Instruction = 0x2e
	GetSolutionAngles Instruction = 0x2e
	SetSolutionAngles Instruction = 0x2f
	// Encoder
	SetEncoder     Instruction = 0x3a
	GetEncoder     Instruction = 0x3b
	SetEncoders    Instruction = 0x3c
	GetEncoders    Instruction = 0x3d
	SetEncoderDrag Instruction = 0x3e
	// Coordinate transformation
	SetToolReference  Instruction = 0x81
	GetToolReference  Instruction = 0x82
	SetWorldReference Instruction = 0x83
	GetWorldReference Instruction = 0x84
	SetReferenceFrame Instruction = 0x85
	GetReferenceFrame Instruction = 0x86
	SetMovementType   Instruction = 0x87
	GetMovementType   Instruction = 0x88
	SetEndType        Instruction = 0x89
	GetEndType        Instruction = 0x8a
	WriteMoveC        Instruction = 0x8c
)
