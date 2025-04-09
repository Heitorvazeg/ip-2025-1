package main

import "fmt"

func main() {
	var dia, mes, ano int
	fmt.Scan(&dia, &mes, &ano)

	if dia <= 0 || dia > 31 {
		fmt.Println("Data invalida")
		return
	}

	switch true {
	case mes == 1:
		fmt.Printf("%d de Janeiro de %d", dia, ano)
	case mes == 2:
		fmt.Printf("%d de Fevereiro de %d", dia, ano)
	case mes == 3:
		fmt.Printf("%d de Março de %d", dia, ano)
	case mes == 4:
		fmt.Printf("%d de Abril de %d", dia, ano)
	case mes == 5:
		fmt.Printf("%d de Maio de %d", dia, ano)
	case mes == 6:
		fmt.Printf("%d de Junho de %d", dia, ano)
	case mes == 7:
		fmt.Printf("%d de Julho de %d", dia, ano)
	case mes == 8:
		fmt.Printf("%d de Agosto de %d", dia, ano)
	case mes == 9:
		fmt.Printf("%d de Setembro de %d", dia, ano)
	case mes == 10:
		fmt.Printf("%d de Outubro de %d", dia, ano)
	case mes == 11:
		fmt.Printf("%d de Novembro de %d", dia, ano)
	case mes == 12:
		fmt.Printf("%d de Dezembro de %d", dia, ano)
	default:
		fmt.Printf("Data invalida")
	}
}
