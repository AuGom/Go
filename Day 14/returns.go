package main

import "fmt"

// the (int, int) indicates multiple returns, in this case two integers
func vals() (int, int) {
	return 3, 4

}

func main() {
	a, b := vals() // a and b will get respectively the return from vals
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // same case as above, but c will get only the value 4
	fmt.Println(c)

}
