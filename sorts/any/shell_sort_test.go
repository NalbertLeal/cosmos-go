package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestShellSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	ShellSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("ShellSort don't sort")
	}
}

func TestShellSortRandomInput_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	ShellSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("ShellSort don't sort big random arrays")
	}
}

func TestShellSort_0_100k(t *testing.T) {
	t.Parallel()

	arrLength := 100000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	ShellSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("ShellSort don't sort big random arrays")
	}
}
