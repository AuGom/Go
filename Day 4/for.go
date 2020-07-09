package main

import "fmt"

//example of how for is implemented on golang
func main() {
	i := 1       //declare and initialize a variable. Another method is: var i int = 1
	for i <= 3 { //similar to while on other programming languages
		fmt.Println(i)
		i = i + 1
	}

	for j := 10; j >= 5; j-- { //how its commonly seen
		fmt.Println(j)
	}

	for { //without a condition it works like a while(true)
		fmt.Println("loop")
		break //get out of the loop
	}

	for n := 0; n <= 5; n++ { //this case will print only the odd numbers
		if n%2 == 0 {
			continue //pass/jump to the next iteration
		}
		fmt.Println(n)
	}

}
