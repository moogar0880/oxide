package iter

import "github.com/moogar0880/oxide"

type Interface[T any] interface {
	Next() (T, bool)
}

// SizeHinter defines an optional interface that an iterator may implement in
// order to express the lower and, optionally, the upper bounds on it's
// underlying collection.
type SizeHinter interface {
	SizeHint() (int64, oxide.Option[int64])
}

type Peekable[T any] interface {
	Interface[T]

	Peek() oxide.Option[T]
}

type MapEntry[K comparable, V any] struct {
	Key K
	Val V
}
