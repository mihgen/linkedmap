## Doubly Linked Hash Table
This implementation uses built-in map, and extends it to track order in which (key, value) pairs are added.

This library, and built-in map type are not safe for concurrent use.

#### Example usage
```go
package main

import (
	"fmt"
	"github.com/mihgen/linkedmap"
)

func main() {
	lm := linkedmap.New()

	lm.Add("key1", "value1")
	lm.Add("key2", "value2")
	lm.Add("key3", "value3")
	lm.Add("key2", "00updated00") // update doesn't change order

	// Show value of previously added k-v pair
	fmt.Println("Value of prev added k-v pair: ", lm.Last().Prev().Value())

	// Get value knowing a key
	fmt.Println("Value for key1 is", lm.Get("key1"))

	// List all stored (k,v) pairs
	for e := lm.First(); e != nil; e = e.Next() {
		fmt.Println(e.Key(), "->", e.Value())
	}

	// Pairs before the first and after the last are nil
	fmt.Println("Before the first -", lm.First().Prev())
	fmt.Println("After the last -", lm.Last().Next())

}
```

#### TODO:
- tests
- lazyinit, so that lm := new(linkedmap) works
- delete k,v pair
- check if we need to cover any other methods, available for map built-in
