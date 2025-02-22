package iter

import (
	"testing"

	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/assert"
)

func TestCount(t *testing.T) {
	testIO := []struct {
		name   string
		data   []string
		expect int
	}{
		{
			name:   "should handle slice of size 1",
			data:   []string{"foo"},
			expect: 1,
		},
		{
			name:   "should handle slice of size 2",
			data:   []string{"foo", "bar"},
			expect: 2,
		},
		{
			name: "should handle slice of size 10",
			data: []string{
				"foo", "bar", "baz",
				"bag", "bad", "bat",
				"baa", "bab", "ban",
				"dan",
			},
			expect: 10,
		},
		{
			name:   "should handle slice of size 0",
			data:   []string{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Count(FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestLast(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should handle slice of size 1",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should handle slice of size 2",
			data:   []int{1, 2},
			expect: 2,
		},
		{
			name:   "should handle slice of size 10",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expect: 10,
		},
		{
			name:   "should handle slice of size 0",
			data:   []int{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Last(FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestAdvanceBy(t *testing.T) {
	testIO := []struct {
		name      string
		data      []int
		advanceBy int
		expect    int
	}{
		{
			name:      "should handle advancing by 0",
			data:      []int{1},
			advanceBy: 0,
			expect:    1,
		},
		{
			name:      "should handle advancing by 1",
			data:      []int{1, 2},
			advanceBy: 1,
			expect:    2,
		},
		{
			name:      "should handle advancing by 5",
			data:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			advanceBy: 5,
			expect:    6,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := AdvanceBy(FromSlice(test.data), test.advanceBy).Next()
			assert.Equal(t, test.expect, actual)

			iter := FromSlice(test.data)
			iter = AdvanceBy(iter, test.advanceBy)
			actual, _ = iter.Next()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestNth(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		n      int
		expect int
	}{
		{
			name:   "should handle advancing by 0",
			data:   []int{1},
			n:      0,
			expect: 1,
		},
		{
			name:   "should handle advancing by 1",
			data:   []int{1, 2},
			n:      1,
			expect: 2,
		},
		{
			name:   "should handle advancing by 5",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:      5,
			expect: 6,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := Nth(FromSlice(test.data), test.n)
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestStepBy(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		stepBy int
		expect []int
	}{
		{
			name:   "should handle stepping by 1",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			stepBy: 1,
			expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "should handle stepping by 2",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			stepBy: 2,
			expect: []int{1, 3, 5, 7, 9},
		},
		{
			name:   "should handle advancing by 5",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			stepBy: 5,
			expect: []int{1, 6},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(StepBy(FromSlice(test.data), test.stepBy))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestForEach(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should handle advancing by 0",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should handle advancing by 1",
			data:   []int{1, 2},
			expect: 2,
		},
		{
			name:   "should handle advancing by 5",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expect: 10,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			visited := 0

			ForEach(FromSlice(test.data), func(_ *int) {
				visited = visited + 1
			})

			assert.Equal(t, test.expect, visited)
		})
	}
}

func TestFind(t *testing.T) {
	testIO := []struct {
		name     string
		data     []int
		fn       Predicate[int]
		expect   int
		expectOk bool
	}{
		{
			name:     "should handle finding an item in a list of 1",
			data:     []int{1},
			fn:       assert.IsPositive,
			expect:   1,
			expectOk: true,
		},
		{
			name:     "should handle finding a value in a larger list",
			data:     []int{1, 2},
			fn:       assert.IsEven,
			expect:   2,
			expectOk: true,
		},
		{
			name:     "should handle not finding an item",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn:       assert.Is100,
			expect:   0,
			expectOk: false,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual, ok := Find(FromSlice(test.data), test.fn)

			assert.Equal(t, test.expect, actual)
			assert.Equal(t, test.expectOk, ok)
		})
	}
}

func TestFilter(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     Predicate[int]
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{-1, 0, 1, 2},
			fn:     assert.IsPositive,
			expect: []int{0, 1, 2},
		},
		{
			name:   "should handle finding a value in a larger list",
			data:   []int{0, 1, 2, 3, 4, 5, 6},
			fn:     assert.IsEven,
			expect: []int{0, 2, 4, 6},
		},
		{
			name:   "should handle not finding an item",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn:     assert.Is100,
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Filter(FromSlice(test.data), test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestSkipWhile(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     Predicate[int]
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			fn:     assert.IsNegative,
			expect: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "should handle finding a value in a larger list",
			data: []int{0, 1, 2, 3, 4, 5, 6},
			fn: func(i *int) bool {
				return *i < 3
			},
			expect: []int{3, 4, 5, 6},
		},
		{
			name: "should handle not finding an item",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn: func(i *int) bool {
				return *i < 100
			},
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(SkipWhile(FromSlice(test.data), test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestTakeWhile(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     Predicate[int]
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			fn:     assert.IsNegative,
			expect: []int{-5, -4, -3, -2, -1},
		},
		{
			name:   "should handle finding a value in a larger list",
			data:   []int{2, 4, 5, 6},
			fn:     assert.IsEven,
			expect: []int{2, 4},
		},
		{
			name:   "should handle not finding an item",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn:     assert.Is100,
			expect: []int{},
		},
		{
			name:   "should consume entire iterator",
			data:   []int{2, 4, 6, 8, 10},
			fn:     assert.IsEven,
			expect: []int{2, 4, 6, 8, 10},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(TakeWhile(FromSlice(test.data), test.fn))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestSkip(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		input  int
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			input:  5,
			expect: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:   "should handle finding a value in a larger list",
			data:   []int{2, 4, 5, 6},
			input:  1,
			expect: []int{4, 5, 6},
		},
		{
			name:   "should handle not finding an item",
			data:   []int{1, 2, 3, 4, 5},
			input:  8,
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Skip(FromSlice(test.data), test.input))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestTake(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		input  int
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			input:  5,
			expect: []int{-5, -4, -3, -2, -1},
		},
		{
			name:   "should handle finding a value in a larger list",
			data:   []int{2, 4, 5, 6},
			input:  1,
			expect: []int{2},
		},
		{
			name:   "should handle not finding an item",
			data:   []int{1, 2, 3, 4, 5},
			input:  8,
			expect: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Take(FromSlice(test.data), test.input))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestInspect(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should handle advancing by 0",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should handle advancing by 1",
			data:   []int{1, 2},
			expect: 2,
		},
		{
			name:   "should handle advancing by 5",
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expect: 10,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			visited := 0

			collected := CollectSlice(Inspect(FromSlice(test.data), func(_ *int) {
				visited = visited + 1
			}))

			assert.Equal(t, test.expect, visited)
			assert.Equal(t, test.data, collected)
		})
	}
}

func TestPartition(t *testing.T) {
	testIO := []struct {
		name        string
		data        []int
		fn          Predicate[int]
		expectLeft  []int
		expectRight []int
	}{
		{
			name:        "should partition even and odds",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn:          assert.IsEven,
			expectLeft:  []int{2, 4, 6, 8, 10},
			expectRight: []int{1, 3, 5, 7, 9},
		},
		{
			name:        "should partition negative and positive",
			data:        []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			fn:          assert.IsPositive,
			expectLeft:  []int{0, 1, 2, 3, 4, 5},
			expectRight: []int{-5, -4, -3, -2, -1},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			left, right := Partition(FromSlice(test.data), test.fn)

			assert.Equal(t, test.expectLeft, left)
			assert.Equal(t, test.expectRight, right)
		})
	}
}

func TestChain(t *testing.T) {
	testIO := []struct {
		name   string
		data1  []int
		data2  []int
		expect []int
	}{
		{
			name:   "should partition even and odds",
			data1:  []int{1, 2, 3},
			data2:  []int{4, 5, 6},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "should partition negative and positive",
			data1:  []int{-5, -4, -3, -2, -1},
			data2:  []int{4, 5, 6},
			expect: []int{-5, -4, -3, -2, -1, 4, 5, 6},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Chain(FromSlice(test.data1), FromSlice(test.data2)))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestZip(t *testing.T) {
	testIO := []struct {
		name   string
		data1  []int
		data2  []int
		expect [][2]int
	}{
		{
			name:  "should zip two collections of equal length",
			data1: []int{1, 2, 3},
			data2: []int{4, 5, 6},
			expect: [][2]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			name:  "should zip collections of different lengths",
			data1: []int{-5, -4, -3, -2, -1},
			data2: []int{4, 5, 6},
			expect: [][2]int{
				{-5, 4},
				{-4, 5},
				{-3, 6},
				{-2, 0},
				{-1, 0},
			},
		},
		{
			name:  "should zip collections of different lengths",
			data1: []int{-3, -2, -1},
			data2: []int{1, 2, 3, 4, 5, 6},
			expect: [][2]int{
				{-3, 1},
				{-2, 2},
				{-1, 3},
				{0, 4},
				{0, 5},
				{0, 6},
			},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Zip(FromSlice(test.data1), FromSlice(test.data2)))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestAll(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     Predicate[int]
		expect bool
	}{
		{
			name:   "should find all are even",
			data:   []int{2, 4, 6, 8, 10},
			fn:     assert.IsEven,
			expect: true,
		},
		{
			name:   "should find all are not even",
			data:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			fn:     assert.IsPositive,
			expect: false,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := All(FromSlice(test.data), test.fn)

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestAny(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     Predicate[int]
		expect bool
	}{
		{
			name:   "should find all are even",
			data:   []int{2, 4, 6, 8, 10},
			fn:     assert.IsEven,
			expect: true,
		},
		{
			name:   "should find none are even",
			data:   []int{1, 3, 5},
			fn:     assert.IsEven,
			expect: false,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Any(FromSlice(test.data), test.fn)

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestSizeHint(t *testing.T) {
	testIO := []struct {
		name        string
		iter        Interface[int]
		expectLower int64
		expectUpper oxide.Option[int64]
	}{
		{
			name:        "should find all are even",
			iter:        FromSlice([]int{2, 4, 6, 8, 10}),
			expectUpper: oxide.Some(int64(5)),
		},
		{
			name:        "should find none are even",
			iter:        FromSlice([]int{1, 3, 5}),
			expectUpper: oxide.Some(int64(3)),
		},
		{
			name:        "should find none are even",
			iter:        new(unboundedIterator),
			expectUpper: oxide.None[int64](),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			lower, maybeUpper := SizeHint(test.iter)

			assert.Equal(t, test.expectLower, lower)
			assert.Equal(t, test.expectUpper, maybeUpper)
		})
	}
}

func TestIntoPeekable(t *testing.T) {
	testIO := []struct {
		name string
		iter Interface[int]
	}{
		{
			name: "should make peekable iterator",
			iter: FromSlice([]int{2, 4, 6, 8, 10}),
		},
		{
			name: "should make peekable iterator",
			iter: FromSlice([]int{1, 3, 5}),
		},
		{
			name: "should make peekable iterator",
			iter: new(unboundedIterator),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iter := IntoPeekable(test.iter)

			for {
				peeked := iter.Peek()

				value, ok := iter.Next()

				if peeked.IsSome() {
					assert.Equal(t, true, ok)
					assert.Equal(t, peeked.Value(), value)
				} else {
					assert.Equal(t, false, ok)
					assert.Equal(t, peeked.Value(), value)
				}

				if !ok {
					break
				}
			}
		})
	}
}

func TestIntersperse(t *testing.T) {
	testIO := []struct {
		name   string
		iter   Interface[int]
		sep    int
		expect []int
	}{
		{
			name:   "should intersperse zeros",
			iter:   FromSlice([]int{2, 4, 6, 8, 10}),
			sep:    0,
			expect: []int{2, 0, 4, 0, 6, 0, 8, 0, 10},
		},
		{
			name:   "should make peekable iterator",
			iter:   FromSlice([]int{1, 3, 5}),
			sep:    100,
			expect: []int{1, 100, 3, 100, 5},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := CollectSlice(Intersperse(test.iter, test.sep))

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestInterleave(t *testing.T) {
	testIO := []struct {
		name       string
		iter1      Interface[int]
		iter2      Interface[int]
		expect     []int
		expectSize oxide.Option[int64]
	}{
		{
			name:       "should interleave balanced values",
			iter1:      FromSlice([]int{1, 3, 5, 7, 9}),
			iter2:      FromSlice([]int{2, 4, 6, 8, 10}),
			expect:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectSize: oxide.Some[int64](10),
		},
		{
			name:       "should interleave unbalanced values (left heavy)",
			iter1:      FromSlice([]int{1, 3, 5}),
			iter2:      FromSlice([]int{2}),
			expect:     []int{1, 2, 3, 5},
			expectSize: oxide.Some[int64](4),
		},
		{
			name:       "should interleave unbalanced values (right heavy)",
			iter1:      FromSlice([]int{1}),
			iter2:      FromSlice([]int{2, 4, 6}),
			expect:     []int{1, 2, 4, 6},
			expectSize: oxide.Some[int64](4),
		},
		{
			name:       "should return no values and have no size hint if both iterators do not implement SizeHint",
			iter1:      new(unboundedIterator),
			iter2:      new(unboundedIterator),
			expect:     []int{},
			expectSize: oxide.None[int64](),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iterator := Interleave(test.iter1, test.iter2)
			lower, size := iterator.(SizeHinter).SizeHint()
			assert.Equal(t, int64(0), lower)
			assert.Equal(t, test.expectSize, size)

			actual := CollectSlice(iterator)

			assert.Equal(t, test.expect, actual)
		})
	}
}

type unboundedIterator struct{}

func (*unboundedIterator) Next() (data int, present bool) { return }
