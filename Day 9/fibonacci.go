// Generates fibonacci sequence for n numbers

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter how many numbers you want in the sequence:")
	var val int
	fmt.Scanf("%d", &val)
	a, b := 0, 1
	fmt.Printf("%d ", a)
	fmt.Printf("%d ", b)
	for i := 0; i < val-2; i++ {
		fmt.Printf("%d ", a+b)
		a, b = b, a+b
	}

}
