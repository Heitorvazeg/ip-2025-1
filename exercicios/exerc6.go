package main

import "fmt"

func main() {
	var N int
	var fahre float32
	var celsius float32

	fmt.Println("Digite o numero de casos: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Println("Digite a temperatura em Fahrenheit: ")
		fmt.Scan(&fahre)

		celsius = 5 * (fahre - 32) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahre, celsius)
	}
}
