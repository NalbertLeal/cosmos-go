package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func CountSort[T interfaces.Integers](arr []T, max int, min int) {
	countSlice := make([]int, max-min+1)

	for _, v := range arr {
		countSlice[v-min] += 1
	}

	j := 0
	for i := 0; i < len(countSlice); i += 1 {
		for e := 0; e < countSlice[i]; e += 1 {
			arr[j] = i + min
			j += 1
		}
	}
}
