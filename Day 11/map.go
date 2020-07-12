package main

import "fmt"

func main() {

	// use make to create an empty map: make(map[key-type]val-type)
	m := make(map[string]int)

	// declaring and initializing a map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map N:", n)

	// setting key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map M:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len M:", len(m))

	// delete method removes a key-value pair from a map
	delete(m, "k2")
	fmt.Println("map M:", m)

	// by getting a map position in this way its possible to know if it really
	// exists on the map, if you use the first return too you can get the value
	// pointed by the key or a null value if the key/position doesn't exist
	_, prs := m["k1"]
	fmt.Println("prs:", prs)
}
