package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var fullPackage = `033500000C076B5C208F4401085D01B5960000003FE6E59AFC28ABB6023B1BC61500000500070E000500001B16000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D656B001E0EC20016000400170001008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01B9190000003FE6E59AFC28ABB6026C1BC61600000500070E000500001B16000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D6566001E0EC10016000600170002008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01BC9D0000003FE6E59AFC28ABB6024C1BC61500000500070E000500001B16000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D6559001E0EC10016000400170001008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01C0210000003FE6E59AFC28ABB6024E1BC61600000500070E000500001B16000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D6563001E0EBF0016000400170001008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01C3A50000003FE6E59AFC28ABB6020E1BC61400000500070E000500001B17000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D654F001E0EC10016000400170003008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01C7290000003FE6E59AFC28ABB6020F1BC61600000500070E000500001B00000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D6540001E0EBF0016000400170001008B000000890000008300000300410144B38A00960000000000AF00020A05005D01C8230000003FE6E59AFC28ABB602021BC61500000500070E000500001B15000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D651B001E0EBE0016000D00170001008B000000890000008300000300410144B38A00960000C74200AF00020A05005D01C8270000013FE6E4D2FC28ACD102024A600F00000700050E000501001B15000200000300000400001C0100AD00008600008700008800008200008F0001950101960107001D7013001E0EBF001600DF00170001008B000300890000008300000300410144B38A00960000C74200AF00020A0500ADE9`

func Test_hex2bin(t *testing.T) {
	in := `0x1EC0`
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test", args{in}, 12345334},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hex2bin(tt.args.input); got != tt.want {
				t.Errorf("hex2bin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hex2int64(t *testing.T) {
	in := `0113`
	in2 := `00000B1A29F64B1A`
	in3 := `46E2`

	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test", args{in}, 275},
		{"test imei", args{in2}, 12207001062170},
		{"test crc", args{in3}, 18146},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hex2int64(tt.args.input); got != tt.want {
				t.Errorf("hex2int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binstring2int(t *testing.T) {
	in := `11010100011010001001000101001010`
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test", args{in}, 12345334},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binstring2int(tt.args.input); got != tt.want {
				t.Errorf("binstring2int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrc(t *testing.T) {
	// buff := bufio.NewReader(f)
	// len := buff.Size()
	// fmt.Println(len)
	// body := make([]byte, 1024)
	// content, _ := buff.Read(body)

	fmt.Printf("% q\n", fullPackage)
	c1 := fullPackage[2 : len(fullPackage)-2]
	content, _ := strconv.ParseUint(c1, 16, 64)
	fmt.Println(content)
	// input := Hex2int("0723")
	// bytes := content
	type args struct {
		byteArray []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"crc kermit", args{[]byte(c1)}, 947},
		{"crc kermit-2", args{[]byte(c1)}, 947},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc(tt.args.byteArray); got != tt.want {
				t.Errorf("Crc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexT(t *testing.T) {
	in := "00000B1A29F64B1A"
	in2 := "46E2"
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"imei test", args{in}, 8},
		{"imei test", args{in2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexT(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexT() = %v, want %v", got, tt.want)
			}
		})
	}
}
