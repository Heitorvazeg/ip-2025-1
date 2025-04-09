package main

import "fmt"

func main() {
	var destino, retorno int
	fmt.Scan(&destino, &retorno)

	if retorno == 1 {
		switch true {
		case destino == 1:
			fmt.Println("Valor: $900.00")
		case destino == 2:
			fmt.Println("Valor: $650.00")
		case destino == 3:
			fmt.Println("Valor: $600.00")
		case destino == 4:
			fmt.Println("Valor: $550.00")
		}

	} else {
		switch true {
		case destino == 1:
			fmt.Println("Valor: $500.00")
		case destino == 2:
			fmt.Println("Valor: $350.00")
		case destino == 3:
			fmt.Println("Valor: $350.00")
		case destino == 4:
			fmt.Println("Valor: $300.00")
		}
	}
}
