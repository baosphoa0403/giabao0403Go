package array

import (
	"testing"
)

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := Filter(data, func(item int) bool {
		return item%2 == 0
	})

	if !AreSlicesEqual(result, []int{2, 4}) {
		t.Errorf("error")
	}
}
