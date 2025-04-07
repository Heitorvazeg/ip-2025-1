package main

import "fmt"

func main() {
	var horas, minutos, segundos int

	fmt.Println("Digite os valores: ")
	fmt.Scan(&horas, &minutos, &segundos)

	segundos = segundos + (minutos * 60) + (horas * 60 * 60)

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d", segundos)
}
