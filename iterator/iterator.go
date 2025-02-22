package iterator

import (
	"context"

	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/iter"
)

// An Iterator provides you a standard API for iterating over collections of
// data.
type Iterator[T any] struct {
	inner iter.Interface[T]
}

// NewIterator returns a new Iterator instance which wraps the provided
// iter.Interface.
func NewIterator[T any](iter iter.Interface[T]) *Iterator[T] {
	return &Iterator[T]{inner: iter}
}

// FromSlice returns a new Iterator instance which wraps the provided slice.
func FromSlice[T any](slice []T) *Iterator[T] {
	return NewIterator(iter.FromSlice(slice))
}

// FromMap returns a new Iterator instance which wraps the provided maps and
// yields instances of the iter.MapEntry type.
func FromMap[K comparable, V any](data map[K]V) *Iterator[iter.MapEntry[K, V]] {
	return NewIterator(iter.FromMap(data))
}

// FromChan returns a new Iterator instance which wraps the provided channel.
func FromChan[T any](data chan T) *Iterator[T] {
	return NewIterator[T](iter.FromChannel(data))
}

// Next implements iter.Interface and allows Iterator[T] to be used as a bare
// iterator.
func (i *Iterator[T]) Next() (T, bool) {
	return i.inner.Next()
}

// AdvanceBy advances the iterator by n values.
func (i *Iterator[T]) AdvanceBy(n int) *Iterator[T] {
	return NewIterator(iter.AdvanceBy(i.inner, n))
}

// StepBy returns a new Iterator which starts at this iterators next value, but
// which steps by the specified number of items on each subsequent iteration.
func (i *Iterator[T]) StepBy(step int) *Iterator[T] {
	return NewIterator(iter.StepBy(i.inner, step))
}

// Filter returns an Iterator which will only yield elements for which
// satisfies the provided Predicate.
func (i *Iterator[T]) Filter(fn iter.Predicate[T]) *Iterator[T] {
	return NewIterator(iter.Filter(i.inner, fn))
}

// SkipWhile returns an Iterator which skips elements for as long as the
// provided predicate is satisfied. Once a false value is returned by the
// predicate all values will be yielded by the iterator as normal.
func (i *Iterator[T]) SkipWhile(fn iter.Predicate[T]) *Iterator[T] {
	return NewIterator(iter.SkipWhile(i.inner, fn))
}

// TakeWhile returns an Iterator which yields elements for as long as the
// provided predicate is satisfied. Once a false value is returned by the
// predicate the iterator will cease to yield further values.
func (i *Iterator[T]) TakeWhile(fn iter.Predicate[T]) *Iterator[T] {
	return NewIterator(iter.TakeWhile(i.inner, fn))
}

// Skip returns an Iterator which skips over the first n elements. The
// remaining elements are all yielded as normal.
func (i *Iterator[T]) Skip(n int) *Iterator[T] {
	return NewIterator(iter.Skip(i.inner, n))
}

// Take returns an Iterator which yields the first n elements, or all elements
// if the iterator contains fewer than n elements, and then ceases to yield
// values.
func (i *Iterator[T]) Take(n int) *Iterator[T] {
	return NewIterator(iter.Take(i.inner, n))
}

// Inspect returns an Iterator which calls the specified closure on each
// element yielded by the iterator until the iterator is exhausted.
func (i *Iterator[T]) Inspect(fn iter.InspectFunc[T]) *Iterator[T] {
	return NewIterator(iter.Inspect(i.inner, fn))
}

// Chain returns an Iterator which iterates over both of the provided iterators
// in the order in which they are provided.
func (i *Iterator[T]) Chain(other iter.Interface[T]) *Iterator[T] {
	return NewIterator(iter.Chain(i.inner, other))
}

// Intersperse returns a new Iterator which injects a copy of the provided
// separator between items yielded by the Iterator.
func (i *Iterator[T]) Intersperse(sep T) *Iterator[T] {
	return NewIterator(iter.Intersperse(i.inner, sep))
}

// Interleave returns a new Iterator which alternates elements from two
// iterators until both Iterators are fully consumed.
func (i *Iterator[T]) Interleave(other iter.Interface[T]) *Iterator[T] {
	return NewIterator(iter.Interleave(i.inner, other))
}

// Count consumes the Iterator and returns the count of all items in the
// iterator.
func (i *Iterator[T]) Count() int {
	return iter.Count(i.inner)
}

// Last consumes the Iterator and returns the last element.
func (i *Iterator[T]) Last() (last T) {
	return iter.Last(i.inner)
}

// Nth returns the nth item in the Iterator.
func (i *Iterator[T]) Nth(n int) (value T, ok bool) {
	return iter.Nth(i.inner, n)
}

// ForEach consumes the Iterator and calls the provided closure on each element.
func (i *Iterator[T]) ForEach(fn func(*T)) {
	iter.ForEach(i.inner, fn)
}

// Find searches the Iterator for an element that satisfies the provided
// predicate.
func (i *Iterator[T]) Find(fn iter.Predicate[T]) (T, bool) {
	return iter.Find(i.inner, fn)
}

// SizeHint implements the iter.SizeHinter interface and attempts to provide
// size hint information about the Iterator.
//
// Note: not all Iterator implements or can implement SizeHint, so ensure that
// any logic which uses these values handles all possible variations of the
// returned lower and upper bounds.
func (i *Iterator[T]) SizeHint() (int64, oxide.Option[int64]) {
	return iter.SizeHint(i.inner)
}

// Partition consumes the Iterator and produces two collections:
// 1: One collection contains all values yielded by the iterator for which the
// provided Predicate is satisfied (returned true)
// 2: The other collection contains all values yielded by the iterator for
// which the Predicate was not satisfied (returned false).
func (i *Iterator[T]) Partition(fn iter.Predicate[T]) ([]T, []T) {
	return iter.Partition(i.inner, fn)
}

// All consumes the Iterator, returning a boolean value which indicates whether
// all the values yielded by the iterator satisfied the provided predicate.
func (i *Iterator[T]) All(fn iter.Predicate[T]) bool {
	return iter.All(i.inner, fn)
}

// Any consumes the Iterator, returning a boolean value which indicates whether
// any of the values yielded by the iterator satisfied the provided predicate.
func (i *Iterator[T]) Any(fn iter.Predicate[T]) bool {
	return iter.Any(i.inner, fn)
}

// Peekable returns a new iter.Peekable iterator which, in addition to the
// standard Next() method, also implements Peek() which allows callers to view
// the next value that an iterator would yield, without consuming it.
func (i *Iterator[T]) Peekable() iter.Peekable[T] {
	return iter.IntoPeekable(i.inner)
}

// CollectSlice collects the Iterator into a slice of type T.
func (i *Iterator[T]) CollectSlice() []T {
	return iter.CollectSlice(i.inner)
}

// CollectChan collects the Iterator into a channel of type T.
func (i *Iterator[T]) CollectChan(ctx context.Context, buffer int) <-chan T {
	return iter.CollectChan[T](ctx, i.inner, buffer)
}
