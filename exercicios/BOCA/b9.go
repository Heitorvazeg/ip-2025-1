package main

import "fmt"

func main() {
	var peso int
	var altura float32

	fmt.Scan(&peso, &altura)

	imc := float32(peso) / (altura * altura)

	switch true {
	case imc < 18.5:
		fmt.Println("Abaixo do peso")
	case imc >= 18.5 && imc < 25:
		fmt.Println("Peso normal")
	case imc >= 25 && imc <= 30:
		fmt.Println("Sobrepeso")
	case imc >= 30:
		fmt.Println("Obesidade")
	default:
		fmt.Println("Numeros invalidos")
	}
}
