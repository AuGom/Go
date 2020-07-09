package main

import "fmt"

func main() {

	fmt.Print("Enter a integer: ")
	var num int
	_, _ = fmt.Scanf("%d", &num)
	//the number typed above will be stored at "num", so i'm not using a variable
	//before the = sign
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	if num%5 == 0 { //if this statement is false, then nothing will be printed
		fmt.Println(num, " is divisible by 5")
	}

	//declare and make an if condition.
	//Variables stays available even when declared like this
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
