package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func CycleSort[T interfaces.Order[T]](arr []T) {
	cycleSort(arr, 0, len(arr)-1)
}

func cycleSort[T interfaces.Order[T]](arr []T, start int, end int) {
	for cycleStart := start; cycleStart < end; cycleStart += 1 {
		item := arr[cycleStart]

		pos := cycleStart
		for i := cycleStart + 1; i <= end; i += 1 {
			if arr[i].CompareTo(item) < 0 {
				pos += 1
			}
		}

		if pos == cycleStart {
			continue
		}

		for item.CompareTo(arr[pos]) == 0 {
			pos += 1
		}

		if pos != cycleStart {
			item, arr[pos] = arr[pos], item
		}

		for pos != cycleStart {
			pos = cycleStart

			for i := cycleStart + 1; i <= end; i += 1 {
				if arr[i].CompareTo(item) < 0 {
					pos += 1
				}
			}
			for item.CompareTo(arr[pos]) == 0 {
				pos += 1
			}
			if item.CompareTo(arr[pos]) != 0 {
				item, arr[pos] = arr[pos], item
			}
		}
	}
}
