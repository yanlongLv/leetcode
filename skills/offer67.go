func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	bytes := []byte(str)
	i := 0
	ch := bytes[i]
	for i < len(bytes)-1 && ch == ' ' {
		i++
		ch = bytes[i]
	}
	bytes = bytes[i:]
	if len(bytes) == 0 {
		return 0
	}
	na := false
	if bytes[0] == '-' {
		na = true
		bytes = bytes[1:]
	} else if bytes[0] == '+' {
		bytes = bytes[1:]
	}
	n := 0
	for _, ch := range bytes {
		ch -= '0'
		if ch > 9 {
			break
		}
		if a, ok := checkN(n, na); !ok {
			return a
		}
		n = n*10 + int(ch)
	}
	if a, ok := checkN(n, na); !ok {
		return a
	}
	if na {
		n = -n
	}
	return n

}

func isGtInt32(n int) bool {
	if n > (1<<31)-1 {
		return true
	}
	return false
}

func checkN(n int, na bool) (int, bool) {
	if isGtInt32(n) {
		if na {
			return math.MinInt32, false
		}
		return (1 << 31) - 1, false
	}
	return 0, true
}