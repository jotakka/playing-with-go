package exercism

func SquareOfSum(n int) int {
	nVal := n * (n + 1) / 2

	return nVal * nVal
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
