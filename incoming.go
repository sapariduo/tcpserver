package main

import (
	"fmt"

	"github.com/sapariduo/tcpserver/message"
	"github.com/sapariduo/tcpserver/utils"
)

var (
	packetsize = 2
	imei       = 8
	command    = 1
	crc        = 2
	recleft    = 1
	nbrec      = 1
)

var (
	packet = []Pair{{"packet", 4}, {"imei", 16}, {"command", 2}, {"payload", 0}, {"crc", 4}}
)

type Pair struct {
	label string
	value int
}

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

func msgParser(input string) *message.Message {
	msg := message.New()
	// cmd := type1.New()
	// output := []int{}
	for _, v := range packet {
		switch x := v.label; x {
		case "packet":
			val := utils.Hex2int(input[:v.value])
			msg.PacketLength = val
			input = input[v.value:]
			fmt.Println(v.label, val)
		case "imei":
			val := utils.Hex2int(input[:v.value])
			msg.Imei = val
			input = input[v.value:]
			fmt.Println(v.label, val)
		case "command":
			val := utils.Hex2int(input[:v.value])
			msg.Command = val
			input = input[v.value:]
			fmt.Println(v.label, val)
		case "payload":
			msg.Payload = input[v.value : len(input)-4]

		case "crc":
			msg.Crc = input[len(input)-4:]
			fmt.Println(v.label, msg.Crc)
		}

	}
	return msg
}
