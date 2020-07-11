package main

import "fmt"

// Slices are similar to arrays, but with some differences
// 1) slices support new items
// 2) they are typed only by the elements they contain
func main() {

	s := make([]string, 3) // declaring a slice typed as string and length 3
	// make is being used here to not declare a slice without length
	fmt.Println("empty s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("populated s:", s)
	fmt.Println("unique position:", s[2])
	fmt.Println("length s:", len(s))

	// appending new items
	// append returns a slice with the new items added
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("new s:", s)

	c := make([]string, len(s))
	copy(c, s) //copying a slice
	fmt.Println("copy:", c)

	t := []string{"g", "h", "i"} // declaring and populating
	fmt.Println("slice t:", t)

	// making a "slice" of the slice
	// first item (2) is included, last item (5) is excluded
	l := s[2:5]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2:", l)

	l = s[2:]
	fmt.Println("slice 3:", l)

	// like arrays, slices can be used in more than one dimension
	// and, on the contrary of the arrays, the inner slices can vary its length
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d slice: ", twoD)
}
