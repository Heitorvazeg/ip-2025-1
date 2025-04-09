package main

import "fmt"

func main() {
	var (
		id int
		nota_a, nota_b, nota_c float64
		exercicios float64
		resultado string
	)
	fmt.Scan(&id, &nota_a, &nota_b, &nota_c, &exercicios)

	mediafinal := (nota_a + (nota_b*2) + (nota_c*3) + exercicios) / 7

	switch true {
	case mediafinal < 4.0:
		resultado = "E"
	case mediafinal >= 4.0 && mediafinal < 6.0:
		resultado = "D"
	case mediafinal >= 6.0 && mediafinal < 7.5:
		resultado = "C"
	case mediafinal >= 7.5 && mediafinal < 9.0:
		resultado = "B"
	case mediafinal >= 9.0 && mediafinal < 10.0:
		resultado = "A"
	default:
		fmt.Println("Algo deu errado")
	}

	fmt.Printf("%d %.2f %.2f %s", id, exercicios, mediafinal, resultado)
}