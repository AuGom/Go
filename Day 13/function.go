package main

import "fmt"

// takes two ints and returns an int
func sum(a int, b int) int {

	return a + b
}

// you may omit the type name for the like-typed parameters up to the final
// parameter that declares the type
func sub(a, b int) int {
	return a - b
}

func main() {

	res := sum(1, 2)
	fmt.Println("1+2 =", res)

	res = sub(3, 2)
	fmt.Println("3-2 =", res)
}
