package backTracking

func restoreIpAddresses(s string) []string {
	result := []string{}
	var backTracking func(cur, pointNum int)
	backTracking = func(cur, pointNum int) {
		if pointNum == 3 {
			if isVaid(s, cur, len(s)-1) {
				temp := s
				result = append(result, temp)
			}
			return
		}
		for i := cur; i < len(s); i++ {
			if isVaid(s, cur, i) {
				s = string(s[:i+1]) + "." + string(s[i+1:])
				pointNum++
				backTracking(i+2, pointNum)
				pointNum--
				s = string(s[:i+1]) + string(s[i+2:])
			} else {
				break
			}
		}
	}
	backTracking(0, 0)
	return result
}

func isVaid(s string, start, end int) bool {
	if start > end {
		return false
	}
	if s[start]-'0' == 0 && start != end {
		return false
	}
	num := 0
	for i := start; i <= end; i++ {
		if s[i]-'0' > 9 || s[i]-'0' < 0 {
			return false
		}
		num = (num * 10) + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}
	return true
}
