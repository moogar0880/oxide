package iter

import (
	"testing"

	"github.com/moogar0880/oxide/assert"
)

func TestFromMap(t *testing.T) {
	data := map[int]string{
		0: "foo",
		1: "bar",
		2: "baz",
	}
	expect := []MapEntry[int, string]{
		{
			Key: 0,
			Val: "foo",
		},
		{
			Key: 1,
			Val: "bar",
		},
		{
			Key: 2,
			Val: "baz",
		},
	}

	slice := CollectSlice(FromMap(data))
	assert.Equal(t, len(expect), len(slice))
}

func TestFromChannel(t *testing.T) {
	channel := make(chan int, 10)

	channel <- 0
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	channel <- 5
	channel <- 6
	channel <- 7
	channel <- 8
	channel <- 9

	iter := FromChannel(channel)
	_, upper := iter.(SizeHinter).SizeHint()
	assert.Equal(t, true, upper.IsSome())
	assert.Equal(t, int64(10), upper.Value())

	slice := CollectSlice(iter)
	assert.Equal(t, 10, len(slice))
}

func TestFromNilChan(t *testing.T) {
	var channel chan int

	iter := FromChannel(channel)
	lower, upper := iter.(SizeHinter).SizeHint()
	assert.Equal(t, true, upper.IsNone())
	assert.Equal(t, int64(0), lower)
}

func TestRange(t *testing.T) {
	testIO := []struct {
		name   string
		from   int
		to     int
		expect int
	}{
		{
			name:   "should yield zero elements",
			from:   0,
			to:     0,
			expect: 0,
		},
		{
			name:   "should yield one element",
			from:   0,
			to:     1,
			expect: 1,
		},
		{
			name:   "should yield 10 elements",
			from:   0,
			to:     9,
			expect: 9,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			count := Count(Range(test.from, test.to))
			assert.Equal(t, test.expect, count)
		})
	}
}
