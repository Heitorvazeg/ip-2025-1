package main

import "fmt"

func main() {
	var A, B, C, media float32
	var faltas, permitido int
	var aulas int = 64

	fmt.Scan(&A, &B, &C, &faltas)

	permitido = aulas / 4
	media = (A + B + C) / 3

	switch true {
	case faltas > permitido:
		fmt.Println("Reprovado por falta")
	case media >= 7:
		fmt.Println("Aprovado")
	case media < 7 && media >= 5:
		fmt.Println("Prova final")
	case media < 5 && media >= 0:
		fmt.Println("Reprovado")
	default:
		fmt.Println("Nota invalidade")
	}
}
