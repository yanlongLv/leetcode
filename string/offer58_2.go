package string

func reverseLeftWords(s string, n int) string {
	t := []byte(s)
	var reverseWords func(start, end int)
	reverseWords = func(start, end int) {
		i, j := start, end
		for j > i {
			t[j], t[i] = t[i], t[j]
			i++
			j--
		}
	}
	reverseWords(0, n-1)
	reverseWords(n, len(s)-1)
	reverseWords(0, len(s)-1)
	return string(t)
}
