package functional

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"
)

func generateRandomIntSlice(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(1000) + 1
	}
	return arr
}

func isZippedValuesFromslice(zipped [][]int, arr []int, valuePosition int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != zipped[i][valuePosition] {
			return false
		}
	}
	return true
}

func isZippedValuesZero(zipped [][]int, sliceSize int, valuePosition int) (bool, error) {
	var zeroValue int
	for i := 0; i < sliceSize; i++ {
		if zipped[i][valuePosition] == zeroValue {
			return false, errors.New("at zipped[" + strconv.Itoa(i) + "][" + strconv.Itoa(valuePosition) + "] the value is zero, but should not be")
		}
	}
	for i := sliceSize; i < len(zipped); i++ {
		if zipped[i][valuePosition] != zeroValue {
			return false, errors.New("at zipped[" + strconv.Itoa(i) + "][" + strconv.Itoa(valuePosition) + "] the value is not zero, but should be")
		}
	}
	return true, nil
}

func TestZip3Slices1000Len(t *testing.T) {
	arr1 := generateRandomIntSlice(1000)
	arr2 := generateRandomIntSlice(1000)
	arr3 := generateRandomIntSlice(1000)

	zipped := Zip(arr1, arr2, arr3)
	if len(zipped) != 1000 {
		t.Error("The zipped slice should have length of 1000, instead length is ", len(zipped))
	}
	for i := 0; i < len(zipped); i++ {
		if len(zipped[0]) != 3 {
			t.Error("The slice of values inside the zipped slice should have length 3, instead at possition ", i, " the langth is ", len(zipped[i]))
		}
	}

	if !isZippedValuesFromslice(zipped, arr1, 0) {
		t.Error("Values of arr1 is not in the zipped")
	}
	if !isZippedValuesFromslice(zipped, arr2, 1) {
		t.Error("Values of arr2 is not in the zipped")
	}
	if !isZippedValuesFromslice(zipped, arr3, 2) {
		t.Error("Values of arr3 is not in the zipped")
	}
}

func TestZip4SlicesDifferentLen(t *testing.T) {
	arr1 := generateRandomIntSlice(10000)
	arr2 := generateRandomIntSlice(5000)
	arr3 := generateRandomIntSlice(400)
	arr4 := generateRandomIntSlice(1500)

	zipped := Zip(arr1, arr2, arr3, arr4)
	if len(zipped) != 10000 {
		t.Error("The zipped slice should have length of 1000, instead length is ", len(zipped))
	}
	for i := 0; i < len(zipped); i++ {
		if len(zipped[0]) != 4 {
			t.Error("The slice of values inside the zipped slice should have length 3, instead at possition ", i, " the langth is ", len(zipped[i]))
		}
	}

	if !isZippedValuesFromslice(zipped, arr1, 0) {
		t.Error("Values at possition 0 of arr1 is not in the zipped")
	}
	if ok, err := isZippedValuesZero(zipped, len(arr1), 0); !ok {
		t.Error(err)
	}

	if !isZippedValuesFromslice(zipped, arr2, 1) {
		t.Error("Values at possition 1 of arr2 is not in the zipped")
	}
	if ok, err := isZippedValuesZero(zipped, len(arr2), 1); !ok {
		t.Error(err)
	}

	if !isZippedValuesFromslice(zipped, arr3, 2) {
		t.Error("Values at possition 2 of arr3 is not in the zipped")
	}
	if ok, err := isZippedValuesZero(zipped, len(arr3), 2); !ok {
		t.Error(err)
	}

	if !isZippedValuesFromslice(zipped, arr4, 3) {
		t.Error("Values at possition 3 of arr4 is not in the zipped")
	}
	if ok, err := isZippedValuesZero(zipped, len(arr4), 3); !ok {
		t.Error(err)
	}
}
