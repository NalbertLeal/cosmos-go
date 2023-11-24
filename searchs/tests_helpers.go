package searchs

type myInt int

func (m myInt) CompareTo(o myInt) int {
	if m < o {
		return -1
	} else if m == o {
		return 0
	}
	return 1
}

func isSorted(arr []myInt) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}