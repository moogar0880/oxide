package math

import (
	"github.com/moogar0880/oxide/constraints"
	"github.com/moogar0880/oxide/iter"
)

// A Number is a generic type which accounts for all the builtin
// numerical types.
//
// Deprecated. Use the interface definition from `constraints` directly.
type Number = constraints.Number

// Sum returns the sum of all values yielded by the provided iterator.
func Sum[T constraints.Number](iterator iter.Interface[T]) (value T) {
	return iter.Fold(iterator, value, func(accum T, value *T) T {
		accum += *value

		return accum
	})
}

// Product returns the product of all values yielded by the provided iterator.
func Product[T constraints.Number](iterator iter.Interface[T]) (value T) {
	return iter.Fold(iterator, value, func(accum T, value *T) T {
		if accum == 0 {
			return *value
		}

		return accum * *value
	})
}

// Min returns the minimum of all values yielded by the provided iterator.
func Min[T constraints.Number](iterator iter.Interface[T]) T {
	min, _ := iterator.Next()

	for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
		if item < min {
			min = item
		}
	}

	return min
}

// Max returns the maximum of all values yielded by the provided iterator.
func Max[T constraints.Number](iterator iter.Interface[T]) (value T) {
	max, _ := iterator.Next()

	for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
		if item > max {
			max = item
		}
	}

	return max
}
