package string

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(s); i += (2 * k) {
		if i+k < len(s) {
			b = reverse(b, i, i+k-1)
			continue
		}
		b = reverse(b, i, len(s)-1)
	}
	return string(b)
}

func reverse(s []byte, start, end int) []byte {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return s
}
