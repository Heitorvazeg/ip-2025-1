package main

import "fmt"

func main() {
	var (
		vetorA = [10]int{4, 7, 5, 8, 2, 15, 9, 6, 10, 11}
		vetorB = [5]int{3, 4, 5, 8, 2}
	)

	vetorRA := make([]int, 0)
	vetorRB := make([]int, 0)

	for i := 0; i < 10; i++ {
		if vetorA[i]%2 == 0 {
			vetorRA = append(vetorRA, vetorA[i])

		} else {
			vetorRB = append(vetorRB, vetorA[i])
		}
	}

	for y := 0; y < len(vetorRA); y++ {
		for l := 0; l < 5; l++ {
			vetorRA[y] += vetorB[l]
		}
	}

	for z := 0; z < len(vetorRB); z++ {
		for l := 0; l < 5; l++ {
			vetorRB[z] += vetorB[l]
		}
	}

	fmt.Print(vetorRA)
	fmt.Print(vetorRB)
}
