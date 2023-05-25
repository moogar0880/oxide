package iter

import "github.com/moogar0880/oxide"

// The functions and types in this file are denoted as "extensions" of the core
// iterator API because they can not currently be properly expressed in the
// higher-level Iterator API due to current compiler limitations regarding
// adding generic trait bounds on methods, namely the fact that you are
// expressly forbidden from doing so.

// A FoldFunc defines a function which "folds" every element into an
// accumulator (A), returning the new state of the accumulator on each
// invocation.
type FoldFunc[T, A any] func(A, *T) A

// A MapFunc defines a function which maps from one type (F) to another (T).
type MapFunc[F, T any] func(F) T

type FilterMapFunc[F, T any] func(F) oxide.Option[T]
type FindMapFunc[F, T any] func(F) oxide.Option[T]

// Fold returns the final value of the accumulator (A) after consuming the
// provided Interface.
func Fold[T, A any](iter Interface[T], init A, fn FoldFunc[T, A]) A {
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		init = fn(init, &item)
	}

	return init
}

// Map returns an Interface which will call the provided MapFunc as the iterator
// is consumed.
func Map[F, T any](iter Interface[F], fn MapFunc[F, T]) Interface[T] {
	return &mappingIterator[F, T]{
		inner: iter,
		fn:    fn,
	}
}

type mappingIterator[F, T any] struct {
	inner Interface[F]
	fn    MapFunc[F, T]
}

func (i *mappingIterator[F, T]) Next() (T, bool) {
	val, ok := i.inner.Next()
	if !ok {
		var zero T
		return zero, false
	}

	return i.fn(val), true
}

func FilterMap[F, T any](iter Interface[F], fn FilterMapFunc[F, T]) Interface[T] {
	return &filterMapIterator[F, T]{
		inner: iter,
		fn:    fn,
	}
}

type filterMapIterator[F, T any] struct {
	inner Interface[F]
	fn    FilterMapFunc[F, T]
}

func (i *filterMapIterator[F, T]) Next() (T, bool) {
	var result oxide.Option[T]

	for item, ok := i.inner.Next(); ok; item, ok = i.inner.Next() {
		if result = i.fn(item); result.IsSome() {
			return result.Value(), true
		}
	}

	return result.Value(), false
}

type Enumerated[T any] struct {
	Index int
	Value T
}

func Enumerate[T any](iter Interface[T]) Interface[Enumerated[T]] {
	return &enumeratedIterator[T]{inner: iter}
}

type enumeratedIterator[T any] struct {
	inner   Interface[T]
	current int
}

func (i *enumeratedIterator[T]) Next() (Enumerated[T], bool) {
	value, ok := i.inner.Next()
	if !ok {
		var zero Enumerated[T]
		return zero, false
	}

	enum := Enumerated[T]{Index: i.current, Value: value}
	i.current += 1

	return enum, true
}

func MapWhile[F, T any](iter Interface[F], fn FilterMapFunc[F, T]) Interface[T] {
	return &mapWhileIterator[F, T]{inner: iter, fn: fn}
}

type mapWhileIterator[F, T any] struct {
	inner Interface[F]
	fn    FilterMapFunc[F, T]
}

func (i *mapWhileIterator[F, T]) Next() (T, bool) {
	value, ok := i.inner.Next()
	if !ok {
		var zero T
		return zero, ok
	}

	result := i.fn(value)
	return result.Value(), result.IsSome()
}

func Fuse[T any](iter Interface[T]) Interface[T] {
	return &fuseIter[T]{inner: iter}
}

type fuseIter[T any] struct {
	inner Interface[T]
	done  bool
}

func (i *fuseIter[T]) Next() (zero T, ok bool) {
	if i.done {
		return
	}

	value, ok := i.inner.Next()
	if !ok {
		i.done = true
	}

	return value, ok
}

func FindMap[F, T any](iter Interface[F], fn FindMapFunc[F, T]) oxide.Option[T] {
	var result oxide.Option[T]
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if result = fn(item); result.IsSome() {
			return result
		}
	}

	return oxide.None[T]()
}
