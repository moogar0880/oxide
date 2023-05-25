package math

import "github.com/moogar0880/oxide/iter"

// A Number is a generic type which accounts for all of the builtin numerical
// types.
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Sum returns the sum of all values yielded by the provided iterator.
func Sum[T Number](iterator iter.Interface[T]) (value T) {
	return iter.Fold(iterator, value, func(accum T, value *T) T {
		accum += *value

		return accum
	})
}

// Product returns the product of all values yielded by the provided iterator.
func Product[T Number](iterator iter.Interface[T]) (value T) {
	return iter.Fold(iterator, value, func(accum T, value *T) T {
		if accum == 0 {
			return *value
		}

		return accum * *value
	})
}

// Min returns the minimum of all values yielded by the provided iterator.
func Min[T Number](iterator iter.Interface[T]) T {
	min, _ := iterator.Next()

	for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
		if item < min {
			min = item
		}
	}

	return min
}

// Max returns the maximum of all values yielded by the provided iterator.
func Max[T Number](iterator iter.Interface[T]) (value T) {
	max, _ := iterator.Next()

	for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
		if item > max {
			max = item
		}
	}

	return max
}
