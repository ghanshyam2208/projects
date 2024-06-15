package arrays

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {

	t.Run("Sum of tails", func(t *testing.T) {
		actual := SumAllTails([]int{2, 3}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v expected %v", actual, expected)
		}
	})

	t.Run("Sum of empty tails", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v expected %v", actual, expected)
		}
	})
}
