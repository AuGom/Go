//Crie um programa que utilize a declaração switch, onde o switch statement
//seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	esporteFavorito := "natacao"
	switch esporteFavorito {
	case "luta":
		fmt.Println("luta")
	case "futebol":
		fmt.Println("futebol")
	case "natacao":
		fmt.Println("natacao")
	}
}
