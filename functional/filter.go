package functional

import "github.com/NalbertLeal/cosmos-go/interfaces"

func Filter[T interfaces.Order[T]](arr []T, f func(T) bool) []T {
	result := []T{}
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}