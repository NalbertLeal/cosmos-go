package sorts

import (
	"github.com/NalbertLeal/cosmos-go/interfaces"
)

func median3[T interfaces.Order[T]](arr []T, p1 int, p2 int, p3 int) int {
	a := arr[p1]
	b := arr[p2]
	c := arr[p3]

	if a.CompareTo(b) < 0 && b.CompareTo(c) < 0 {
		return p2
	} else if b.CompareTo(a) < 0 && a.CompareTo(c) < 0 {
		return p1
	} else {
		return p3
	}
}

func pivotMedian3[T interfaces.Order[T]](arr []T, start int, end int) int {
	return median3(arr, start, (end-start)/2, end)
}

func pivotTukey[T interfaces.Order[T]](arr []T, start int, end int) int {
	i := median3(arr, start, (end-start)/2, end)
	j := median3(arr, start+1, (end-start+1)/2, end-1)
	k := median3(arr, start+2, (end-start-1)/2, end-2)

	if arr[i].CompareTo(arr[j]) < 0 && arr[j].CompareTo(arr[k]) < 0 {
		return j
	} else if arr[j].CompareTo(arr[i]) < 0 && arr[i].CompareTo(arr[k]) < 0 {
		return i
	}
	return k
}

func choosePivot[T interfaces.Order[T]](arr []T, start int, end int) int {
	pivotIndex := (start + end) / 2
	if end-start < 50 && end-start >= 8 {
		pivotIndex = pivotMedian3(arr, start, end)
	}
	if end-start >= 50 {
		pivotIndex = pivotTukey(arr, start, end)
	}
	return pivotIndex
}

func hoaresPartition[T interfaces.Order[T]](arr []T, start int, end int) int {
	pivotIndex := choosePivot(arr, start, end)
	if pivotIndex != start {
		arr[start], arr[pivotIndex] = arr[pivotIndex], arr[start]
	}
	pivot := arr[start]

	left := start + 1
	right := end
	for left <= right {
		for left <= right && arr[left].CompareTo(pivot) <= 0 {
			left += 1
		}
		for left <= right && arr[right].CompareTo(pivot) >= 0 {
			right -= 1
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left += 1
			right -= 1
		}
	}
	arr[start], arr[right] = arr[right], arr[start]
	return right
}

func nalbert3WayPartition[T interfaces.Order[T]](arr []T, start int, end int) (int, int) {
	pivotIndex := choosePivot(arr, start, end)
	if pivotIndex != start {
		arr[pivotIndex], arr[start] = arr[start], arr[pivotIndex]
	}
	pivot := arr[start]

	left := start + 1
	right := end
	pivotEnd := start

	for left <= right {
		if arr[left].CompareTo(pivot) < 0 {
			left += 1
		} else if arr[left].CompareTo(pivot) > 0 {
			arr[left], arr[right] = arr[right], arr[left]
			right -= 1
		} else {
			pivotEnd += 1
			arr[left], arr[pivotEnd] = arr[pivotEnd], arr[left]
			left += 1
		}
	}

	for i := pivotEnd; i >= start; i-- {
		arr[right], arr[i] = arr[i], arr[right]
		right -= 1
	}

	return left, right
}

func partition2Pivots[T interfaces.Order[T]](arr []T, start int, end int) (int, int) {
	if arr[start].CompareTo(arr[end]) > 0 {
		arr[start], arr[end] = arr[end], arr[start]
	}

	lowPivot := arr[start]
	lowPivotFinalIndex := start + 1

	highPivot := arr[end]
	highPivotFinalIndex := end - 1

	index := start + 1
	for index <= highPivotFinalIndex {
		if arr[index].CompareTo(lowPivot) < 0 {
			arr[index], arr[lowPivotFinalIndex] = arr[lowPivotFinalIndex], arr[index]
			lowPivotFinalIndex++
			index++
		} else if arr[index].CompareTo(highPivot) > 0 {
			arr[index], arr[highPivotFinalIndex] = arr[highPivotFinalIndex], arr[index]
			highPivotFinalIndex--
		} else {
			index++
		}
	}
	lowPivotFinalIndex--
	highPivotFinalIndex++
	arr[start], arr[lowPivotFinalIndex] = arr[lowPivotFinalIndex], arr[start]
	arr[end], arr[highPivotFinalIndex] = arr[highPivotFinalIndex], arr[end]
	return lowPivotFinalIndex, highPivotFinalIndex
}
