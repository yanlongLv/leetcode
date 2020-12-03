package greed

func minCostToMoveChips(position []int) int {
	nums1, num2 := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			num2++
			continue
		}
		nums1++
	}
	if nums1 > num2 {
		return num2
	}
	return nums1
}
