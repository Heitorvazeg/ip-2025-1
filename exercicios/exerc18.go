package main

import "fmt"

func main() {
	var inicial, razao, N int

	fmt.Println("Digite os valores:")
	fmt.Scan(&inicial, &razao, &N)

	aN := inicial + (N-1)*razao
	resultado := N * (inicial + aN) / 2

	fmt.Printf("%d", resultado)
}
