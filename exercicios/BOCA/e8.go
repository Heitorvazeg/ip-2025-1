package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	distintos := make(map[int]bool)

	for _, valor := range vetor {
		distintos[valor] = true
	}

	fmt.Println(len(distintos))
}
