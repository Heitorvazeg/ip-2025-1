package main

import "fmt"

func main() {
	var fahre, polegadas, celsius, chuva float32

	fmt.Println("Digite os valores: ")
	fmt.Scan(&fahre, &polegadas)

	celsius = (5*fahre - 160) / 9

	chuva = polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", chuva)
}
