package sorts

import "github.com/NalbertLeal/cosmos-go/interfaces"

func TimSort[T interfaces.Order[T]](arr []T) {
	timSort(arr, 0, len(arr)-1)
}

func timSort[T interfaces.Order[T]](arr []T, start int, end int) {
	// calculate runs length
	runLength := calculateRunLengthBits(arr)
	// insertion sort on each run
	for insertionStart := 0; insertionStart < len(arr); insertionStart += runLength {
		insertionEnd := len(arr) - 1
		if insertionEnd > insertionStart+runLength-1 {
			insertionEnd = insertionStart + runLength - 1
		}
		insertionSort(arr, insertionStart, insertionEnd)
	}
	// merge the runs
	mergeSize := runLength
	for mergeSize < len(arr) {
		for left := 0; left < len(arr); left += mergeSize * 2 {
			middle := left + mergeSize - 1
			if middle > len(arr)-1 {
				middle = len(arr) - 1
			}
			right := left + (mergeSize * 2) - 1
			if right > len(arr)-1 {
				right = len(arr) - 1
			}

			if middle < right {
				merge(arr, left, middle, right)
			}
		}
		mergeSize = mergeSize * 2
	}
}

func calculateRunLengthDiv2[T interfaces.Order[T]](arr []T) int {
	const THRESHOLD int = 64
	runLength := len(arr)
	remainder := 0
	for runLength > THRESHOLD {
		if runLength%2 == 1 {
			remainder = 1
		}
		runLength = runLength / 2
	}
	return runLength + remainder
}

/// O valor de Minrun é determinado com base no valor de N, seguindo os seguintes princípios:
/// 1 -Não deve ser muito longo, pois o Minrun será posteriormente submetido ao InsertionSort, que é efetivo para tamanhos menores de vetores (Runs).
/// 2 - Não deve ser muito curto, pois quanto mais curto for o Run, mais Runs terão que ser mescladas no próximo passo.
/// 3 - Seria bom se N / Minrun fosse uma potência de 2 (ou próximo disso). Esse requisito resulta do fato de que o MergeSort funciona melhor nos Runs que têm aproximadamente o mesmo comprimento.
/// No artigo da proposta do timsort o autor se refere a seus próprios experimentos, mostrando que com Minrun > 256, o princípio 1 não está satisfeito, e com Minrun <8, o princípio 2 não está satisfeito e os comprimentos que atingem melhores performances variam de 32 a 65.
/// Exceção: se N < 64, então Minrun = N
func calculateRunLengthBits[T interfaces.Order[T]](arr []T) int {
	const THRESHOLD int = 64
	r := 0
	n := len(arr)
	for n >= THRESHOLD {
		r |= n & 1
		n >>= 1
	}
	return n + r
}
