package stack

import (
	"strconv"
)

func calculate(s string) int {
	stkn := make([]int, 0, len(s))
	stkf := make([]string, 0, len(s))

	n := ""
	for _, v1 := range s {
		v := string(v1)
		switch v {
		case " ":
		case "(":
			stkf = append(stkf, "(")
		case ")":
			if n != "" {
				b, _ := strconv.Atoi(n)
				stkn = append(stkn, b)
				n = ""
			}
			i := 1
			for ; ; i++ {
				if stkf[len(stkf)-i] == "(" {
					break
				}
			}
			sstkf := stkf[len(stkf)-i+1:]
			stkf = stkf[:len(stkf)-i]
			sstkn := stkn[len(stkn)-i:]
			stkn = stkn[:len(stkn)-i]
			sum := sstkn[0]

			for i, v := range sstkf {
				if v == "+" {
					sum += sstkn[i+1]
				} else {
					sum -= sstkn[i+1]
				}
			}
			stkn = append(stkn, sum)

		case "+":
			if n != "" {
				b, _ := strconv.Atoi(n)
				stkn = append(stkn, b)
				n = ""
			}
			stkf = append(stkf, "+")
		case "-":
			if n != "" {
				b, _ := strconv.Atoi(n)
				stkn = append(stkn, b)
				n = ""
			}
			stkf = append(stkf, v)
		default:
			n += v
		}
	}
	if n != "" {
		b, _ := strconv.Atoi(n)
		stkn = append(stkn, b)
	}
	sum := stkn[0]

	for i, v := range stkf {
		if v == "+" {
			sum += stkn[i+1]
		} else {
			sum -= stkn[i+1]
		}
	}

	return sum
}
