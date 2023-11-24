package sorts

import (
	"github.com/NalbertLeal/cosmos-go/interfaces"
)

func InsertionSort[T interfaces.Order[T]](arr []T) {
	insertionSort(arr, 0, len(arr)-1)
}

func insertionSort[T interfaces.Order[T]](arr []T, start int, end int) {
	for i := start + 1; i <= end; i++ {
		curr := i
		value := arr[i]
		for curr > start && value.CompareTo(arr[curr-1]) < 0 {
			arr[curr] = arr[curr-1]
			curr--
		}
		if curr != i {
			arr[curr] = value
		}
	}
}
