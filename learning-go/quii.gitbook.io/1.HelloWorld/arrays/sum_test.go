package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test array with len 5", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		actual := Sum(nums)
		expected := 15

		if actual != expected {
			t.Errorf("Expected %d, but got %d", expected, actual)
		}
	})

	t.Run("test array with any len", func(t *testing.T) {
		nums := []int{1, 2, 3}

		actual := Sum(nums)
		expected := 6

		if actual != expected {
			t.Errorf("Expected %d, but got %d", expected, actual)
		}
	})

}
