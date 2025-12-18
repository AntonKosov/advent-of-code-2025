package math

func CountBits[T Numbers](value T) int {
	count := 0
	for value != 0 {
		count++
		value &= value - 1
	}

	return count
}
