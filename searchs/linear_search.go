package searchs

import "github.com/NalbertLeal/cosmos-go/interfaces"

func LinearSearch[T interfaces.Order[T]](arr []T, value T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].CompareTo(value) == 0 {
			return i
		}
	}
	return len(arr)
}