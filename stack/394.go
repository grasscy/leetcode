package stack

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stk := make([]string, 0, len(s))
	for _, v1 := range s {
		v := string(v1)
		switch v {
		case "]":
			s := ""
			for stk[len(stk)-1] != "[" {
				s = stk[len(stk)-1] + s
				stk = stk[:len(stk)-1]
			}
			stk = stk[:len(stk)-1]

			k := ""
			for len(stk) != 0 {
				_, err := strconv.Atoi(stk[len(stk)-1])
				if err == nil {
					k = stk[len(stk)-1] + k
					stk = stk[:len(stk)-1]
				} else {
					break
				}
			}

			r, _ := strconv.Atoi(k)
			tr := ""
			for r != 0 {
				tr += s
				r--
			}
			stk = append(stk, tr)
		default:
			stk = append(stk, v)
		}
	}
	return strings.Join(stk, "")
}
