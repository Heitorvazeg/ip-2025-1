package main

import "fmt"

func main() {
	var A, B, C int
	var delta float32

	fmt.Println("Digite A: ")
	fmt.Scan(&A)

	fmt.Println("Digite B: ")
	fmt.Scan(&B)

	fmt.Println("Digite C: ")
	fmt.Scan(&C)

	delta = float32(B*B - 4*A*C)

	fmt.Printf("O VALOR DE DELTA E = %.2f", delta)
}
