package iterator_test

import (
	"fmt"

	"github.com/moogar0880/oxide/iterator"
)

func ExampleIterator_Count() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		fmt.Println(iter.Count())
	}
	// Output: 10
}

func ExampleIterator_Last() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.Last to get the last elemnt from the iterator.
		fmt.Println(iter.Last())
	}
	// Output: 9
}

func ExampleIterator_AdvanceBy() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.AdvanceBy to advance the state of our iterator.
		iter = iter.AdvanceBy(5)

		// Use iter.CollectSlice to consume our iterator into a slice of the
		// values yielded by the iterator.
		slice := iter.CollectSlice()
		fmt.Println(slice)
	}
	// Output: [5 6 7 8 9]
}

func ExampleIterator_Nth() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		value, ok := iter.Nth(5)
		fmt.Println(value, ok)
	}
	// Output: 5 true
}

func ExampleIterator_StepBy() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.StepBy to ensure our iterator will have a configured step.
		iter = iter.StepBy(3)

		// Use iter.CollectSlice to consume our iterator into a slice of the
		// values yielded by the iterator.
		slice := iter.CollectSlice()
		fmt.Println(slice)
	}
	// Output: [0 3 6 9]
}

func ExampleIterator_ForEach() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		fmt.Printf("[")
		iter.ForEach(func(i *int) {
			fmt.Printf("%d ", *i)
		})
		fmt.Printf("]\n")
	}
	// Output: [0 1 2 3 4 5 6 7 8 9 ]
}

func ExampleIterator_Find() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		value, found := iter.Find(func(i *int) bool {
			return *i == 7
		})
		fmt.Println(value, found)
	}
	// Output: 7 true
}

func ExampleIterator_SkipWhile() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		slice := iter.SkipWhile(isNegative).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [0 1 2 3 4 5]
}

func ExampleIterator_TakeWhile() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		slice := iter.TakeWhile(isNegative).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [-5 -4 -3 -2 -1]
}

func ExampleIterator_Skip() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.Skip(3).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [3 4 5]
}

func ExampleIterator_Take() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.Take(3).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [0 1 2]
}

func ExampleIterator_Inspect() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5})

		handler := func(i *int) { fmt.Println(*i) }
		slice := iter.Inspect(handler).CollectSlice()
		fmt.Println(slice)
	}
	// Output: 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// [0 1 2 3 4 5]
}

func ExampleIterator_Partition() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		neg, pos := iter.Partition(isNegative)
		fmt.Println(neg, pos)
	}
	// Output: [-5 -4 -3 -2 -1] [0 1 2 3 4 5]
}

func ExampleIterator_Chain() {
	{
		// Define an iterator over our pre-defined data.
		i1 := iterator.FromSlice([]int{-5, -4, -3, -2, -1})
		i2 := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5})
		slice := i1.Chain(i2).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [-5 -4 -3 -2 -1 0 1 2 3 4 5]
}

func ExampleIterator_All() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		fmt.Println(iter.All(isNegative))
	}
	// Output: false
}

func ExampleIterator_Any() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		fmt.Println(iter.Any(isNegative))
	}
	// Output: true
}

func ExampleIterator_Peekable() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})
		peekable := iter.Peekable()

		maybeValue := peekable.Peek()
		fmt.Println(maybeValue.Unpack())

		value, ok := peekable.Next()
		fmt.Println(value, ok)
	}
	// Output: -5 true
	// -5 true
}

func ExampleIterator_Intersperse() {
	{
		// Define an iterator over our pre-defined data.
		iter := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.Intersperse(100).CollectSlice()
		fmt.Println(slice)
	}
	// Output: [0 100 1 100 2 100 3 100 4 100 5]
}

func ExampleIterator() {
	{
		// Generate an iterator from a pre-defined slice.
		slice := iterator.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}).
			// Use the Inspect method to debug values at any point in the
			// iterators chain of operations.
			Inspect(func(i *int) {
				fmt.Printf("inspected: %d\n", *i)
			}).
			// Only consume every other value after the first value.
			StepBy(3).
			// Drop any values which are not even.
			Filter(func(i *int) bool {
				return *i%2 == 0
			}).
			// Inject a value of 100 between every pair of values yielded by
			// the iterator.
			Intersperse(100).
			// Use the Inspect method one last time to debug the output values
			// of our iteration.
			Inspect(func(i *int) {
				fmt.Printf("inspected: %d\n", *i)
			}).
			// Collect the results back into slice.
			CollectSlice()

		fmt.Println(slice)
	}
	// Output: inspected: 0
	// inspected: 0
	// inspected: 1
	// inspected: 2
	// inspected: 3
	// inspected: 4
	// inspected: 5
	// inspected: 6
	// inspected: 100
	// inspected: 6
	// inspected: 7
	// inspected: 8
	// inspected: 9
	// [0 100 6]
}
