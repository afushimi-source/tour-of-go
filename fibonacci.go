package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n, next_n := 0, 1
	return func() int {
		now_n := n
		n, next_n = next_n, n + next_n
		return now_n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
