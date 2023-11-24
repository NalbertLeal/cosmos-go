package searchs

import "github.com/NalbertLeal/cosmos-go/interfaces"

func BinSearch[T interfaces.Order[T]](arr []T, value T) int {
	index := 0
	for curr := len(arr) / 2; curr >= 1; curr /= 2 {
		for curr+index < len(arr) && arr[curr+index].CompareTo(value) <= 0 {
			index += curr
		}
	}
	if arr[index].CompareTo(value) == 0 {
		return index
	}
	return len(arr)
}