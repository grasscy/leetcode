package stack

import "regexp"

func isValid591(code string) bool {
	if len(code) < 7 || code[0] != '<' {
		return false
	}

	stk := append([]string{""}, make([]string, 0, len(code))...)

	for i := 0; i < len(code); i++ {
		v := code[i]
		pop := stk[len(stk)-1]

		if v == '>' {
			if pop == "<![CDATA[" && code[i-1] == ']' && code[i-2] == ']' {
				stk = stk[:len(stk)-1]
				continue
			}
		} else if v == '<' {
			if pop == "<![CDATA[" {
				continue
			}
			if i+len("<![CDATA[") < len(code) && code[i:i+len("<![CDATA[")] == "<![CDATA[" {
				if len(stk) == 1 {
					return false
				}
				stk = append(stk, "<![CDATA[")
				i += len("<![CDATA[")
				continue
			}

			tag := ""
			j := i
			for ; j < len(code); j++ {
				tag += string(code[j])
				if code[j] == '>' {
					break
				}
			}
			if tag[1] == '/' {
				if pop == "" || pop[1:] != tag[2:] {
					return false
				}
				stk = stk[:len(stk)-1]
				if len(stk) == 1 && j != len(code)-1 {
					return false
				}
				continue
			} else {
				valid := regexp.MustCompile(`^<[A-Z]{1,9}>$`)
				if !valid.MatchString(tag) {
					return false
				}
				stk = append(stk, tag)
				continue
			}
			i = j
		}
	}
	return len(stk) == 1
}
