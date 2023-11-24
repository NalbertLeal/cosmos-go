package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func HeapSort[T interfaces.Order[T]](arr []T) {
	heapSort(arr, 0, len(arr)-1)
}

func heapSort[T interfaces.Order[T]](arr []T, start int, end int) {
	length := end - start
	for i := ((length / 2) - 1) + start; i >= start; i-- {
		heapfy(arr, start, end, i)
	}
	for i := end; i > start; i-- {
		arr[i], arr[start] = arr[start], arr[i]
		heapfy(arr, start, i, start)
	}
}

func heapfy[T interfaces.Order[T]](arr []T, start int, end int, initRoot int) {
	root := initRoot
	largest := initRoot

	for {
		child := (root - start) * 2

		if child+1+start < end && arr[child+1+start].CompareTo(arr[largest]) > 0 {
			largest = child + 1 + start
		}
		if child+2+start < end && arr[child+2+start].CompareTo(arr[largest]) > 0 {
			largest = child + 2 + start
		}
		if largest == root {
			break
		}
		arr[root], arr[largest] = arr[largest], arr[root]
		root = largest
	}
}
