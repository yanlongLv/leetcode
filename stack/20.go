package stack

func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "{" || string(s[i]) == "(" || string(s[i]) == "[" {
			stack = append(stack, string(s[i]))
			continue
		}
		if string(s[i]) == "}" && len(stack) > 0 && stack[len(stack)-1] == "{" {
			stack = stack[:len(stack)-1]
		} else if string(s[i]) == ")" && len(stack) > 0 && stack[len(stack)-1] == "(" {
			stack = stack[:len(stack)-1]
		} else if string(s[i]) == "]" && len(stack) > 0 && stack[len(stack)-1] == "[" {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
