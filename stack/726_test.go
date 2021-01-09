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
		{"", args{"(H)"}, "H"},
		{"", args{"K4(ON(SO3)2)2"}, "K4N2O14S4"},
		{"", args{"Mg(OH)2"}, "H2MgO2"},
		{"", args{"Be32"}, "Be32"},
		{"", args{"H2O"}, "H2O"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}
