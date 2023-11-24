package branchless

// Based on package branchless from https://github.com/alasdairforsythe/branchless

func MinInt64(a int64, b int64) int64 {
	return b ^ ((a ^ b) & ((a - b) >> 63))
}

func MaxInt64(a int64, b int64) int64 {
	return a ^ ((a ^ b) & ((a - b) >> 63))
}

func EvenInt64(a int64) int64 {
	return (a & 1) ^ 1
}

func OddInt64(a int64) int64 {
	return a & 1
}

func LessThanInt64(a int64, b int64) int64 {
	return ((a - b) >> 63) & 1
}

func GreaterThanInt64(a int64, b int64) int64 {
	return ((b - a) >> 63) & 1
}

func EqualToInt64(a int64, b int64) int64 {
	return (((a ^ b) | -(a ^ b)) >> 63 & 1) ^ 1
}

func NotEqualInt64(a int64, b int64) int64 {
	return (((a ^ b) | -(a ^ b)) >> 63 & 1)
}

func IsZeroInt64(a int64) int64 {
	return (((a | -a) >> 63) & 1) ^ 1
}

func IsDivisibleByInt64(a int64, b int64) int64 {
	return (a % b & 1) ^ 1
}

func AbsInt64(a int64) int64 {
	return (a + (a >> 63)) ^ (a >> 63)
}

func SignInt64(a int64) int64 {
	return (a >> 63) | (-(a >> 63))
}
