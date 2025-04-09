package main

import "fmt"

func main() {
	var compra, venda float64
	fmt.Scan(&compra)

	if compra < 10.0 && compra >= 0.0 {
		venda = compra * (70.0 / 100.0)
	}

	if compra < 30.0 && compra >= 10.0 {
		venda = compra * (50.0 / 100.0)
	}

	if compra < 50.0 && compra >= 30.0 {
		venda = compra * (40.0 / 100.0)
	}

	if compra >= 50.0 {
		venda = compra * (30.0 / 100.0)
	}

	fmt.Printf("Lucro de $%.2f", venda)
}
