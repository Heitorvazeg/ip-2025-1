package main

import "fmt"

func main() {
	var (
		precoI      float64
		qtdI        int
		despesasF   float64
		decrementoP float64
		incrementoV int
		precoM      float64
	)
	fmt.Scan(&precoI, &qtdI, &despesasF, &decrementoP, &incrementoV, &precoM)

	fmt.Println("Preco	Ingressos Vendidos	Lucro")

	lucroM := 0.0
	faixaP := 0.0
	ingressoF := 0

	for i := precoI; i >= precoM; i -= decrementoP {
		lucro := (i * float64(qtdI)) - despesasF

		fmt.Printf("%.2f		%d		%.2f\n", i, qtdI, lucro)
		qtdI += incrementoV

		if lucro > lucroM {
			lucroM = lucro
			faixaP = i
			ingressoF = qtdI - incrementoV
		}
	}
	fmt.Printf("Lucro maximo: %.2f\n", lucroM)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.", faixaP, ingressoF)
}
