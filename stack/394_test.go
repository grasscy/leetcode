package stack

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"3[a]2[bc]"}, "aaabcbc"},
		{"", args{"3[a2[c]]"}, "accaccacc"},
		{"", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"", args{"abc3[cd]xyz"}, "abccdcdcdxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
