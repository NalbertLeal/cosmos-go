package searchs

import "testing"

func TestBinSearch(t *testing.T) {
	arr := []myInt{myInt(0), myInt(19), myInt(23), myInt(32), myInt(49), myInt(53), myInt(63), myInt(78), myInt(82), myInt(96)}
	if BinSearch(arr, 32) != 3 {
		t.Fatal("BinSearch can't find the correct position")
	}
}