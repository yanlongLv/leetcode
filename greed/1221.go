package greed

func balancedStringSplit(s string) int {
	l := 0
	r := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			l++
		} else {
			r++
		}
		if l == r {
			count++
		}
	}
	return count
}
