package math

import (
	"testing"

	"github.com/moogar0880/oxide/assert"
	"github.com/moogar0880/oxide/iter"
)

func TestSum(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should sum 1 number",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should sum 3 numbers",
			data:   []int{1, 2, 3},
			expect: 6,
		},
		{
			name:   "should sum 5 numbers",
			data:   []int{1, 2, 3, 4, 5},
			expect: 15,
		},
		{
			name:   "should sum 0 numbers",
			data:   []int{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Sum(iter.FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestProduct(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should multiply 1 number",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should multiply 3 numbers",
			data:   []int{1, 2, 3},
			expect: 6,
		},
		{
			name:   "should multiply 5 numbers",
			data:   []int{1, 2, 3, 4, 5},
			expect: 120,
		},
		{
			name:   "should multiply 0 numbers",
			data:   []int{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Product(iter.FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestMin(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should find min from 1 number",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should find min from 3 numbers",
			data:   []int{3, 2, 1},
			expect: 1,
		},
		{
			name:   "should find min from 5 numbers",
			data:   []int{5, 7, 6, 2},
			expect: 2,
		},
		{
			name:   "should find min from 0 numbers",
			data:   []int{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Min(iter.FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestMax(t *testing.T) {
	testIO := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "should find max from 1 number",
			data:   []int{1},
			expect: 1,
		},
		{
			name:   "should find max from 3 numbers",
			data:   []int{3, 2, 1},
			expect: 3,
		},
		{
			name:   "should find max from 5 numbers",
			data:   []int{5, 7, 6, 2},
			expect: 7,
		},
		{
			name:   "should find max from 0 numbers",
			data:   []int{},
			expect: 0,
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			actual := Max(iter.FromSlice(test.data))
			assert.Equal(t, test.expect, actual)
		})
	}
}

func BenchmarkSum(b *testing.B) {
	iterator := iter.Range(0, b.N)
	for i := 0; i < b.N; i++ {
		_ = Sum(iterator)
	}
}

func BenchmarkProduct(b *testing.B) {
	iterator := iter.Range(0, b.N)
	for i := 0; i < b.N; i++ {
		_ = Product(iterator)
	}
}

func BenchmarkMin(b *testing.B) {
	iterator := iter.Range(0, b.N)
	for i := 0; i < b.N; i++ {
		_ = Min(iterator)
	}
}

func BenchmarkMax(b *testing.B) {
	iterator := iter.Range(0, b.N)
	for i := 0; i < b.N; i++ {
		_ = Max(iterator)
	}
}
