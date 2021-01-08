package stack

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stk := make([]Info, 0, len(logs))
	last := getInfo(logs[len(logs)-1])

	ts := make([]byte, last.T+1)
	for i := range ts {
		ts[i] = 255
	}

	for _, log := range logs {
		info := getInfo(log)
		if info.Type == "end" {
			si := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			for i := si.T; i <= info.T; i++ {
				if ts[i] > 111 {
					ts[i] = info.Id
				}
			}
		} else {
			stk = append(stk, info)
		}

	}
	for _, v := range ts {
		res[v]++
	}
	return res
}

func exclusiveTime1(n int, logs []string) []int {
	res := make([]int, n)
	stk := make([]Info, 0, len(logs))

	for _, log := range logs {
		info := getInfo(log)
		if info.Type == "end" {
			si := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res[info.Id] += info.T - si.T + 1
			if len(stk) > 0 {
				stk[len(stk)-1].T = info.T + 1
			}
		} else {
			if len(stk) > 0 {
				res[stk[len(stk)-1].Id] += info.T - stk[len(stk)-1].T
			}

			stk = append(stk, info)
		}

	}
	return res
}

func getInfo(s string) Info {
	split := strings.Split(s, ":")
	id, _ := strconv.Atoi(split[0])
	t, _ := strconv.Atoi(split[2])

	return Info{
		Id:   byte(id),
		Type: split[1],
		T:    t,
	}
}

type Info struct {
	Id   byte
	Type string
	T    int
}
