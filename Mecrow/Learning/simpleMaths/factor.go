package simpleMaths

func IsPrime(input int) bool {
	if input < 2 {
		return false
	}
	result := true
	for index := 2; index*index <= input; index++ {
		if input%index == 0 {
			result = false
			break
		}
	}
	return result
}

func getNextPrime(current_prime int) int {
	next_prime := current_prime + 1
	for !IsPrime(next_prime) {
		next_prime++
	}
	return next_prime
}

func PrimeFactors(input int) []int {
	factors := make([]int, 0, 64)
	n := 2
	for input > 1 {
		if input%n == 0 {
			factors = append(factors, n)
			input /= n
			n = 2
		} else {
			n = getNextPrime(n)
		}
	}
	return factors
}
