package sorts

import (
	"github.com/NalbertLeal/cosmos-go/interfaces"
)

func IntroSort[T interfaces.Order[T]](arr []T) {
	depth := introsortDepth(len(arr))
	introSort(arr, 0, len(arr)-1, depth)
}

func introSort[T interfaces.Order[T]](arr []T, start int, end int, deepth int) {
	if end-start < 12 {
		insertionSort(arr, start, end)
		return
	}
	if deepth == 0 {
		heapSort(arr, start, end)
		return
	}

	left, right := nalbert3WayPartition(arr, start, end)
	introSort(arr, start, right, deepth-1)
	introSort(arr, left, end, deepth-1)
}

func introsortDepth(arrLength int) int {
	maxDepth := 0
	for i := arrLength; i > 0; i >>= 1 {
		maxDepth++
	}
	return maxDepth * 2
}
