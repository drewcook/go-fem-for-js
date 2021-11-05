package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 15
	actual := Add(1, 2, 3, 4, 5)
	if expected != actual {
		t.Errorf("Sum was incorrect! Expected %d but got %d.", expected, actual)
	}
}
