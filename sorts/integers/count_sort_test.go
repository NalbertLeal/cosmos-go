package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestCountSort(t *testing.T) {
	t.Parallel()
	arr := []int{9, 3, 4, 7, 4, 56, 6, 7, 8, 2, 3, 4, 6, 7}
	CountSort(arr, 56, 2)
	if !slices.IsSorted(arr) {
		t.Fatal("CountSort don't sort")
	}
}

func TestCountSortRandomInput_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = rand.Intn(1000)
	}

	max := -1000000
	min := 1000000
	for i := 0; i < len(arr); i += 1 {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}

	CountSort(arr, max, min)
	if !slices.IsSorted(arr) {
		t.Fatal("CountSort don't sort big random arrays")
	}
}

func TestCountSort_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = i
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	max := -1000000
	min := 1000000
	for i := 0; i < len(arr); i += 1 {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}

	CountSort(arr, max, min)
	if !slices.IsSorted(arr) {
		t.Fatal("HeapSort don't sort big random arrays")
	}
}
