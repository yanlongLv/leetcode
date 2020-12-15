package greed

func candy(ratings []int) int {
	candyVec := make([]int, len(ratings))
	candyVec[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candyVec[i] = candyVec[i-1] + 1
		} else {
			candyVec[i] = 1
		}

	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j+1] < ratings[j] {
			candyVec[j] = maxNumber(candyVec[j+1]+1, candyVec[j])
		}
	}
	result := 0
	for k := 0; k < len(candyVec); k++ {
		result += candyVec[k]
	}
	return result
}

func maxNumber(x, y int) int {
	if x > y {
		return x
	}
	return y
}
