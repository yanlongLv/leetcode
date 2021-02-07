package recursive

func generateParenthesis(n int) []string {
	result := []string{}
	var generate func(left, right int, s string)
	generate = func(left, right int, s string) {
		if left == n && right == n {
			result = append(result, s)
		}
		if left < n {
			generate(left+1, right, s+"(")
		}
		if left > right {
			generate(left, right+1, s+")")
		}
	}
	generate(0, 0, "")
	return result
}
