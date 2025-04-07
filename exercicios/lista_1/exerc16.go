package main

import "fmt"

func main() {
	var salario, salarioreaj float32

	fmt.Println("Digite o salario: ")
	fmt.Scan(&salario)

	if salario <= 300.00 && salario > 0 {
		salarioreaj = salario + salario/2

	}
	if salario > 300.00 {
		salarioreaj = salario + salario*30/100
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f", salarioreaj)
}
