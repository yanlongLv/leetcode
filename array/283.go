func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if mums[i] != 0 {
			mums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}