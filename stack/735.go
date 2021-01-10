package stack

func asteroidCollision(asteroids []int) []int {
	stk := make([]int, 0, len(asteroids))

	for _, v := range asteroids {
		len := len(stk)
		if len == 0 {
			stk = append(stk, v)
			continue
		}

		ifIn := true
		for len > 0 && stk[len-1] > 0 && v < 0 {
			pabs := stk[len-1]

			if pabs < -v {
				stk = stk[:len-1]
				len--
			} else if pabs > -v {
				ifIn = false
				break
			} else {
				ifIn = false
				stk = stk[:len-1]
				break
			}
		}

		if ifIn {
			stk = append(stk, v)
		}
	}
	return stk
}
