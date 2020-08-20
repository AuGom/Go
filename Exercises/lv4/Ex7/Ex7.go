package main

import "fmt"

func main() {
	x := [][]string{
		[]string{
			"aaa",
			"bbb",
			"ccc",
		},
		[]string{
			"ddd",
			"eee",
			"fff",
		},
		[]string{
			"ggg",
			"hhh",
			"iii",
		},
	}

	for _, v := range x {
		fmt.Println(v)
	}
}
