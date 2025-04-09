package main

import "fmt"

func main() {
	var (
		dia               int
		preco, precofinal float64
		tipo              string
	)

	fmt.Scan(&dia, &preco)
	fmt.Println("Digite o tipo: ")
	fmt.Scan(&tipo)

	if tipo != "normal" && tipo != "lançamento" {
		fmt.Println("Tipo invalido")
		return
	}

	if dia < 1 || dia > 7 {
		fmt.Println("Dia invalido")
	}

	if preco < 0.0 {
		fmt.Println("Preço invalido")
	}

	if dia == 4 || dia == 6 || dia == 7 || dia == 1 {
		if tipo == "normal" {
			precofinal = preco

		} else if tipo == "lançamento" {
			precofinal = preco + (preco * 15.0 / 100.0)
		}

	} else if dia == 2 || dia == 3 || dia == 5 {
		if tipo == "normal" {
			precofinal = preco - (preco * 40.0 / 100.0)

		} else if tipo == "lançamento" {
			precofinal = preco - (preco * 40.0 / 100.0) + (preco * 15.0 / 100.0)
		}
	}

	fmt.Printf("O preço final é: %.2f", precofinal)
}
