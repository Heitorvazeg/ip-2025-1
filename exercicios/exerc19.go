package main

import "fmt"

func main() {
	var inteiro int

	fmt.Println("Digite o inteiro: ")
	fmt.Scan(&inteiro)

	var soma float64 = 0.0

	for i := 1; i <= inteiro; i++ {
		soma += 1.0 / float64(i)
	}
	fmt.Printf("%.6f", soma)
}
