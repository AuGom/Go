package main

import "fmt"

func main() {
	for count := 33; count < 123; count++ {
		fmt.Printf("%d\t%#x\t%s\n", count, count, string(count))
	}
}
