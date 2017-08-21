package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	pprev := 0
	prev := 1
	return func() int {
		switch count {
			case 0:
				count++
				return 0
			case 1:
				count++
				return 1
			default:
				count++
				res := pprev+prev
				pprev = prev
				prev = res
				return res
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
