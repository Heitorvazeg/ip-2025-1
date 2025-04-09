package main

import (
	"fmt"
)

func main() {
	var X, fx int
	fmt.Scan(&X)

	if X != 2 {
		fx = 8 / (2 - X)
		fmt.Printf("f(x) = %d", fx)

	} else {
		fmt.Println("Erro de divisão por zero")
	}
}
