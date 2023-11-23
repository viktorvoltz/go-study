package main

import "fmt"

func main() {
	x := 5

	y := increment(x)

	fmt.Println(y)
}

func increment(x int) int {
	x++
	return x
}
