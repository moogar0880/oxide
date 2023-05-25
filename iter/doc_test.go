package iter_test

import (
	"fmt"

	"github.com/moogar0880/oxide/iter"
)

func ExampleCount() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.Count to get the total count of all items in the iterator.
		count := iter.Count(iterator)
		fmt.Println(count)
	}
	// Output: 10
}

func ExampleLast() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.Last to get the last elemnt from the iterator.
		last := iter.Last(iterator)
		fmt.Println(last)
	}
	// Output: 9
}

func ExampleAdvanceBy() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.AdvanceBy to advance the state of our iterator.
		iterator = iter.AdvanceBy(iterator, 5)

		// Use iter.CollectSlice to consume our iterator into a slice of the
		// values yielded by the iterator.
		slice := iter.CollectSlice(iterator)
		fmt.Println(slice)
	}
	// Output: [5 6 7 8 9]
}

func ExampleNth() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		value, ok := iter.Nth(iterator, 5)
		fmt.Println(value, ok)
	}
	// Output: 5 true
}

func ExampleStepBy() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		// Use iter.StepBy to ensure our iterator will have a configured step.
		iterator = iter.StepBy(iterator, 3)

		// Use iter.CollectSlice to consume our iterator into a slice of the
		// values yielded by the iterator.
		slice := iter.CollectSlice(iterator)
		fmt.Println(slice)
	}
	// Output: [0 3 6 9]
}

func ExampleForEach() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		fmt.Printf("[")
		iter.ForEach(iterator, func(i *int) {
			fmt.Printf("%d ", *i)
		})
		fmt.Printf("]\n")
	}
	// Output: [0 1 2 3 4 5 6 7 8 9 ]
}

func ExampleFind() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		value, found := iter.Find(iterator, func(i *int) bool {
			return *i == 7
		})
		fmt.Println(value, found)
	}
	// Output: 7 true
}

func ExampleSkipWhile() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		slice := iter.CollectSlice(iter.SkipWhile(iterator, isNegative))
		fmt.Println(slice)
	}
	// Output: [0 1 2 3 4 5]
}

func ExampleTakeWhile() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		slice := iter.CollectSlice(iter.TakeWhile(iterator, isNegative))
		fmt.Println(slice)
	}
	// Output: [-5 -4 -3 -2 -1]
}

func ExampleSkip() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.CollectSlice(iter.Skip(iterator, 3))
		fmt.Println(slice)
	}
	// Output: [3 4 5]
}

func ExampleTake() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.CollectSlice(iter.Take(iterator, 3))
		fmt.Println(slice)
	}
	// Output: [0 1 2]
}

func ExampleInspect() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})

		handler := func(i *int) { fmt.Println(*i) }
		slice := iter.CollectSlice(iter.Inspect(iterator, handler))
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

func ExamplePartition() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		neg, pos := iter.Partition(iterator, isNegative)
		fmt.Println(neg, pos)
	}
	// Output: [-5 -4 -3 -2 -1] [0 1 2 3 4 5]
}

func ExampleChain() {
	{
		// Define an iterator over our pre-defined data.
		i1 := iter.FromSlice([]int{-5, -4, -3, -2, -1})
		i2 := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})
		iterator := iter.Chain(i1, i2)

		slice := iter.CollectSlice(iterator)
		fmt.Println(slice)
	}
	// Output: [-5 -4 -3 -2 -1 0 1 2 3 4 5]
}

func ExampleZip() {
	{
		// Define an iterator over our pre-defined data.
		i1 := iter.FromSlice([]int{-5, -4, -3, -2, -1})
		i2 := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.CollectSlice(iter.Zip(i1, i2))
		fmt.Println(slice)
	}
	// Output: [[-5 0] [-4 1] [-3 2] [-2 3] [-1 4] [0 5]]
}

func ExampleAll() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		fmt.Println(iter.All(iterator, isNegative))
	}
	// Output: false
}

func ExampleAny() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})

		isNegative := func(i *int) bool { return *i < 0 }
		fmt.Println(iter.Any(iterator, isNegative))
	}
	// Output: true
}

func ExampleIntoPeekable() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})
		peekable := iter.IntoPeekable(iterator)

		maybeValue := peekable.Peek()
		fmt.Println(maybeValue.Unpack())

		value, ok := peekable.Next()
		fmt.Println(value, ok)
	}
	// Output: -5 true
	// -5 true
}

func ExampleIntersperse() {
	{
		// Define an iterator over our pre-defined data.
		iterator := iter.FromSlice([]int{0, 1, 2, 3, 4, 5})

		slice := iter.CollectSlice(iter.Intersperse(iterator, 100))
		fmt.Println(slice)
	}
	// Output: [0 100 1 100 2 100 3 100 4 100 5]
}
