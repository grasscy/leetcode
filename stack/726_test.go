package stack

import "testing"

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"H2O"}, "H2O"},
		{"", args{"Mg(OH)2"}, "H2MgO2"},
		{"", args{"K4(ON(SO3)2)2"}, "K4N2O14S4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}
