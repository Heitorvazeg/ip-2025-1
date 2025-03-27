package main

import "fmt"

func main() {
	var raio, largura float32
	var areatotal, areabase, arearet float32
	var valor float32
	var pi float32
	pi = 3.14

	fmt.Println("Digite o raio: ")
	fmt.Scan(&raio)

	fmt.Println("Digite a largura: ")
	fmt.Scan(&largura)

	areabase = 2 * (pi * raio * raio)
	arearet = 2 * pi * raio * largura
	areatotal = areabase + arearet

	valor = areatotal * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f", valor)
}
