package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, x := range sli {
		fmt.Println(x)
	}
	fmt.Printf("%T", sli)
}
