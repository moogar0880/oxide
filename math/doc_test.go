package math_test

import (
	"fmt"

	"github.com/moogar0880/oxide/iter"
	"github.com/moogar0880/oxide/math"
)

func ExampleMax() {
	{
		// Define an iterator over our pre-defined data.
		max := math.Max(iter.FromSlice([]int{1, -2, -3, 0, 6, 1, 2, 3}))
		fmt.Println(max)
	}
	// Output: 6
}

func ExampleMin() {
	{
		// Define an iterator over our pre-defined data.
		min := math.Min(iter.FromSlice([]int{-1, -2, -3, 0, 1, 2, 3}))
		fmt.Println(min)
	}
	// Output: -3
}

func ExampleProduct() {
	{
		// Define an iterator over our pre-defined data.
		product := math.Product(iter.FromSlice([]int{1, 2, 3, 4, 5}))
		fmt.Println(product)
	}
	// Output: 120
}

func ExampleSum() {
	{
		// Define an iterator over our pre-defined data.
		sum := math.Sum(iter.FromSlice([]int{0, 1, 2, 3, 4, 5}))
		fmt.Println(sum)
	}
	// Output: 15
}
