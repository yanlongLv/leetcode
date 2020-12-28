package greed

import "strconv"

func monotoneIncreasingDigits(N int) int {
	ns := []byte(strconv.Itoa(N))
	flag := len(ns)
	for i := len(ns) - 1; i > 0; i-- {
		if ns[i-1] > ns[i] {
			flag = i
			ns[i-1]--

		}
	}
	for j := flag; j < len(ns); j++ {
		ns[j] = '9'
	}
	nj, _ := strconv.Atoi(string(ns))
	return nj
}
