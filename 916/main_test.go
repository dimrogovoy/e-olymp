package main

import "testing"

func Test_realMain(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{1, 10, 1, 10}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := realMain(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("realMain() = %v, want %v", got, tt.want)
			}
		})
	}
}
