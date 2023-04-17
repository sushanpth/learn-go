package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// first, second, third
	// third = first + second
	// first = third - second
	first := 0
	second := 1
	return func() int {
		// short form
		first, second = second, first+second
		return second - first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
