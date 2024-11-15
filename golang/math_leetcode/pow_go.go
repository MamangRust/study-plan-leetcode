package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	return quickPow(x, n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := quickPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
