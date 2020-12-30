package string

import "strings"

func generateTheString(n int) string {
	var str strings.Builder
	c := 1
	if n > 1 && n%2 != 0 {
		c = 2
	}
	char := []byte{'a', 'b', 'c'}
	for i := 0; i < c; i++ {
		str.WriteByte(char[i])
	}
	for i := 0; i < n-c; i++ {
		str.WriteByte(char[c])
	}
	return str.String()
}
