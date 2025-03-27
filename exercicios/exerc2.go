package main

import "fmt"

func main() {
	var total int32
	var popularp, geralp, arquip, cadeirap float32
	var popularqtd, geralqtd, arquiqtd, cadeiraqtd float32
	var populardin, geraldin, arquidin, cadeiradin float32
	var N int
	var totaldin float32

	fmt.Println("Fale a quantidade de testes: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Println("Digite os valores: ")
		fmt.Scan(&total, &popularp, &geralp, &arquip, &cadeirap)
		popularqtd = popularp * float32(total)
		geralqtd = geralp * float32(total)
		arquiqtd = arquip * float32(total)
		cadeiraqtd = cadeirap * float32(total)

		populardin = popularqtd
		geraldin = geralqtd * 5.00
		arquidin = arquiqtd * 10.00
		cadeiradin = cadeiraqtd * 20.00

		totaldin = (populardin + geraldin + arquidin + cadeiradin) / 100.00

		fmt.Printf("A RENDA DO JOGO N. %d E: %.2f\n", i, totaldin)
	}
}
