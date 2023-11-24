package sorts

import (
	"math/rand"
	"slices"
	"testing"
)

func TestIntroSort(t *testing.T) {
	t.Parallel()
	arr := []myInt{myInt(9), myInt(3), myInt(4), myInt(7), myInt(4), myInt(56), myInt(6), myInt(7), myInt(8), myInt(2), myInt(3), myInt(4), myInt(6), myInt(7)}
	IntroSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("IntroSort don't sort")
	}
}

func TestIntroSortRandomInput_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(rand.Intn(1000))
	}

	IntroSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("IntroSort don't sort big random arrays")
	}
}

func TestIntroSort_0_1M(t *testing.T) {
	t.Parallel()

	arrLength := 1000000
	arr := make([]myInt, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = myInt(i)
	}
	rand.Shuffle(arrLength, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	IntroSort(arr)
	if !slices.IsSorted(arr) {
		t.Fatal("IntroSort don't sort big random arrays")
	}
}
