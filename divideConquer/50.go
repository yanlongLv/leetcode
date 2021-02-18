package divideConpquer

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0/divideConpquer(x, -n)
	}
	return divideConpquer(x,n)
 }
 
 func divideConpquer(x float64, n int) float64 {
	 if n == 0 {
		 return 1.0
	 }
	 y := divideConpquer(x, n/2)
	 if  n % 2 == 1{
		 return y * y * x
	 }
	 return y * y
 }