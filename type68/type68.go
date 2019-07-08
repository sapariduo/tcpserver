package type68

import (
	"github.com/sapariduo/tcpserver/message"
	"github.com/sapariduo/tcpserver/utils"
)

type Records struct {
	Left  int
	Total int
	Data  string
}
type Pair struct {
	label string
	value int
}

var (
	record      = []Pair{{"recleft", 2}, {"nbrecord", 2}, {"data", 0}}
	header      = []int{4, 1, 1, 1, 4, 4, 2, 2, 1, 2, 1, 2}
	headerLabel = []string{"timestamp", "timestampext", "recext", "priority", "long", "lat", "alt", "angle", "sat", "speed", "hdop", "eid"}
)

func New() *Records {
	return new(Records)
}

func (recd *Records) Header() map[string]int {

	// var len int
	ret := map[string]int{}

	for i, idx := range header {
		if headerLabel[i] == "lat" || headerLabel[i] == "long" {
			val := utils.Hex2binary(recd.Data[:idx*2])
			recd.Data = recd.Data[idx*2:]
			ret[headerLabel[i]] = val
		} else {
			val := utils.Hex2int(recd.Data[:idx*2])
			recd.Data = recd.Data[idx*2:]
			ret[headerLabel[i]] = val
		}

	}

	return ret

}

func (rcrd *Records) Records(msg *message.Message) *Records {
	// rcrd := new(Records)
	input := msg.Payload
	for _, v := range record {
		switch x := v.label; x {
		case "recleft":
			val := utils.Hex2int(input[:v.value])
			rcrd.Left = val
			input = input[v.value:]
			// fmt.Println(v.label, val)
		case "nbrecord":
			val := utils.Hex2int(input[:v.value])
			rcrd.Total = val
			input = input[v.value:]
			// fmt.Println(v.label, val)
		case "data":
			rcrd.Data = input[v.value:]
		}
	}

	return rcrd
}
