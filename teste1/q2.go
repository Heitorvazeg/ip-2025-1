package main

import "fmt"

func main() {
	var numero int
	fmt.Scan(&numero)

	if numero > 20 && numero < 90 {
		fmt.Println("Está compreendido")

	} else {
		fmt.Println("Não está compreendido")
	}
}