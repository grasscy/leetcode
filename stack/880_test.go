package stack

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{"", args{"y959q969u3hb22odq595222280369", 222280369}, "o"},
		{"", args{"leet2code3", 10}, "o"},
		{"", args{"ha22", 5}, "h"},
		{"", args{"a2345678999999999999999", 1}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeAtIndex(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("decodeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
