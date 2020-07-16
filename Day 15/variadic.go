package main

import "fmt"

// variadic functions can take an arbitrary number of arguments
func sum(vals ...int) {
	fmt.Print(vals, " ")
	total := 0
	for _, num := range vals {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1)
	sum(1, 2)
	sum(1, 2, 3)

	// it can be applied to slices too
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
