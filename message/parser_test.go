package message

import (
	"reflect"
	"testing"
)

func Test_msgParser(t *testing.T) {
	data := `4568004009ce40007806c340727fde0367097db53a98272d1094abbe00000000b0022a80533900000204055001030300010104020101080a0003654f00000000`
	msg := new(Message)
	msg.PacketLength = 100
	msg.Imei = 898989787979
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"incoming", args{data}, msg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := msgParser(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("msgParser() = %+v, want %v", got, tt.want)
			}
		})
	}
}
