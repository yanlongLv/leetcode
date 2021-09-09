package hash

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
