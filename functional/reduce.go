package functional

func Reduce[T any, O any](arr []T, init O, f func(O, T) O) O {
	result := init
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}