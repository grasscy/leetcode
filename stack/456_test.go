package stack

import "testing"

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{-1, 3, 2, 0}}, true},
		{"", args{[]int{1, 0, 1, -4, -3}}, false},
		{"", args{[]int{3, 1, 4, 2}}, true},
		{"", args{[]int{3, 5, 0, 3, 4}}, true},
		{"", args{[]int{1, 2, 3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern1(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
