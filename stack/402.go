package stack

import "strings"

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	res := make([]uint8, 0, len(num))
	for i := 0; i < len(num); i++ {
		for len(res) > 0 && res[len(res)-1] > num[i] && k > 0 {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, num[i])
	}

	if k > 0 {
		ii := 0
		for i := len(res) - 1; i > 0; i-- {
			if res[i] == '0' {
				ii++
			} else {
				break
			}
		}

		res = res[:len(res)-k-ii]
	}

	index := -1
	for i, v := range res {
		if v != '0' {
			index = i
			break
		}
	}
	if index < 0 {
		return "0"
	}
	if index > 0 {
		res = res[index:]
	}
	return string(res)
}

func removeKdigits1(num string, k int) string {

	remain := len(num) - k

	res := make([]uint8, 0, len(num))
	for i := 0; i < len(num); i++ {
		for len(res) > 0 && res[len(res)-1] > num[i] && k > 0 {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, num[i])
	}
	s := string(res[:remain])
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		return "0"
	}
	return s
}
