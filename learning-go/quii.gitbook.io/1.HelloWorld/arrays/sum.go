package arrays

func Sum(nums []int) int {
	// Sum function
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
