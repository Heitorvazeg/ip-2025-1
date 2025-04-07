package main

import "fmt"

func main() {
	var inteiro, quadrado int

	fmt.Println("Digite o valor inteiro: ")
	fmt.Scan(&inteiro)

	for i := 1; i <= inteiro; i++ {
		if i%2 == 0 {
			quadrado = i * i
			fmt.Printf("%d^%d = %d\n", i, i, quadrado)
		}
	}
}
