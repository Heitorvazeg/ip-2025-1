package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B, C int
	var concatenado string
	var elevado int
	var err error
	fmt.Println("Digite os inteiros: ")
	fmt.Scan(&A, &B, &C)

	if A > 9 || B > 9 || C > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		concatenado = strconv.Itoa(A) + strconv.Itoa(B) + strconv.Itoa(C)
		elevado, err = strconv.Atoi(concatenado)
		if err != nil {
			fmt.Println("Erro na conversão")
		}

		elevado = elevado * elevado

		fmt.Printf("%s, %d", concatenado, elevado)
	}

}
