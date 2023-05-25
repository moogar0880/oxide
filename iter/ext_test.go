package iter

import (
	"strconv"
	"testing"

	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/assert"
)

func TestMap(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     MapFunc[int, string]
		expect []string
	}{
		{
			name:   "should map all values in list of 1",
			data:   []int{1},
			fn:     strconv.Itoa,
			expect: []string{"1"},
		},
		{
			name:   "should map all values in list of 10",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn:     strconv.Itoa,
			expect: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
		{
			name:   "should map all values in empty list",
			data:   []int{},
			fn:     strconv.Itoa,
			expect: []string{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := FromSlice(test.data)

			actual := CollectSlice(Map(iter, test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TryParseInt(s string) oxide.Option[int] {
	value, err := strconv.Atoi(s)
	if err != nil {
		return oxide.None[int]()
	}

	return oxide.Some(value)
}

func TestFilterMap(t *testing.T) {
	testIO := []struct {
		name   string
		data   []string
		fn     FilterMapFunc[string, int]
		expect []int
	}{
		{
			name:   "should filter map all values in list of 1",
			data:   []string{"1"},
			fn:     TryParseInt,
			expect: []int{1},
		},
		{
			name:   "should filter map over multiple values, some will be filtered",
			data:   []string{"1", "two", "3", "FoUr", "V", "VI", "7", "8", "9", "X"},
			fn:     TryParseInt,
			expect: []int{1, 3, 7, 8, 9},
		},
		{
			name:   "should filter map all values in empty list",
			data:   []string{},
			fn:     func(_ string) oxide.Option[int] { return oxide.None[int]() },
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := FromSlice(test.data)

			actual := CollectSlice(FilterMap(iter, test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestEnumerate(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect []Enumerated[int]
	}{
		{
			name: "should enumerate collection of length 1",
			data: []int{1},
			expect: []Enumerated[int]{
				{0, 1},
			},
		},
		{
			name: "should enumerate collections of longer lengths",
			data: []int{-5, -4, -3, -2, -1},
			expect: []Enumerated[int]{
				{0, -5},
				{1, -4},
				{2, -3},
				{3, -2},
				{4, -1},
			},
		},
		{
			name:   "should enumerate empty collection",
			data:   []int{},
			expect: []Enumerated[int]{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Enumerate(FromSlice(test.data)))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestMapWhile(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     FilterMapFunc[int, int]
		expect []int
	}{
		{
			data: []int{-1, 4, 0, 1},
			fn: func(i int) oxide.Option[int] {
				if i == 0 {
					return oxide.None[int]()
				}

				return oxide.Some(16 / i)
			},
			expect: []int{-16, 4},
		},
		{
			name:   "should map all values in empty list",
			data:   []int{},
			fn:     func(i int) oxide.Option[int] { return oxide.Some(i) },
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := FromSlice(test.data)

			actual := CollectSlice(MapWhile(iter, test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestFuse(t *testing.T) {
	testIO := []struct {
		name   string
		iter   Interface[int]
		expect []int
	}{
		{
			iter:   FromSlice([]int{-1, 4, 0, 1}),
			expect: []int{-1, 4, 0, 1},
		},
		{
			iter:   FromSlice([]int{}),
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := Fuse(test.iter)
			actual := CollectSlice(iter)

			assert.Equal(t, test.expect, actual)

			// Next should now always return no more values.
			for i := 0; i < 10; i++ {
				_, ok := iter.Next()
				assert.Equal(t, false, ok)
			}
		})
	}
}

func TestFindMap(t *testing.T) {
	testIO := []struct {
		name   string
		input  []string
		expect oxide.Option[int]
	}{
		{
			input:  []string{"lol", "NaN", "2", "5"},
			expect: oxide.Some(2),
		},
		{
			input:  []string{"lol", "NaN", "foo", "bar"},
			expect: oxide.None[int](),
		},
		{
			input:  []string{},
			expect: oxide.None[int](),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := FromSlice(test.input)

			actual := FindMap(iter, func(value string) oxide.Option[int] {
				result, err := strconv.Atoi(value)
				if err != nil {
					return oxide.None[int]()
				}

				return oxide.Some(result)
			})

			assert.Equal(t, test.expect, actual)
		})
	}
}
