package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	InsertionSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("InsertionSort don't sort")
	}
}

func TestInsertionSortRandomInput_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	InsertionSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("InsertionSort don't sort big random arrays")
	}
}

func TestInsertionSort_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	InsertionSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("InsertionSort don't sort big random arrays")
	}
}

// func TestFastInsertionSortRandomInput_0_100k(t *testing.T) {
// 	t.Parallel()

// 	arrLength := 100000
// 	arr := make([]myInt, arrLength)
// 	for i := 0; i < arrLength; i++ {
// 		arr[i] = myInt(rand.Intn(1000))
// 	}

// 	FastInsertionSort(arr)
// 	if !slices.IsSorted(arr) {
// 		t.Fatal("FastInsertionSort don't sort big random arrays")
// 	}
// }
