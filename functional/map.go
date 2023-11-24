package functional

func Map[T any, O any](arr []T, mapper func(T) O) []O {
	result := make([]O, len(arr))
	for i, v := range arr {
		result[i] = mapper(v)
	}
	return result
}