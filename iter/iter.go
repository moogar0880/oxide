package iter

import "github.com/moogar0880/oxide"

// Count consumes the provided iterator, returning the number of iterations
// that were made before the first None value is encountered.
func Count[T any](iter Interface[T]) int {
	return Fold(iter, 0, func(count int, _ *T) int {
		return count + 1
	})
}

// Last consumes the provided iterator, returning the final element.
func Last[T any](iter Interface[T]) (last T) {
	return Fold(iter, last, func(accum T, value *T) T {
		accum = *value

		return accum
	})
}

// AdvanceBy eagerly consumes n elements from the provided iterator by calling
// Next either n times or until a None value is encountered.
func AdvanceBy[T any](iter Interface[T], n int) Interface[T] {
	for index := 0; index < n; index++ {
		iter.Next()
	}

	return iter
}

// Nth returns the nth element from the provided iterator.
//
// Note: the provided iterator is considered to be zero indexed, meaning n=0
// will return the first element of the iterator.
func Nth[T any](iter Interface[T], n int) (T, bool) {
	AdvanceBy(iter, n)

	return iter.Next()
}

type stepByIterator[T any] struct {
	inner      Interface[T]
	stepBy     int
	firstTaken bool
}

// StepBy returns a new iterator which starts at the same value as the provided
// iterator, but which steps by the specified number of items on each
// iteration.
func StepBy[T any](iter Interface[T], step int) Interface[T] {
	return &stepByIterator[T]{inner: iter, stepBy: step - 1}
}

func (i *stepByIterator[T]) Next() (T, bool) {
	if !i.firstTaken || i.stepBy == 0 {
		i.firstTaken = true
		return i.inner.Next()
	}

	return Nth(i.inner, i.stepBy)
}

// ForEach consumes the provided iterator and calls the provided closure on
// each element.
func ForEach[T any](iter Interface[T], fn func(*T)) {
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		fn(&item)
	}
}

type Predicate[T any] func(*T) bool

// Find searches the provided iterator for an element that satisfies the
// provided predicate.
func Find[T any](iter Interface[T], fn Predicate[T]) (T, bool) {
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if fn(&item) {
			return item, ok
		}
	}

	var zero T
	return zero, false
}

type filterIterator[T any] struct {
	inner Interface[T]
	fn    Predicate[T]
}

func (i *filterIterator[T]) Next() (T, bool) {
	return Find(i.inner, i.fn)
}

// Filter returns an iterator which will only yield elements for which
// satisfies the provided Predicate.
func Filter[T any](iter Interface[T], fn Predicate[T]) Interface[T] {
	return &filterIterator[T]{inner: iter, fn: fn}
}

type skipWhileIterator[T any] struct {
	inner Interface[T]
	fn    Predicate[T]
	done  bool
}

func (i *skipWhileIterator[T]) Next() (T, bool) {
	if i.done {
		return i.inner.Next()
	}

	for item, ok := i.inner.Next(); ok; item, ok = i.inner.Next() {
		if !i.fn(&item) {
			i.done = true
			return item, ok
		}
	}

	var zero T
	return zero, false
}

// SkipWhile returns an iterator which skips elements for as long as the
// provided predicate is satisfied. Once a false value is returned by the
// predicate all values will be yielded by the iterator as normal.
func SkipWhile[T any](iter Interface[T], fn Predicate[T]) Interface[T] {
	return &skipWhileIterator[T]{inner: iter, fn: fn}
}

type takeWhileIterator[T any] struct {
	inner Interface[T]
	fn    Predicate[T]
}

func (i *takeWhileIterator[T]) Next() (T, bool) {
	value, ok := i.inner.Next()
	if !ok {
		return value, ok
	}

	if i.fn(&value) {
		return value, ok
	}

	var zero T
	return zero, false
}

// TakeWhile returns an iterator which yields elements for as long as the
// provided predicate is satisfied. Once a false value is returned by the
// predicate the iterator will cease to yield further values.
func TakeWhile[T any](iter Interface[T], fn Predicate[T]) Interface[T] {
	return &takeWhileIterator[T]{inner: iter, fn: fn}
}

// TODO: MapWhile

type skipIterator[T any] struct {
	inner   Interface[T]
	n       int
	skipped bool
}

func (i *skipIterator[T]) Next() (T, bool) {
	if i.skipped {
		return i.inner.Next()
	}

	i.skipped = true
	return Nth(i.inner, i.n)
}

// Skip returns an Iterator which skips over the first n elements. The
// remaining elements are all yielded as normal.
func Skip[T any](iter Interface[T], n int) Interface[T] {
	return &skipIterator[T]{inner: iter, n: n}
}

type takeIterator[T any] struct {
	inner Interface[T]
	n     int
	taken int
}

func (i *takeIterator[T]) Next() (T, bool) {
	if i.taken == i.n {
		var zero T
		return zero, false
	}

	i.taken = i.taken + 1
	return i.inner.Next()
}

// Take returns an iterator which yields the first n elements, or all elements
// if the iterator contains fewer than n elements, and then ceases to yield
// values.
func Take[T any](iter Interface[T], n int) Interface[T] {
	return &takeIterator[T]{inner: iter, n: n}
}

type InspectFunc[T any] func(*T)

type inspectIterator[T any] struct {
	inner Interface[T]
	fn    InspectFunc[T]
}

func (i *inspectIterator[T]) Next() (T, bool) {
	value, ok := i.inner.Next()
	if !ok {
		return value, ok
	}

	i.fn(&value)

	return value, ok
}

// Inspect returns an iterator which calls the specified closure on each
// element yielded by the iterator until the iterator is exhausted.
func Inspect[T any](iter Interface[T], fn InspectFunc[T]) Interface[T] {
	return &inspectIterator[T]{inner: iter, fn: fn}
}

// Partition consumes the provided iterator and produces two collections - One
// collection contains all values yielded by the iterator for which the
// provided Predicate is satisfied (returned true), and the other collection
// contains all values yielded by the iterator for which the Predicate was not
// satisfied (returned false).
func Partition[T any](iter Interface[T], fn Predicate[T]) ([]T, []T) {
	left := make([]T, 0)
	right := make([]T, 0)

	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if fn(&item) {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	return left, right
}

// Chain returns an iterator which iterates over both of the provided iterators
// in the order in which they are provided.
func Chain[T any](first Interface[T], second Interface[T]) Interface[T] {
	return &chainIterator[T]{first: first, second: second}
}

type chainIterator[T any] struct {
	first  Interface[T]
	second Interface[T]
}

func (i *chainIterator[T]) Next() (T, bool) {
	value, ok := i.first.Next()
	if !ok {
		return i.second.Next()
	}

	return value, ok
}

// Zip returns an iterator which "zips" up the two provided iterators. This
// iterator will return an array of size 2 which contains the next items yielded
// from both iterators.
func Zip[T any](left Interface[T], right Interface[T]) Interface[[2]T] {
	return &zipIterator[T]{left: left, right: right}
}

type zipIterator[T any] struct {
	left  Interface[T]
	right Interface[T]
}

func (i *zipIterator[T]) Next() ([2]T, bool) {
	v1, ok1 := i.left.Next()
	v2, ok2 := i.right.Next()

	if ok1 && ok2 {
		return [2]T{v1, v2}, true
	} else if ok1 {
		var zero T
		return [2]T{v1, zero}, true
	} else if ok2 {
		var zero T
		return [2]T{zero, v2}, true
	}

	var zero T
	return [2]T{zero, zero}, false
}

// All consumes the provided iterator, returning a boolean value which
// indicates whether all the values yielded by the iterator satisfied the
// provided predicate.
func All[T any](iter Interface[T], fn Predicate[T]) bool {
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if !fn(&item) {
			return false
		}
	}

	return true
}

// Any consumes the provided iterator, returning a boolean value which
// indicates whether any of the values yielded by the iterator satisfied the
// provided predicate.
func Any[T any](iter Interface[T], fn Predicate[T]) bool {
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if fn(&item) {
			return true
		}
	}

	return false
}

// SizeHint returns a tuple which represents the lower bound and the upper
// bound of the provided iterator.
//
// An oxide.None value for the upper bound indicates that there is no known
// upper bound on the underlying collection.
func SizeHint[T any](iter Interface[T]) (int64, oxide.Option[int64]) {
	if hint, ok := iter.(SizeHinter); ok {
		return hint.SizeHint()
	}

	return 0, oxide.None[int64]()
}

type peekableIterator[T any] struct {
	inner  Interface[T]
	peeked oxide.Option[oxide.Option[T]]
}

func (i *peekableIterator[T]) Next() (T, bool) {
	if i.peeked.IsSome() {
		value := i.peeked.Value()

		// Clear the peeked value that we're about to return.
		i.peeked = oxide.None[oxide.Option[T]]()

		return value.Unpack()
	}

	return i.inner.Next()
}

func (i *peekableIterator[T]) Peek() oxide.Option[T] {
	value, ok := i.inner.Next()
	if !ok {
		i.peeked = oxide.None[oxide.Option[T]]()

		return i.peeked.Value()
	}

	i.peeked = oxide.Some[oxide.Option[T]](oxide.Some(value))
	return i.peeked.Value()
}

// IntoPeekable returns a new Peekable iterator which, in addition to the
// standard Next() method, also implements Peek() which allows callers to view
// the next value that an iterator would yield, without consuming it.
func IntoPeekable[T any](iter Interface[T]) Peekable[T] {
	return &peekableIterator[T]{
		inner:  iter,
		peeked: oxide.None[oxide.Option[T]](),
	}
}

type intersperseIterator[T any] struct {
	inner    Peekable[T]
	sep      T
	needsSep bool
}

func (i *intersperseIterator[T]) Next() (T, bool) {
	// The inner iterator is guaranteed to be Peekable due to the type logic in
	// the Intersperse constructor.
	if i.needsSep && i.inner.(Peekable[T]).Peek().IsSome() {
		i.needsSep = false
		return i.sep, true
	}

	i.needsSep = true
	return i.inner.Next()
}

// Intersperse returns a new iterator which injects a copy of the provided
// separator between items yielded by the provided iterator.
func Intersperse[T any](iter Interface[T], sep T) Interface[T] {
	return &intersperseIterator[T]{
		inner:    IntoPeekable(iter),
		sep:      sep,
		needsSep: false,
	}
}
