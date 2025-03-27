package main

import "fmt"

func main() {
	var A int

	fmt.Println("Digite o valor inteiro: ")
	fmt.Scan(&A)

	if A%3 == 0 && A%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")

	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
