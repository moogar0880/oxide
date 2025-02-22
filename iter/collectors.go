package iter

import (
	"context"
	"fmt"
	"strings"
)

// CollectSlice consumes the provided Interface into a slice of type T.
func CollectSlice[T any](from Interface[T]) []T {
	slice := make([]T, 0)
	for item, ok := from.Next(); ok; item, ok = from.Next() {
		slice = append(slice, item)
	}

	return slice
}

type MapEntryFunc[K comparable, V any] func(V) (K, V)

// CollectMap consumes the provided Interface, converting each yielded value
// into a key-value pair that are inserted into the generated map.
func CollectMap[K comparable, V any](iter Interface[V], fn MapEntryFunc[K, V]) (out map[K]V) {
	if sized, ok := iter.(SizeHinter); ok {
		if _, upper := sized.SizeHint(); upper.IsSome() {
			out = make(map[K]V, upper.Value())
		} else {
			out = make(map[K]V)
		}
	} else {
		out = make(map[K]V)
	}

	var key K
	var val V
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		key, val = fn(item)

		out[key] = val
	}
	return out
}

// CollectChan consumes the provided Interface and writes each of it's yielded
// values onto the returned channel.
//
// Note: once all values have been consumed from the Interface, the returned
// channel will be closed.
func CollectChan[T any](ctx context.Context, from Interface[T], buffer int) <-chan T {
	out := make(chan T, buffer)
	go func(out chan<- T) {
		defer close(out)

		for item, ok := from.Next(); ok; item, ok = from.Next() {
			select {
			case out <- item:
			case <-ctx.Done():
				return
			}
		}
	}(out)
	return out
}

// Join consumes all values in the provided iterator and concatenates the
// elements to create a single string.
//
// The separator string sep is placed between elements in the resulting string.
func Join(iter Interface[string], sep string) string {
	return strings.Join(CollectSlice(iter), sep)
}

// JoinStringer behaves similarly to Join but accepts any value which can be
// converted to a string via the fmt.Stringer interface.
//
// For additional details see Join and fmt.Stringer.
func JoinStringer[T fmt.Stringer](iter Interface[T], sep string) string {
	return Join(Map(iter, func(f T) string {
		return f.String()
	}), sep)
}
