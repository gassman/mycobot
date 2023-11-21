package mycobot

import (
	"mycobot/cmds"
)

type Message struct {
	Id        cmds.Instruction
	Reply     bool
	ReplySize int
	Size      int
	Data      []byte
}

type Command struct {
	Send    Message
	Receive Message
}

var msgs = [...]Message{
	{Id: cmds.PowerOn, Reply: false, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x10, 0xfa}},
	{Id: cmds.PowerOff, Reply: false, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x11, 0xfa}},
	{Id: cmds.IsPowerOn, Reply: true, ReplySize: 6, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x12, 0xfa}},
	{Id: cmds.ReleaseAllServos, Reply: false, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x13, 0xfa}},
	{Id: cmds.IsControllerConnected, Reply: true, ReplySize: 6, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x14, 0xfa}},
	{Id: cmds.SetFreshMode, Reply: false, Size: 6, Data: []byte{0xfe, 0xfe, 0x03, 0x16, 0x01, 0xfa}},
	{Id: cmds.SetFreeMode, Reply: false, Size: 6, Data: []byte{0xfe, 0xfe, 0x03, 0x1a, 0x01, 0xfa}},
	{Id: cmds.IsFreeMode, Reply: true, ReplySize: 6, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x1b, 0xfa}},
	{Id: cmds.GetAngles, Reply: true, ReplySize: 17, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x20, 0xfa}},
	{Id: cmds.SendAngle, Reply: false, Size: 9, Data: []byte{0xfe, 0xfe, 0x06, 0x21, 0x00, 0x00, 0x00, 0x00, 0xfa}},
	{Id: cmds.SendAngles, Reply: false, Size: 18, Data: []byte{0xfe, 0xfe, 0x06, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfa}},
	{Id: cmds.GetCoOrds, Reply: true, ReplySize: 17, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x23, 0xfa}},
	{Id: cmds.SendCoOrd, Reply: false, Size: 9, Data: []byte{0xfe, 0xfe, 0x06, 0x24, 0x00, 0x00, 0x00, 0x00, 0xfa}},
	{Id: cmds.SendCoOrds, Reply: false, Size: 19, Data: []byte{0xfe, 0xfe, 0x06, 0x25, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfa}},
	{Id: cmds.Pause, Reply: false, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x26, 0xfa}},
	{Id: cmds.IsPaused, Reply: true, ReplySize: 6, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x27, 0xfa}},
	{Id: cmds.Resume, Reply: false, Size: 6, Data: []byte{0xfe, 0xfe, 0x03, 0x28, 0x01, 0xfa}},
	{Id: cmds.Stop, Reply: false, Size: 5, Data: []byte{0xfe, 0xfe, 0x02, 0x20, 0xfa}},
	{Id: cmds.IsInPosition, Reply: true, ReplySize: 6, Size: 18, Data: []byte{0xfe, 0xfe, 0x06, 0x2a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfa}},
}

func NewMessage(msgId cmds.Instruction, data []byte) *Message {
	newMsg := new(Message)
	newMsg.Id = msgId
	newMsg.BuildMessageData(data)
	return newMsg
}

func (m *Message) BuildMessageData(data []byte) {
	m.Size = len(data) + 3
	m.Data = make([]byte, m.Size+2)
	m.Data = append(m.Data, byte(cmds.Header), byte(cmds.Header))
	m.Data = append(m.Data, byte(m.Size))
	m.Data = append(m.Data, data...)
	m.Data = append(m.Data, byte(cmds.Footer))
}

func Short2HighLow(short uint16) (byte, byte) {
	return uint8(short >> 8), uint8(short & 0xff)
}

func HighLow2Short(high byte, low byte) uint16 {
	return uint16(low) | uint16(high)<<8
}
