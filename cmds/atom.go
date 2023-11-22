package cmds

const (
	SetPinMode            Instruction = 0x60
	SetDigitalOutput      Instruction = 0x61
	GetDigitalInput       Instruction = 0x62
	SetPWMMode            Instruction = 0x63
	SetPWMOutput          Instruction = 0x64
	GetGripperValue       Instruction = 0x65
	SetGripperState       Instruction = 0x66
	SetGripperValue       Instruction = 0x67
	SetGripperCalibration Instruction = 0x68
	IsGripperMoving       Instruction = 0x69
	SetColor              Instruction = 0x6a
	SetGripperTorque      Instruction = 0x6f
	SetColorMyarm         Instruction = 0x70
	SetElectricGripper    Instruction = 0x6b
	InitElectricGripper   Instruction = 0x6c
	SetGripperMode        Instruction = 0x6d
	GetGripperMode        Instruction = 0x6e
	// Basic Atom I/O
	SetBasicOutput Instruction = 0xa0
	GetBasicInput  Instruction = 0xa1
	GetBaseInput   Instruction = 0xa2
	SetBasePWM     Instruction = 0xa5
)
