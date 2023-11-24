package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func merge[T interfaces.Order[T]](arr []T, start int, middle int, end int) {
	tempArr := make([]T, end-start+1)

	leftI := start
	rightI := middle + 1
	tempArrI := 0
	for leftI < middle+1 && rightI <= end {
		if arr[leftI].CompareTo(arr[rightI]) <= 0 {
			tempArr[tempArrI] = arr[leftI]
			leftI++
		} else {
			tempArr[tempArrI] = arr[rightI]
			rightI++
		}
		tempArrI++
	}
	for leftI < middle+1 {
		tempArr[tempArrI] = arr[leftI]
		leftI++
		tempArrI++
	}
	for rightI <= end {
		tempArr[tempArrI] = arr[rightI]
		rightI++
		tempArrI++
	}

	copy(arr[start:end+1], tempArr)
}
