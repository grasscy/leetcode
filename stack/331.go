package stack

import "strings"

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	split := strings.Split(preorder, ",")
	stk := make([]string, 0, len(split))
	for _, v := range split {
		if v != "#" {
			stk = append(stk, v)
		} else {
			stk = append(stk, v)
			l := len(stk)
			for l > 1 && stk[l-1] == "#" && stk[l-2] == "#" {
				stk = stk[:l-2]
				if l-2 == 0 {
					return false
				}
				stk[l-3] = "#"
				l = l - 2
			}
		}
	}
	return len(stk) == 1 && stk[0] == "#"
}
