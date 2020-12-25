package stack

import "testing"

func Test_calculate227(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"2*3+4"}, 10},
		{"", args{"0-2147483647"}, -2147483647},
		{"", args{"3+2*2"}, 7},
		{"", args{" 3/2 "}, 1},
		{"", args{" 3+5 / 2 "}, 5},
		{"", args{"1-1+1"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate227(tt.args.s); got != tt.want {
				t.Errorf("calculate227() = %v, want %v", got, tt.want)
			}
		})
	}
}
