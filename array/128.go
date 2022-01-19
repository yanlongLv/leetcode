package array

func longestConsecutive(nums []int) int {
	var (
		numsMap = make(map[int]int)
		result  = 0
	)

	for _, value := range nums {
		numsMap[value] = 1
	}

	for k, v := range numsMap {
		if v == 0 {
			continue
		}
		numsMap[k] = 0
		r := 1
		for m := 1; m < len(nums); m++ {
			if _, ok := numsMap[k+m]; ok {
				r++
				numsMap[k+m] = 0
			} else {
				break
			}
		}
		if result < r {
			result = r
		}
	}
	return result
}
