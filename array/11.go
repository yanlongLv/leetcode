package array

func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > res {
			res = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// func min(x, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }

// func maxArea(height []int) int {
// 	//O(n) 双指针解法。核心思想：当前的最大体积，是由边界中最小值决定的,所以优先动最小的边
// 	var res int
// 	left := 0
// 	right := len(height) - 1
// 	for left < right {
// 		v := (right - left) * min(height[left], height[right])
// 		if v > res {
// 			res = v
// 		}
// 		if height[left] < height[right] {
// 			left++
// 		} else {
// 			right--
// 		}
// 	}
// 	return res
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	} else {
// 		return a
// 	}
// }
