package arrays

func Sum(nums [5]int) int {
	// Sum function
	var total int
	for i := 0; i < 5; i++ {
		total += nums[i]
	}
	return total
}
