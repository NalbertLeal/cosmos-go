package sorts

type myInt int

func (m myInt) CompareTo(o myInt) int {
	if m < o {
		return -1
	} else if m == o {
		return 0
	}
	return 1
}
