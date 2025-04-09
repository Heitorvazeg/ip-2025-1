package main

import (
	"fmt"
	"strings"
)

func main() {
	var numero string
	fmt.Scan(&numero)

	algarismo := strings.Split(strings.TrimSpace(numero), "")

	if len(algarismo) != 3 && !strings.Contains(numero, "-") {
		fmt.Println("Digito invalido")

	} else if len(algarismo) != 4 && strings.Contains(numero, "-") {
		fmt.Println("Digito Invalido")

	} else {
		if len(algarismo) == 3 {
			fmt.Printf("%s", algarismo[1])

		} else {
			fmt.Printf("%s", algarismo[2])
		}
	}
}
