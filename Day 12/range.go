package main

import "fmt"

// range iterates over elements in a variety of data structures
func main() {

	nums := []int{1, 2, 3}
	sum := 0

	// for each number in range of the "nums" slice its executed the sum
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// it provides index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// on map iterates index and value too
	kvs := map[string]string{"a": "blue", "b": "red", "c": "yellow"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// or just the key
	for k := range kvs {
		fmt.Println("key:", k)
	}

}
