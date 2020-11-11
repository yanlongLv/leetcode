package string

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	next := getNext1([]byte(s))
	length := len(s)
	return next[length-1] != -1 && length%(length-(next[length-1]+1)) == 0
}

func getNext1(t []byte) []int {
	next := make([]int, len(t))
	j := -1
	next[0] = -1
	for i := 1; i < len(t); i++ {
		for j >= 0 && t[i] != t[j+1] {
			j = next[j]
		}
		if t[i] == t[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
