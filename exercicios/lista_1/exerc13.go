package main

import "fmt"

func main() {
	var nota float32
	var conceito string

	fmt.Println("Digite a nota: ")
	fmt.Scan(&nota)

	if nota < 6.0 {
		conceito = "D"
	}
	if nota >= 6.0 && nota < 7.5 {
		conceito = "C"
	}
	if nota >= 7.5 && nota < 9.0 {
		conceito = "B"
	}
	if nota >= 9.0 && nota <= 10 {
		conceito = "A"
	}

	fmt.Printf("NOTA = %.2f CONCEITO = %s", nota, conceito)
}
