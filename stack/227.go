package stack

import "strconv"

func calculate227(s string) int {
	stkn := make([]int, 0, len(s))
	n := ""
	f := '+'
	for _, v1 := range s {
		if v1 == ' ' {
			continue
		}
		if '0' <= v1 && v1 <= '9' {
			n += string(v1)
			continue
		}
		switch f {
		case '-':
			atoi, _ := strconv.Atoi(n)
			stkn = append(stkn, -atoi)
		case '+':
			atoi, _ := strconv.Atoi(n)
			stkn = append(stkn, atoi)
		case '*':
			atoi, _ := strconv.Atoi(n)
			i := stkn[len(stkn)-1]
			stkn[len(stkn)-1] = i * atoi
		case '/':
			atoi, _ := strconv.Atoi(n)
			i := stkn[len(stkn)-1]
			stkn[len(stkn)-1] = i / atoi
		}
		f = v1
		n = ""
	}

	if len(n) > 0 {
		atoi, _ := strconv.Atoi(n)
		if f == '*' {
			i := stkn[len(stkn)-1]
			stkn[len(stkn)-1] = i * atoi
		} else if f == '/' {
			i := stkn[len(stkn)-1]
			stkn[len(stkn)-1] = i / atoi
		} else if f == '-' {
			stkn = append(stkn, -atoi)
		} else {
			stkn = append(stkn, atoi)
		}
	}

	sum := 0
	for _, v := range stkn {
		sum += v
	}
	return sum
}
