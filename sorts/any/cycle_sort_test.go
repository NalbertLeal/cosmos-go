package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestCycleSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	CycleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("CycleSort don't sort")
	}
}

func TestCycleSortRandomInput_0_10k(t *testing.T) {
	t.Parallel()

	arrLength := 10000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	CycleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("CycleSort don't sort big random arrays")
	}
}

func TestCycleSort_0_10k(t *testing.T) {
	t.Parallel()

	arrLength := 10000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	CycleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("CycleSort don't sort big random arrays")
	}
}
