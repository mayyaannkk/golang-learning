package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 2}
		want := Sum(numbers)
		got := 6
		if got != want {
			t.Errorf("want %d, got %d, and given %v", want, got, numbers)
		}
	})
}
