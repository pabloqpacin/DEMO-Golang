package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {sum += numbers[i]}
	for _, number := range numbers {
		// On each iteration, range returns two values - the index and the value.
		// We are choosing to ignore the index value by using _ (blank identifier).
		sum += number
	}
	return sum
}
