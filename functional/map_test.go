package functional

import (
	"math/rand"
	"testing"
)

func TestMapFloatToInt(t *testing.T) {
	f := func(f float64) int {
		return int(f)
	}

	arr := make([]float64, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Float64() * 100
	}

	result := Map(arr, f)

	if len(result) != len(arr) {
		t.Error("input length is different of the output length")
	}
	for i := 0; i < len(arr); i++ {
		if int(arr[i]) != result[i] {
			t.Error("The converted value of result in position ", i, " is ", result[i], " but should be ", int(arr[i]))
		}
	}
}
