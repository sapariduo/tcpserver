package message

type Message struct {
	PacketLength int
	Imei         int
	Command      int
	Payload      string
	Crc          string
}

func (msg *Message) Type() int {
	return msg.Command
}

func New() *Message {
	return new(Message)
}