package sort

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{4, 2, 5, 7}}, []int{4, 5, 2, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
