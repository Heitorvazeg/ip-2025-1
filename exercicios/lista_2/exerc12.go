package main

import "fmt"

func main() {
	var idade int
	var class string
	fmt.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		class = "Recém-nascido"

	} else if idade >= 3 && idade <= 11 {
		class = "Criança"

	} else if idade >= 12 && idade <= 19 {
		class = "Adolescente"

	} else if idade >= 20 && idade <= 55 {
		class = "Adulto"

	} else if idade > 55 {
		class = "Idoso"

	} else {
		class = "Idade inválida"
	}

	fmt.Printf("%s", class)
}
