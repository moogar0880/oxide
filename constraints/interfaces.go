package constraints

// An Integer is a generic type which accounts for all the whole number types
// supported by the stdlib.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// A Number is a generic type which accounts for all the builtin numerical
// types.
type Number interface {
	Integer | ~float32 | ~float64
}
