package main

import "fmt"

func main() {
	sli := []string{"a", "b", "c", "d", "e"}
	sli = append(sli[:3], sli[4:]...)
	fmt.Println(sli)
}
