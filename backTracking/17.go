package backTracking

func letterCombinations(digits string) []string {
	cache := [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := []string{}
	var path string
	var backStacking func(digitsIndex int)
	backStacking = func(digitsIndex int) {
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}
		letter := cache[digits[digitsIndex]-'0']
		for i := 0; i < len(letter); i++ {
			path += string(letter[i])
			// 注意细微的差别 这次是digitsIndex
			backStacking(digitsIndex + 1)
			path = path[:len(path)-1]
		}
	}
	backStacking(0)
	return result
}
