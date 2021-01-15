package sort

func isAnagram(s string, t string) bool {
	m := make(map[int32]int)

	for _, s := range s {
		i, ok := m[s]
		if !ok {
			m[s] = 1
			continue
		}
		m[s] = i + 1
	}

	for _, s := range t {
		i, ok := m[s]
		if !ok {
			return false
		}
		if i < 1 {
			return false
		}
		m[s] = m[s] - 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true

}

func isAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
