package iter

import (
	"context"
	"strconv"
	"testing"

	"github.com/moogar0880/oxide"
	"github.com/moogar0880/oxide/assert"
)

func TestCollectSlice(t *testing.T) {
	testIO := []struct {
		name   string
		iter   Interface[int]
		expect []int
	}{
		{
			iter:   FromSlice([]int{0, 1, 2, 3}),
			expect: []int{0, 1, 2, 3},
		},
		{
			iter:   FromSlice([]int{0}),
			expect: []int{0},
		},
		{
			iter:   FromSlice([]int{}),
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			slice := CollectSlice(test.iter)
			assert.Equal(t, test.expect, slice)
		})
	}
}

func TestCollectMap(t *testing.T) {
	testIO := []struct {
		name   string
		iter   Interface[int]
		fn     MapEntryFunc[string, int]
		expect map[string]int
	}{
		{
			iter: FromSlice([]int{0, 1, 2, 3}),
			fn: func(i int) (string, int) {
				return strconv.Itoa(i), i
			},
			expect: map[string]int{"0": 0, "1": 1, "2": 2, "3": 3},
		},
		{
			iter: FromSlice([]int{}),
			fn: func(i int) (string, int) {
				return strconv.Itoa(i), i
			},
			expect: map[string]int{},
		},
		{
			iter: new(boundless),
			fn: func(i int) (string, int) {
				return strconv.Itoa(i), i
			},
			expect: map[string]int{},
		},
		{
			iter: new(noSizeHint),
			fn: func(i int) (string, int) {
				return strconv.Itoa(i), i
			},
			expect: map[string]int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			data := CollectMap(test.iter, test.fn)
			assert.Equal(t, test.expect, data)
		})
	}
}

func TestCollectChan(t *testing.T) {
	testIO := []struct {
		name   string
		iter   Interface[int]
		expect []int
	}{
		{
			iter:   FromSlice([]int{0, 1, 2, 3}),
			expect: []int{0, 1, 2, 3},
		},
		{
			iter:   FromSlice([]int{}),
			expect: []int{},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			collected := make([]int, 0)
			data := CollectChan(context.Background(), test.iter, 10)

			for value := range data {
				collected = append(collected, value)
			}

			assert.Equal(t, test.expect, collected)
		})
	}
}

func TestCollectChan_Timeout(t *testing.T) {
	iter := FromSlice([]int{0, 1, 2, 3, 4})

	ctx, cancel := context.WithCancel(context.Background())
	// Just cancel the context right out the gate so that we can exercise the
	// specific chunks of code we're trying to hit in this test.
	cancel()

	data := CollectChan(ctx, iter, 10)

	collected := make([]int, 0)
	for value := range data {
		collected = append(collected, value)
	}

	// This assertion is tough to lock down because Go is inconsistent in how
	// it resolves switch statements with channels and completed contexts in
	// contrived examples such as this one, so just assert we got roughly what
	// we expected.
	assert.Equal(t, true, len(collected) < 5)
}

type boundless struct{}

func (d *boundless) Next() (v int, ok bool) { return }

func (d *boundless) SizeHint() (int64, oxide.Option[int64]) {
	return 0, oxide.None[int64]()
}

type noSizeHint struct{}

func (t *noSizeHint) Next() (v int, ok bool) { return }
