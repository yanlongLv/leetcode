package string

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext([]byte(needle))
	j := -1                              // 因为next数组里记录的起始位置为-1
	for i := 0; i < len(haystack); i++ { // 注意i就从0开始
		// 不匹配
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // j 寻找之前匹配的位置
		}
		// 匹配，j和i同时向后移动
		if haystack[i] == needle[j+1] {
			j++
		}
		// 文本串s里出现了模式串t
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(s []byte) []int {
	j := -1
	next := make([]int, len(s))
	next[0] = j
	for i := 1; i < len(s); i++ { // 注意i从1开始
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 向前回溯
		}
		if s[i] == s[j+1] { // // 找到相同的前后缀
			j++
		}
		next[i] = j // // 将j（前缀的长度）赋给next[i]
	}
	return next
}

func strStr2(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	if l1 == 0 || l1 < l2 {
		return -1
	}
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
