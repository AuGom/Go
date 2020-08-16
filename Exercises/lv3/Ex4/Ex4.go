//Crie um loop utilizando a sintaxe: for {}
//Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {
	x := 1995
	for {
		if x >= 2021 {
			break
		}
		fmt.Println(x)
		x++
	}
}
