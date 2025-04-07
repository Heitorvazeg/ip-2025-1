package main

import "fmt"

func main() {
	var inteiro int
	fmt.Scan(&inteiro)

	if inteiro > 0 {
		fmt.Println("Positivo")

	} else if inteiro < 0 {
		fmt.Println("Negativo")

	} else {
		fmt.Println("Nulo")
	}
}
