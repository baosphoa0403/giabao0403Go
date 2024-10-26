package math

import "testing"

func TestAdd(t *testing.T) {
	result := Sum(3, 4)
	if result != 7 {
		t.Errorf("error")
	}
}
