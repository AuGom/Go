package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, x := range arr { //1st value==pos and 2nd value==value
		fmt.Println(x)
	}
	fmt.Printf("%T", arr)
}
