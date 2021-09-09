package hash

func romanToInt(s string) int {
	var roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev, result := roman[s[0]], roman[s[0]]
	for i := 1; i < len(s); i++ {
		cur := roman[s[i]]
		if cur > prev {
			result = result - 2*prev + cur
		} else {
			result += cur
		}
		prev = cur
	}
	return result
}
