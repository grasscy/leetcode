package stack

import (
	"fmt"
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
			index := 0
			for i := len(stkf) - 1; i > 0; i-- {
				if stkf[i] == "(" {
					index = i
				}
			}
			sstkf := stkf[index+1:]
			sstkn := stkn[index:]
			fmt.Println(sstkf)
			fmt.Println(sstkn)

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
