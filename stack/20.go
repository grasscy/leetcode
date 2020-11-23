package main

import "container/list"

func main() {
	ss := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, s := range ss {
		println(isValid(s))
	}
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := list.New()
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.PushBack(c)
		default:
			pop := stack.Back()
			if pop == nil {
				return false
			} else {
				if (c == ')' && pop.Value == '(') || (c == ']' && pop.Value == '[') || (c == '}' && pop.Value == '{') {
					stack.Remove(pop)
				} else {
					return false
				}
			}
		}
	}
	return stack.Len() == 0
}
