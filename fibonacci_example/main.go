package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		temp := curr
		curr += prev
		prev = temp
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
