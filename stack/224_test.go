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
		{"", args{"2-4-(8+2-6+(8+4-(1)+8-10))"}, -15},
		{"", args{"1-11"}, -10},
		{"", args{"2-(5-6)"}, 3},
		{"", args{" 2-1 + 2 "}, 3},
		{"", args{"(1+(4+5+2)-3)+(6+8)"}, 23},
		{"", args{"(3-(2-(5-(9-(4)))))"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
