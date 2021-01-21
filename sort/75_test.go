package sort

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{[]int{2, 0, 2, 1, 1, 0}}},
		{"", args{[]int{2, 0, 1}}},
		{"", args{[]int{0}}},
		{"", args{[]int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
