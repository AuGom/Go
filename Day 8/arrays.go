package main

import "fmt"

func main() {

	var a [5]int               //declaring an array that hold integers
	b := [5]int{1, 2, 3, 4, 5} //declaring and populating

	a[2] = 100
	fmt.Println("a:", a)             //printing the entire array
	fmt.Println("a[2]:", a[2])       //printing a unique position
	fmt.Println("a length:", len(a)) //the length of the array

	fmt.Println("b:", b)

	//two dimensional array
	var c [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("c: ", c)
}
