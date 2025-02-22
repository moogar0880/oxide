package iterator

import (
	"context"
	"testing"
	"time"

	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/assert"
	"github.com/moogar0880/oxide/iter"
)

func TestIterator(t *testing.T) {
	var slice iter.Interface[int]
	slice = FromSlice([]int{0, 1, 2})

	var hashMap iter.Interface[iter.MapEntry[string, int]]
	hashMap = FromMap(map[string]int{"0": 0, "1": 1, "2": 2})

	var channel iter.Interface[int]
	channel = FromChan(func() chan int {
		out := make(chan int)
		go func() {
			out <- 0
			out <- 1
			out <- 2
		}()
		return out
	}())

	_, ok := slice.(iter.Interface[int])
	assert.Equal(t, true, ok)

	_, ok = hashMap.(iter.Interface[iter.MapEntry[string, int]])
	assert.Equal(t, true, ok)

	_, ok = channel.(iter.Interface[int])
	assert.Equal(t, true, ok)
}

func TestIterator_CollectChan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	channel := FromSlice([]int{0, 1, 2}).CollectChan(ctx, 3)

	collected := make([]int, 0)
	for value := range channel {
		collected = append(collected, value)
	}

	assert.Equal(t, 3, len(collected))
}

func TestIterator_Count(t *testing.T) {
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
			actual := FromSlice(test.data).Count()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Last(t *testing.T) {
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
			actual := FromSlice(test.data).Last()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_AdvanceBy(t *testing.T) {
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
			actual, _ := FromSlice(test.data).AdvanceBy(test.advanceBy).Next()
			assert.Equal(t, test.expect, actual)

			iter := FromSlice(test.data)
			iter = iter.AdvanceBy(test.advanceBy)
			actual, _ = iter.Next()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Nth(t *testing.T) {
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
			actual, _ := FromSlice(test.data).Nth(test.n)
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_StepBy(t *testing.T) {
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
			actual := FromSlice(test.data).StepBy(test.stepBy).CollectSlice()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_ForEach(t *testing.T) {
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

			FromSlice(test.data).ForEach(func(_ *int) {
				visited = visited + 1
			})

			assert.Equal(t, test.expect, visited)
		})
	}
}

func TestIterator_Find(t *testing.T) {
	testIO := []struct {
		name     string
		data     []int
		fn       iter.Predicate[int]
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
			actual, ok := FromSlice(test.data).Find(test.fn)

			assert.Equal(t, test.expect, actual)
			assert.Equal(t, test.expectOk, ok)
		})
	}
}

func TestIterator_Filter(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     iter.Predicate[int]
		expect []int
	}{
		{
			name:   "should handle finding an item in a list of 1",
			data:   []int{0, 1, 2},
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
			actual := FromSlice(test.data).Filter(test.fn).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_SkipWhile(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     iter.Predicate[int]
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
			data: []int{2, 3, 4, 5, 6},
			fn: func(i *int) bool {
				return *i%2 == 0
			},
			expect: []int{3, 4, 5, 6},
		},
		{
			name: "should handle not finding an item",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn: func(i *int) bool {
				return *i == 100
			},
			expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "should handle not finding an item",
			data: []int{},
			fn: func(i *int) bool {
				return *i == 100
			},
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := FromSlice(test.data).SkipWhile(test.fn).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_TakeWhile(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     iter.Predicate[int]
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
			actual := FromSlice(test.data).TakeWhile(test.fn).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Skip(t *testing.T) {
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
			actual := FromSlice(test.data).Skip(test.input).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Take(t *testing.T) {
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
			actual := FromSlice(test.data).Take(test.input).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Inspect(t *testing.T) {
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

			collected := FromSlice(test.data).Inspect(func(_ *int) {
				visited = visited + 1
			}).CollectSlice()

			assert.Equal(t, test.expect, visited)
			assert.Equal(t, test.data, collected)
		})
	}
}

func TestIterator_Partition(t *testing.T) {
	testIO := []struct {
		name        string
		data        []int
		fn          iter.Predicate[int]
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
			left, right := FromSlice(test.data).Partition(test.fn)

			assert.Equal(t, test.expectLeft, left)
			assert.Equal(t, test.expectRight, right)
		})
	}
}

func TestIterator_Chain(t *testing.T) {
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
			actual := FromSlice(test.data1).
				Chain(FromSlice(test.data2)).
				CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

// func TestIterator_Zip(t *testing.T) {
// 	testIO := []struct {
// 		name   string
// 		data1  []int
// 		data2  []int
// 		expect [][2]int
// 	}{
// 		{
// 			name:  "should partition even and odds",
// 			data1: []int{1, 2, 3},
// 			data2: []int{4, 5, 6},
// 			expect: [][2]int{
// 				{1, 4},
// 				{2, 5},
// 				{3, 6},
// 			},
// 		},
// 		{
// 			name:  "should partition negative and positive",
// 			data1: []int{-5, -4, -3, -2, -1},
// 			data2: []int{4, 5, 6},
// 			expect: [][2]int{
// 				{-5, 4},
// 				{-4, 5},
// 				{-3, 6},
// 				{-2, 0},
// 				{-1, 0},
// 			},
// 		},
// 	}
//
// 	for _, test := range testIO {
// 		t.Run(test.name, func(t *testing.T) {
// 			actual := FromSlice(test.data1).
// 				Zip(FromSlice(test.data2)).
// 				CollectSlice()
//
// 			assert.Equal(t, test.expect, actual)
// 		})
// 	}
// }

func TestIterator_All(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     iter.Predicate[int]
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
			actual := FromSlice(test.data).All(test.fn)

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Any(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		fn     iter.Predicate[int]
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
			actual := FromSlice(test.data).Any(test.fn)

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_SizeHint(t *testing.T) {
	testIO := []struct {
		name        string
		data        []int
		expectLower int64
		expectUpper oxide.Option[int64]
	}{
		{
			name:        "should find all are even",
			data:        []int{2, 4, 6, 8, 10},
			expectUpper: oxide.Some(int64(5)),
		},
		{
			name:        "should find none are even",
			data:        []int{1, 3, 5},
			expectUpper: oxide.Some(int64(3)),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			lower, maybeUpper := FromSlice(test.data).SizeHint()

			assert.Equal(t, test.expectLower, lower)
			assert.Equal(t, test.expectUpper, maybeUpper)
		})
	}
}

func TestIterator_Peekable(t *testing.T) {
	testIO := []struct {
		name string
		iter *Iterator[int]
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
			iter: FromSlice([]int{}),
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			peekable := test.iter.Peekable()

			for {
				peeked := peekable.Peek()
				value, ok := peekable.Next()

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

func TestIterator_Intersperse(t *testing.T) {
	testIO := []struct {
		name   string
		iter   *Iterator[int]
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
		{
			name:   "should not intersperse into empty iter",
			iter:   FromSlice([]int{}),
			sep:    100,
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := test.iter.Intersperse(test.sep).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Interleave(t *testing.T) {
	testIO := []struct {
		name   string
		iter1  *Iterator[int]
		iter2  iter.Interface[int]
		expect []int
	}{
		{
			name:   "should interleave balanced values",
			iter1:  FromSlice([]int{1, 3, 5, 7, 9}),
			iter2:  FromSlice([]int{2, 4, 6, 8, 10}),
			expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "should interleave unbalanced values (left heavy)",
			iter1:  FromSlice([]int{1, 3, 5}),
			iter2:  FromSlice([]int{2}),
			expect: []int{1, 2, 3, 5},
		},
		{
			name:   "should interleave unbalanced values (right heavy)",
			iter1:  FromSlice([]int{1}),
			iter2:  FromSlice([]int{2, 4, 6}),
			expect: []int{1, 2, 4, 6},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := test.iter1.Interleave(test.iter2).CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestIterator_Sort(t *testing.T) {
	testIO := []struct {
		name     string
		iter     *Iterator[int]
		lessFunc func(i, j int) bool
		expect   []int
	}{
		{
			name:     "should sort empty iterator",
			iter:     FromSlice([]int{}),
			lessFunc: func(i, j int) bool { return i < j },
			expect:   []int{},
		},
		{
			name:     "should sort iterator with single value",
			iter:     FromSlice([]int{1}),
			lessFunc: func(i, j int) bool { return i < j },
			expect:   []int{1},
		},
		{
			name:     "should sort already sorted iterator",
			iter:     FromSlice([]int{1, 2, 3, 4, 5}),
			lessFunc: func(i, j int) bool { return i < j },
			expect:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "should sort unsorted iterator",
			iter:     FromSlice([]int{2, 6, 3, 1, 4, 5}),
			lessFunc: func(i, j int) bool { return i < j },
			expect:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "should sort in reverse order",
			iter:     FromSlice([]int{2, 6, 3, 1, 4, 5}),
			lessFunc: func(i, j int) bool { return i > j },
			expect:   []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			iterator := test.iter.Sorted(test.lessFunc)
			actual := iterator.CollectSlice()

			assert.Equal(t, test.expect, actual)
		})
	}
}
