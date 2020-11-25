package greed

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	preDiff := 0
	curDiff := 0
	result := 1
	for i := 1; i < len(nums); i++ {
		curDiff = nums[i] - nums[i-1]
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0) {
			result++
			preDiff = curDiff
		}
	}
	return result
}
