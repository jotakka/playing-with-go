package exercism

const (
	Udefined = 0
	NotPrime = 1
	Prime    = 2
)

func Sieve(limit int) []int {
	if limit < 2 {
		return make([]int, 0)
	}
	numbers := make([]int, limit+1)

	numbers[0] = NotPrime
	numbers[1] = NotPrime
	primeCount := 0

	for idx, _ := range numbers {
		if numbers[idx] == NotPrime {
			continue
		}

		numbers[idx] = Prime
		primeCount++

		for i := 2; i*idx <= limit; i++ {
			numbers[idx*i] = NotPrime
		}
	}

	primes := make([]int, primeCount)
	j := 0
	for idx, val := range numbers {
		if val == Prime {
			primes[j] = idx
			j++
		}
	}
	return primes
}
