package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(sli[:3])
	fmt.Println(sli[4:])
	fmt.Println(sli[1:6])
	fmt.Println(sli[3:9])
}
