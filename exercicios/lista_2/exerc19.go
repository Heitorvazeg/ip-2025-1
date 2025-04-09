package main

import (
	"fmt"
	"math"
)

func main() {
	var figura string
	var raio, altura float64

	fmt.Scan(&figura)

	if figura != "cone" && figura != "cilindro" && figura != "esfera" {
		fmt.Println("Figura invalida")
		return
	}

	if figura == "cone" {
		fmt.Scan(&raio, &altura)

		volume := (3.14 * (raio * raio) * altura) / 3
		area := (3.14 * raio * math.Sqrt(math.Pow(raio, 2)+math.Pow(altura, 2)))
		fmt.Printf("Volume: %.2f Area %.2f", volume, area)

	} else if figura == "cilindro" {
		volume := 3.14 * raio * raio * altura
		area := 2.0 * 3.14 * raio * altura
		fmt.Printf("Volume: %.2f Area %.2f", volume, area)

	} else if figura == "esfera" {
		volume := (4.0 * 3.14 * math.Pow(raio, 3)) / 3
		area := 4.0 * 3.14 * raio * raio
		fmt.Printf("Volume: %.2f Area %.2f", volume, area)
	}
}
