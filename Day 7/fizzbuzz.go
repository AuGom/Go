/*
FizzBuzz
In this problem, you should display a list from 1 to 100, one on each line,
with the following exceptions:
-Numbers divisible by 3 should appear as 'Fizz' instead of number;
-Numbers divisible by 5 should appear as 'Buzz' instead of number;
-Numbers divisible by 3 and 5 should appear as' FizzBuzz 'instead of number'.
*/

package main

import "fmt"

func main() {

	for i := 1; i < 101; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
