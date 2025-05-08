package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(0)

	} else {
		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		vetorFim := make([]int, n)

		for j := 0; j < n-1; j++ {
			if j == 0 {
				vetorFim[j] = vetor[j+1]
				continue
			}
			vetorFim[j] = vetor[j-1] + vetor[j+1]
		}
		vetorFim[n-1] = vetor[n-2]

		for l := 0; l < n; l++ {
			fmt.Printf("%d ", vetorFim[l])
		}
	}
}
