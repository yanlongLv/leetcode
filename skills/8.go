package skills

import "math"

func myAtoi(str string) int {
	n := len(str)
	var i, j int
	neg := false
	for i = 0; i < n; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			break
		} else if str[i] == '+' {
			i++
			break
		} else if str[i] == '-' {
			neg = true
			i++
			break
		} else if str[i] != ' ' {
			return 0
		}
	}
	for j = i; j < n; j++ {
		if str[j] < '0' || str[j] > '9' {
			break
		}
	}
	ret := 0
	for k := i; k < j; k++ {
		cur := int(str[k] - '0')
		if !neg {
			ret = ret*10 + cur
			if ret > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			ret = ret*10 - cur
			if ret < math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}
	return ret
}
