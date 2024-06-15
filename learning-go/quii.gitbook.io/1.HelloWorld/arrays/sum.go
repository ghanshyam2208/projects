package arrays

func Sum(nums []int) int {
	// Sum function
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAllTails(numbersToSum ...[]int) []int {
	// var sums []int
	sums := make([]int, len(numbersToSum))
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, tail...)
		}
	}
	return sums
}
