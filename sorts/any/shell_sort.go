package sorts

import (
	"github.com/NalbertLeal/cosmos-go/interfaces"
)

func ShellSort[T interfaces.Order[T]](arr []T) {
	shellSort(arr, 0, len(arr)-1)
}

func shellSort[T interfaces.Order[T]](arr []T, start int, end int) {
	for gap := end - start; gap > 0; gap /= 2 {
		for i := start + gap; i <= end; i += gap {
			curr := i
			value := arr[i]
			for curr > start && value.CompareTo(arr[curr-gap]) < 0 {
				arr[curr] = arr[curr-gap]
				curr -= gap
			}
			if curr != i {
				arr[curr] = value
			}
		}
	}
}
