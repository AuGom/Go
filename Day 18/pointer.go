package main

import "fmt"

// go supports pointers, allowing you to pass references to values and
// records within your program
func ptr(iptr *int) {
	*iptr = 0 // assigning a new value to the referenced variable
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	ptr(&i)
	fmt.Println("zeroptr:", i)  // prints the new value of i
	fmt.Println("pointer:", &i) // prints the memory position of i
}
