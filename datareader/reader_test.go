package datareader

import "testing"

func TestRead(t *testing.T) {
	in := `030500ad001b18011d666b03410491d08996000056c272013a3e9700`
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{"type1", args{in}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Read(tt.args.data)
		})
	}
}
