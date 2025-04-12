package main

import "fmt"

func main() {
	var N int
	var nota_a, nota_b float64
	var countA, countE, countR int
	fmt.Scan(&N)

	countR = 0
	countE = 0
	countA = 0

	for i := 1; i <= N; i++ {
		fmt.Scan(&nota_a, &nota_b)

		media := (nota_a + nota_b) / 2

		if media <= 3 {
			fmt.Printf("Aluno %d: Reprovado\n", i)
			countR += 1

		} else if media > 3 && media < 7 {
			fmt.Printf("Aluno %d: Exame\n", i)
			countE += 1

		} else if media >= 7 && media <= 10 {
			fmt.Printf("Aluno %d: Aprovado\n", i)
			countA += 1
		}
	}

	fmt.Printf("Total Aprovados: %d\n", countA)
	fmt.Printf("Total Exame: %d\n", countE)
	fmt.Printf("Total Reprovados: %d", countR)
}
