package backTracking

func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}
	var partString func(cur int)
	partString = func(cur int) {
		if cur >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := cur; i < len(s); i++ {
			if isPalindrome(s, cur, i) {
				str := s[cur : i+1]
				path = append(path, str)
			} else {
				continue
			}
			partString(i + 1)
			path = path[:len(path)-1]
		}
	}
	partString(0)
	return result
}

func isPalindrome(s string, start, end int) bool {
	j := end
	i := start
	for i < j {
		if s[i] != s[j] {
			return false
		}
		j--
		i++
	}
	return true
}
