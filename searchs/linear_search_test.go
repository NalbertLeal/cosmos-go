package searchs

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []myInt{myInt(0), myInt(19), myInt(23), myInt(32), myInt(49), myInt(53), myInt(63), myInt(78), myInt(82), myInt(96)}
	if LinearSearch(arr, 32) != 3 {
		t.Fatal("LinearSearch can't find the correct position")
	}
}