package main

import "fmt"

func main() {
	var idade int
	fmt.Scan(&idade)

	if idade < 0 {
		fmt.Println("Idade invalida")
		return
	}

	if idade < 16 {
		fmt.Println("Nao eleitor")

	} else if idade >= 16 && idade < 18 || idade > 65 {
		fmt.Println("Eleitor facultativo")

	} else {
		fmt.Println("Eleitor obrigatorio")
	}
}