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
