package stack

import (
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	stk := make([]string, 0, len(formula))

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
				n += string(formula[i+1])
				i++
			}
			nn, _ := strconv.Atoi(n)

			if stk[len(stk)-1] != ")" {
				stk = append(stk, n)
			} else {
				stk = stk[:len(stk)-1]

				var ss []string
				for stk[len(stk)-1] != "(" {
					ss = append(ss, stk[len(stk)-1])
					stk = stk[:len(stk)-1]
				}
				stk = stk[:len(stk)-1]

				n1 := 1
				for _, v := range ss {
					atoi, err := strconv.Atoi(v)
					if err != nil {
						stk = append(stk, v)
						stk = append(stk, strconv.Itoa(nn*n1))
						n1 = 1
					} else {
						n1 = atoi
					}
				}
			}
		}
	}

	m := make(map[string]int)
	var eles []string
	prex := ""
	for _, v := range stk {
		if v == "(" || v == ")" {
			continue
		}

		atoi, err := strconv.Atoi(v)
		if err != nil {
			_, ok := m[v]
			if ok {
				m[v] += 1
			} else {
				eles = append(eles, v)
				m[v] += 1
			}
		} else {
			m[prex] += atoi - 1
		}
		prex = v
	}

	sort.Strings(eles)

	res := ""
	for _, v := range eles {
		res += v
		if m[v] > 1 {
			res += strconv.Itoa(m[v])
		}
	}

	return res
}
