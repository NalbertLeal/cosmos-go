package sorts

import (
	"github.com/NalbertLeal/cosmos-go/interfaces"
)

func Mergesort[T interfaces.Order[T]](arr []T) {
	mergesort(arr, 0, len(arr)-1)
}

func mergesort[T interfaces.Order[T]](arr []T, start int, end int) {
	if start < end {
		middle := (start + end) / 2

		mergesort(arr, start, middle)
		mergesort(arr, middle+1, end)

		merge(arr, start, middle, end)
	}
}
