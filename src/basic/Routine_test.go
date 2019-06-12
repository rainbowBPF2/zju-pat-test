package basic

import "testing"

func Test_setData(t *testing.T) {
	type args struct {
		data int
		ch   chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setData(tt.args.data, tt.args.ch)
		})
	}
}
