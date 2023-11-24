package interfaces

const (
	LESS = -1
	EQUAL = 0
	MORE = 1
)

type Order[T any] interface {
	CompareTo(T) int
}