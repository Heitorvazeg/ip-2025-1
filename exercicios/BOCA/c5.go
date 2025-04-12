package main

import "fmt"

func main() {
	var entrada int
	fmt.Scan(&entrada)

	for i := 0; i < entrada/2; i++ {
		segundo := i + 1
		terceiro := i + 2

		if i*segundo*terceiro == entrada {
			fmt.Printf("%d eh triangular", entrada)
			break
		}
		if i == (entrada/2)-1 {
			fmt.Printf("%d nao eh triangular", entrada)
		}
	}
}
