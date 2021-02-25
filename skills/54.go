func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1
	result := []int{}
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		for row := top + 1; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				result = append(result, matrix[bottom][col])
			}
			for row := bottom; row > top; row-- {
				result = append(result, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}