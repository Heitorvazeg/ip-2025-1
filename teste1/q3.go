package main

import "fmt"

func main() {
	var soma int
	soma = 0

	for i := 1; i <= 100; i++ {
		soma += i
		fmt.Printf("%d ", i)
	}
	fmt.Printf("A soma é: %d", soma)
}