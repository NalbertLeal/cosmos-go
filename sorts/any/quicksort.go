package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func Quicksort[T interfaces.Order[T]](arr []T) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort[T interfaces.Order[T]](arr []T, start int, end int) {
	if start < end {
		sliceIndex := hoaresPartition(arr, start, end)
		quicksort(arr, start, sliceIndex-1)
		quicksort(arr, sliceIndex+1, end)

		// low, high := partition2Pivots(arr, start, end)
		// quicksort(arr, start, low-1)
		// quicksort(arr, low+1, high-1)
		// quicksort(arr, high+1, end)

		// left, right := nalbert3WayPartition(arr, start, end)
		// quicksort(arr, start, right)
		// quicksort(arr, left, end)

	}
}
