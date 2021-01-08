package stack

import (
	"go/ast"
	"strconv"
)

func countOfAtoms(formula string) string {
	stk := make([]string, 0, len(formula))
	var m map[string]int

	for i := 0; i < len(formula); i++ {
		v := formula[i]
		if 'A' <= v && 'Z' >= v {
			ele := string(v)
			if i+1 < len(formula) && 'a' <= formula[i+1] && 'z' >= formula[i+1] {
				ele += string(formula[i+1])
				i++
			}
			stk = append(stk, ele)
			continue
		} else if '(' == v || ')' == v {
			stk = append(stk, string(v))
			continue
		} else {
			n := string(v)
			for i+1 < len(formula) && '0' <= formula[i+1] && '9' >= formula[i+1] {
				n += string(i + 1)
				i++
			}
			if stk[len(stk)-1] != ")" {
				ele := stk[len(stk)-1]
				stk = stk[:len(stk)-1]

			}
		}
	}

}
