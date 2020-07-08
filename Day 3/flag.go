package main

//Flag package its used to support flags inside a CLI
//To run this code and understand its use, type "go run flag.go -name john"
import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	//Params on flag declaration
	//first: the flag word
	//second: a default value
	//third: a description
	nameFlag := flag.String("name", "[name]", "Put your name here")
	flag.Parse() //After all flags declared execute the command-line parsing

	//When a flag is created, it is stored in a pointer.So, you may treat like it one
	//using * to get to its value or & to its position
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Hello, " + *nameFlag + ", please enter some text: ")
	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}
