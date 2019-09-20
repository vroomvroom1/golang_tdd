package arrayslices

// Sum accepts and array of 5 integers and returns the sum
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
