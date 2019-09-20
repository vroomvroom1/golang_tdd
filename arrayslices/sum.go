package arrayslices

// Sum accepts and array of 5 integers and returns the sum
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes in a number of slices and returns a slice with the sum
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns a slice with each tail
func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
