# oxide

Oxide is an experimental library designed to bring the flexibility of some Rust
APIs to the Go programming language.

## Background

`oxide` as a project initially started because I found myself missing certain
tooling from the rust standard library whenever I was working in Go. Obviously
bridging that gap in its entirety is a non-starter and likely far out of scope 
for any one library. However, one pain point that I thought could be 
particularly helpful, now that Go Generics have been around for a while, is 
Iterators. In the end there were several other types that needed to be 
introduced as well in order to keep the APIs sane. At the onset it was unclear 
to me how plausible this would be given how new generics are to the Go type 
system, but in the end the experiment resulted in a library that I believe is
still fairly useful.

### Caveats

Due to limitations in the Go type system, namely the inability to introduce
type constraints onto method receivers, there are several functions in the
functional `iter` API that can not currently be ported to the higher-level 
`oxide.Iterator` API. However, because `oxide.Iterator` implements 
`iter.Interface`, you are able to interoperate between the two APIs fairly 
easily.

## APIs

### `oxide.Option`

The standard `Option` type that Rust programmers are likely accustomed to. This
type allows for concise checks around when an optional value is or is not set,
or in the parlance of the API - when a value is "Some" or "None".

```go
package main

import (
	"fmt"
	
	"github.com/moogar0880/oxide"
)

func main() {
	x := oxide.Some(5)

	if x.IsSome() {
        fmt.Printf("x = oxide.Some(%d)\n", x.Value())
    } else {
		fmt.Println("x is None")
    }
}
```

### `oxide/iter` and `oxide/iterator`

This library offers two different iterator APIs. There's a functional iterator
API which can be found in the `github.com/moogar0880/oxide/iter` module as well
as a higher-level `Iterator` API in `github.com/moogar0880/oxide`.

```go
package main

import (
	"fmt"

	"github.com/moogar0880/oxide/iterator"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
	slice := iterator.FromSlice(data).Filter(func(i *int) bool{
	    return *i%2 == 0
    }).CollectSlice()
	fmt.Println(slice) // []int{0, 2, 4, 6, 8, 10}
}
```

### `oxide/iter` - Functional API 

The `github.com/moogar0880/oxide/iter` module provides a functional iterator API
which provides the foundation for the higher-level `oxide.Iterator` API 
mentioned above. In addition to the methods available on the `oxide.Iterator` 
type there are several functions available in the `iter` package that can not
currently be ported to the `oxide.Iterator` API. These APIs are all currently
located in the `ext.go` file of the `iter` module.

```go
package main

import (
	"fmt"

	"github.com/moogar0880/oxide/iter"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
	slice := iter.CollectSlice(iter.Filter(iter.FromSlice(data), func(i *int) bool{
	    return *i%2 == 0
    }))
	fmt.Println(slice) // []int{0, 2, 4, 6, 8, 10}
}
```

### `oxide/math`

The math module provides several utility functions for performing various 
mathematical operations.

```go
import (
	"fmt"

	"github.com/moogar0880/oxide/iter"
    "github.com/moogar0880/oxide/math"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	sum := math.Sum(iter.FromSlice(data))
	fmt.Println(sum) // 55
}
```
