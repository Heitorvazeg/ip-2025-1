package main

import "fmt"

func main() {
	var entrada int
	fmt.Scan(&entrada)

	if entrada > 20 && entrada < 90 {
		fmt.Println("Está incluso")

	} else {
		fmt.Println("Não está incluso")
	}
}
