package functional

func Zip[T any](arr ...[]T) [][]T {
	biggestLen := 0
	for i := 0; i < len(arr); i++ {
		if biggestLen < len(arr[i]) {
			biggestLen = len(arr[i])
		}
	}

	var zeroValueT T

	result := make([][]T, biggestLen)
	for i := 0; i < biggestLen; i++ {
		values := make([]T, len(arr))
		for j := 0; j < len(arr); j++ {
			if len(arr[j]) > i {
				values[j] = arr[j][i]
			} else {
				values[j] = zeroValueT
			}
		}
		result[i] = values
	}

	return result
}
