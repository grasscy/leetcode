package stack

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"((((((())))()())))"}, 80},
		{"", args{"(()(()))"}, 6},
		{"", args{"()"}, 1},
		{"", args{"(())"}, 2},
		{"", args{"()()"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.args.S); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
