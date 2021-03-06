package stack

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := isValid(tt.args); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
