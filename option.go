package oxide

// An Option is the structural representation of the common tuple API that
// signifies the presence, or lack thereof, of a value.
//
// For example, let's say we have this simple getter method:
//
//	func Get(data map[string]string, key string) (string, bool) {
//		value, ok := data[key]
//		return value, ok
//	}
//
// Although quite trivial, this could be rewritten to leverage the Option type
// like so:
//
//	func Get(data map[string]string, key string) Option[string] {
//		if value, ok := data[key]; ok {
//			return Some(value)
//		}
//
//		return None[string]()
//	}
type Option[T any] struct {
	value   T
	present bool
}

// Some returns an Option value which wraps the provided value as a "Some"
// variant.
//
// Options generated with this constructor will always return `true` for
// calls to `IsSome`, and `false` for calls to `IsNone`.
func Some[T any](value T) Option[T] {
	return Option[T]{value: value, present: true}
}

// None returns an Option value which represents a "None" variant for the inner
// type.
//
// Options generated with this constructor will always return `false` for
// calls to `IsSome`, and `true` for calls to `IsNone`.
func None[T any]() Option[T] {
	return Option[T]{present: false}
}

// IsSome returns true if this Option represents a "Some" value.
func (o Option[T]) IsSome() bool {
	return o.present
}

// IsNone returns true if this Option represents a "None" value.
func (o Option[T]) IsNone() bool {
	return !o.present
}

// Value returns the inner value represented by this Option instance.
//
// If the Option is a "Some" variant, then the raw value will be returned.
//
// If the Option is a "None" variant, then a zero value for type T will be
// returned.
func (o Option[T]) Value() T {
	return o.value
}

// Unpack deconstructs the Option struct into the more idiomatic optional
// tuple representation of `(T, bool)`.
//
// This can be useful for bridging between this API and APIs implemented in
// idiomatic go.
func (o Option[T]) Unpack() (T, bool) {
	return o.value, o.present
}
