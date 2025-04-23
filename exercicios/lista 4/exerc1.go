package main

import (
	"fmt"
)

func main() {
	var vetor [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	j := 0
	for i := 0; i < 10; i++ {
		if vetor[i] > 50 {
			fmt.Printf("%d %d\n", vetor[i], i)

		} else {
			j++
		}
	}
	if j == 10 {
		fmt.Println("Não há nenhum número maior que 50")
	}
}
