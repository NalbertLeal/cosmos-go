package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestBubleSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	BubleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("BubleSort don't sort")
	}
}

func TestBubleSortRandomInput_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	BubleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("BubleSort don't sort big random arrays")
	}
}

func TestBubleSort_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	BubleSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("BubleSort don't sort big random arrays")
	}
}
