package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	HeapSort(arr)
	if !slices.IsSorted(arr) {
		t.Error(arr)
		t.Fatal("HeapSort don't sort")
	}
}

func TestHeapSortRandomInput_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	HeapSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("HeapSort don't sort big random arrays")
	}
}

func TestHeapSort_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	HeapSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("HeapSort don't sort big random arrays")
	}
}
