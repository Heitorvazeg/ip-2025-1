package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	vetorFinal := make([]int, n)
	soma := 0
	for j, _ := range vetorFinal {
		soma += vetor[j]
		vetorFinal[j] = soma
	}

	for l := 0; l < n; l++ {
		fmt.Printf("%d ", vetorFinal[l])
	}
}
