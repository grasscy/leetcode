package stack

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	type args struct {
		n    int
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}}, []int{3, 4}},
		{"", args{8, []string{"0:start:0", "1:start:5", "2:start:6", "3:start:9", "4:start:11", "5:start:12", "6:start:14", "7:start:15", "1:start:24", "1:end:29", "7:end:34", "6:end:37", "5:end:39", "4:end:40", "3:end:45", "0:start:49", "0:end:54", "5:start:55", "5:end:59", "4:start:63", "4:end:66", "2:start:69", "2:end:70", "2:start:74", "6:start:78", "0:start:79", "0:end:80", "6:end:85", "1:start:89", "1:end:93", "2:end:96", "2:end:100", "1:end:102", "2:start:105", "2:end:109", "0:end:114"}},
			[]int{20, 14, 35, 7, 6, 9, 10, 14}},
		{"", args{1, []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"}}, []int{6}},
		{"", args{1, []string{"0:start:0", "0:end:100000000"}}, []int{100000001}},
		{"", args{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}}, []int{8, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime1(tt.args.n, tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
