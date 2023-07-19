package exercism

import "math"

func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	if n == 1 {
		return factors
	}

	val := n
	for ; val%2 == 0; val /= 2 {
		factors = append(factors, int64(2))
	}

	nSqrt := int64(math.Sqrt(float64(val)))

	primes := Sieve(int(nSqrt))

	for _, prime := range primes {
		if val == 1 {
			break
		}
		for ; val%int64(prime) == 0; val /= int64(prime) {
			factors = append(factors, int64(prime))
		}
	}

	if val == 1 {
		return factors
	} else {
		return append(factors, val)
	}
}
