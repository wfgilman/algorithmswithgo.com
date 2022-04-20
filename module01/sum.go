package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0
	if len(numbers) == 0 || numbers == nil {
		return sum
	}
	for _, n := range numbers {
		sum += n
	}
	return sum
}
