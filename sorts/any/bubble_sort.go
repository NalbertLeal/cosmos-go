package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func BubleSort[T interfaces.Order[T]](arr []T) {
	bubblesort2(arr, 0, len(arr)-1)
}

func bubblesort[T interfaces.Order[T]](arr []T, start int, end int) {
	for i := start; i <= end; i++ {
		for j := start + 1; j <= end; j++ {
			if arr[j-1].CompareTo(arr[j]) > 0 {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func bubblesort2[T interfaces.Order[T]](arr []T, start int, end int) {
	for i := end + 1; i > start+1; i -= 2 {
		x := arr[start]
		y := arr[start+1]
		if y.CompareTo(x) < 0 {
			x, y = y, x
		}
		for j := start + 2; j < i; j++ {
			z := arr[j]

			w := y
			if y.CompareTo(z) <= 0 {
				y = z
			} else {
				w = z
			}

			if x.CompareTo(z) <= 0 {
				arr[j-2] = x
				x = w
			} else {
				arr[j-2] = z
			}
		}
		arr[i-2] = x
		arr[i-1] = y
	}
}
