package main

import "fmt"

func main() {
	var a, b, c, d int
	var determinante float32

	fmt.Println("Digite os valores: ")
	fmt.Scan(&a, &b, &c, &d)

	determinante = float32(a*d - b*c)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f", determinante)
}
