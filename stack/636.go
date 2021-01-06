package stack

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stk := make([]int, 0, len(logs))

	for _, v := range logs {
		split := strings.Split(v, ":")
		id, _ := strconv.Atoi(split[0])
		action := split[1]
		t, _ := strconv.Atoi(split[2])
		if action == "start" {
			stk = append(stk, t)
		} else {
			ts := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res[id] += ts
		}
	}
	return res

}
