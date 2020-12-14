package stack

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, " ", "")
	if !strings.Contains(s, "+") && !strings.Contains(s, "-") {
		a, _ := strconv.Atoi(s)
		return a
	}

	stkn := make([]int, 0, len(s))
	stkf := make([]string, 0, len(s))
	n := ""
	for _, v1 := range s {
		v := string(v1)
		switch v {
		case "+", "-":
			stkf = append(stkf, v)
			if len(n) > 0 {
				stkn = append(strconv.Atoi(n))
			}
		default:
			n += v
			atoi, _ := strconv.Atoi(v)
			if len(stkf) > 0 {
				f := stkf[len(stkf)-1]
				stkf = stkf[:len(stkf)-1]
				if f == "+" {
					a := stkn[len(stkn)-1]
					stkn = stkn[:len(stkn)-1]
					c := a + atoi
					stkn = append(stkn, c)
				} else {
					a := stkn[len(stkn)-1]
					stkn = stkn[:len(stkn)-1]
					c := a - atoi
					stkn = append(stkn, c)
				}
			} else {
				stkn = append(stkn, atoi)
			}
		}
	}
	return stkn[0]
}
