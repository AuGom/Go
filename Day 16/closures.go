package main

import "fmt"

// the function value captures its own i value,
// which will be updated each time we call nextInt
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique to that particular function
	newInts := intSeq()
	fmt.Println(newInts())
}
