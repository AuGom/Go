package main

//bufio implements buffered I/O.
//log is a logging package to format output.
//os provides a platform-independent interface to operating system functionality.
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//NewReader instantiates a buffered reader.
//ReadString reads until a delimiter, like press enter on your keyboard.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(text)
}
