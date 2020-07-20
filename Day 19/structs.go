package main

import "fmt"

// structs are typed collections of fields
type person struct {
	name string
	age  int
}

// constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p

}

func main() {

	fmt.Println(person{"Bob", 20})              // creates a new struct
	fmt.Println(person{name: "Alice", age: 30}) // can name the fields
	fmt.Println(person{name: "Fred"})           // empty fields will be zero-valued
	fmt.Println(&person{name: "Ann", age: 40})  // & prefix yields a pointer to the struct
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // access struct fields with a dot

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51 // structs are mutable
	fmt.Println(sp.age)

}
