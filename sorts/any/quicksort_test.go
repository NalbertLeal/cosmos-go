package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQuicksort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	Quicksort(arr)
	if !slices.IsSorted(arr) {
		t.Error("Quicksort don't sort")
		t.Fatal(arr)
	}
}

func TestQuicksortRandomInput_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(100))
	}

	Quicksort(arr)
	if !slices.IsSorted(arr) {
		t.Error(arr)
		t.Fatal("Quicksort don't sort big random arrays")
	}
}

func TestQuicksort_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	Quicksort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("Quicksort don't sort big random arrays")
	}
}
