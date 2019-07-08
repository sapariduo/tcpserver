package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func Hex2int(hexStr string) int {
	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(hexStr, 16, 64)
	return int(result)
}

func binstring2int(input string) int64 {
	ret, _ := strconv.ParseInt(input, 2, 64)
	// fmt.Printf("%032b\n", uint32(ret))
	ret ^= 0xFFFFFFFF
	// rev := bits.Reverse32(uint32(ret))
	// fmt.Printf("%032b\n", ret+1)
	return ret + 1
}

func hex2bin(input string) int64 {
	ret, _ := strconv.ParseInt(input, 16, 64)
	fmt.Printf("%b\n", ret)
	out, _ := strconv.ParseInt(fmt.Sprintf("%b", ret), 2, 64)
	fmt.Println(out)
	return int64(out)
}

func hex2int32(input string) int32 {
	ret, _ := strconv.ParseInt(input, 16, 32)
	return int32(ret)
}

func hex2int64(input string) int64 {
	ret, _ := strconv.ParseInt(input, 16, 64)
	return ret
}

func Hex2binary(hexStr string) int {
	// fmt.Println(hexStr, "initial value")
	result, _ := strconv.ParseUint(hexStr, 16, 64)
	// fmt.Println(result)
	// Convert int to binary representation
	// %024b indicates base 2, padding with 0, with 24 characters.
	bin := fmt.Sprintf("%032b", result)
	if bin[:1] == "1" {
		// fmt.Println("inside 1", bin)

		// bin = fmt.Sprintf(bin[1:])
		// fmt.Printf("%s\n", bin)
		result := binstring2int(bin)
		// bin_ := Bin2int(bin)
		// fmt.Printf("int : %d\n", result)
		return int(-result)
	} else {
		// res, _ := strconv.ParseInt(bin, 2, 64)
		// fmt.Printf("binary :%d \n", res)
		return int(result)
	}

}

func Bin(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0b" + strconv.FormatInt(i64, 2) // base 2 for binary
	} else {
		return strconv.FormatInt(i64, 2) // base 2 for binary
	}
}

func Hex(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0x" + strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	} else {
		return strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	}
}

func Bin2int(binStr string) int {

	// base 2 for binary
	result, _ := strconv.ParseInt(binStr, 2, 64)
	return int(result)
}

func Crc(byteArray []byte) uint16 {
	var crc uint16
	for i := 0; i < len(byteArray); i++ {
		b := uint16(byteArray[i])
		q := (crc ^ b) & 0x0f
		crc = (crc >> 4) ^ (q * 0x1081)
		q = (crc ^ (b >> 4)) & 0xf
		crc = (crc >> 4) ^ (q * 0x1081)
	}
	return (crc >> 8) ^ (crc << 8)
}

func HexT(source string) int {
	dst := make([]byte, hex.DecodedLen(len(source)))
	n, err := hex.Decode(dst, []byte(source))
	if err != nil {
		log.Fatal(err)
	}
	return n

}
