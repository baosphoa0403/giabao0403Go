package main

import "fmt"

func main() {
	fmt.Printf("hello")
	abc := Sum(1, 2)

	fmt.Printf("sum: %d\n", abc)
}

// sum two number
func Sum(a int, b int) int {
	return a + b
}
