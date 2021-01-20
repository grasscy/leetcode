package sort

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}}, 2},
		{"", args{[][]int{{-2147483648, 2147483647}}}, 1},
		{"", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, 2},
		{"", args{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, 4},
		{"", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}}, 2},
		{"", args{[][]int{{1, 2}}}, 1},
		{"", args{[][]int{{2, 3}, {2, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
