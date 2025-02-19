package iter

import (
	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/constraints"
)

// FromSlice returns a new iterator that can be used to iterate over the
// provided slice.
func FromSlice[T any](slice []T) Interface[T] {
	return &sliceIterator[T]{slice: slice}
}

// FromMap returns a new iterator which can be used to traverse through all the
// MapEntry pairs in the provided map.
//
// Note: Just like when iterating through a map's data using a "for / range",
// the order in which the key-value pairs are yielded by the iterator is
// non-deterministic.
func FromMap[K comparable, V any](data map[K]V) Interface[MapEntry[K, V]] {
	return &mapIterator[K, V]{data: data}
}

// FromChannel returns a new iterator which can be used to traverse through all
// the values yielded by the provided channel.
func FromChannel[T any](data <-chan T) Interface[T] {
	return &chanIterator[T]{data: data}
}

// Range returns a new iterator which can be used to iterate over all values in
// a range of numbers.
func Range[T constraints.Integer](from, to T) Interface[T] {
	return &rangeIterator[T]{current: from, max: to}
}

type sliceIterator[T any] struct {
	slice []T
}

func (i *sliceIterator[T]) Next() (value T, ok bool) {
	if len(i.slice) == 0 {
		return
	}

	value = i.slice[0]
	i.slice = i.slice[1:]
	return value, true
}

func (i *sliceIterator[T]) SizeHint() (int64, oxide.Option[int64]) {
	max := cap(i.slice)
	if max == 0 {
		return 0, oxide.None[int64]()
	}

	return 0, oxide.Some(int64(max))
}

type mapIterator[K comparable, V any] struct {
	data map[K]V
}

func (i *mapIterator[K, V]) Next() (MapEntry[K, V], bool) {
	for key, value := range i.data {
		delete(i.data, key)
		return MapEntry[K, V]{Key: key, Val: value}, true
	}

	var zero MapEntry[K, V]
	return zero, false
}

type chanIterator[T any] struct {
	data <-chan T
}

func (i *chanIterator[T]) Next() (T, bool) {
	select {
	case value, ok := <-i.data:
		return value, ok
	default:
		var zero T
		return zero, false
	}
}

func (i *chanIterator[T]) SizeHint() (int64, oxide.Option[int64]) {
	max := cap(i.data)
	if max == 0 {
		return 0, oxide.None[int64]()
	}

	return 0, oxide.Some(int64(max))
}

type rangeIterator[T constraints.Integer] struct {
	current, max T
}

func (r *rangeIterator[T]) Next() (T, bool) {
	if r.current >= r.max {
		return 0, false
	}

	r.current++
	return r.current, true
}
