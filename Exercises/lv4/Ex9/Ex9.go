package main

import "fmt"

func main() {
	x := map[string][]string{
		"augusto":  []string{"luta", "e-sport"},
		"lourival": []string{"futebol", "luta"},
		"fatima":   []string{"natacao", "ginastica"},
	}
	x["doly"] = []string{"bola", "paninho"}
	for k, v := range x {
		fmt.Println(k)
		for _, i := range v {
			fmt.Println(i)
		}
	}
}
