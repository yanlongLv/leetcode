package divideConpquer

func majorityElement(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}
func majorityElementRec(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	mid := (start + end) / 2
	left := majorityElementRec(nums, start, mid)
	right := majorityElementRec(nums, mid+1, end)
	if countMajority(nums, left, start, end) > (end-start+1)/2 {
		return left
	}
	if countMajority(nums, right, start, end) > (end-start+1)/2 {
		return right
	}
	return -1
}
func countMajority(nums []int, target, start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}
