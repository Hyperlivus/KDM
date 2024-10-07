package matf

func Factorial(x int) int {
	res := 1
	for i := x; i > 1; i-- {
		res *= i
	}
	return res
}
func C(n int, k int) int {
	if n < k {
		panic("k must be less than n")
	}
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
func Pow(x int, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}
func Log(a int, b int) int {
	res := 0

	curr := b
	for curr != 0 {
		o := curr % a
		curr -= o
		curr /= a
		res++
	}

	return res
}
