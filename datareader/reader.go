package datareader

import (
	"fmt"

	"github.com/sapariduo/tcpserver/utils"
)

type DataBlock struct {
	byte1 BlockTuple
	byte2 BlockTuple
	byte4 BlockTuple
	byte8 BlockTuple
}

type BlockTuple struct {
	len   int
	start int
	end   int
}

func Read(data string) {
	fmt.Println(data)
	pos := 0
	// fmt.Println("pos", pos)
	var slot []int = []int{1, 2, 4, 8}
	for x, y := range slot {
		fmt.Println("pos", pos)
		len := data[pos : pos+2]
		fmt.Println("hex data len:", len)
		d := utils.Hex2int(len)
		if d == 0 {
			fmt.Println("inside")
			pos += 2
			fmt.Println("block", slot[x], "lenght", d, "bytelenght", y)
			fmt.Println("start", pos, "end", pos+(d*y), "lenght", 2*d*y)
			continue
		}
		pos += 2
		length := (d + (d * y)) * 2
		fmt.Println("block", slot[x], "lenght", d, "bytelength", y)
		fmt.Println("start", pos, "end", pos+length, "lenght", length)
		fmt.Println(data[pos : pos+length])
		// fmt.Println(data[pos : pos+(d*2*y)])
		pos += length
	}
}
