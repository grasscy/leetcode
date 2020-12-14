package stack

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"1-11"}, 0},
		{"", args{"1 + 1"}, 2},
		{"", args{" 2-1 + 2 "}, 3},
		{"", args{"(1+(4+5+2)-3)+(6+8)"}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
