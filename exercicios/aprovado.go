package main

import "fmt"

func main() {
	var nota_a, nota_b, nota_c float32
	var media float32
	fmt.Println("Digite as notas: ")
	fmt.Scanln(&nota_a, &nota_b, &nota_c)
	media = (nota_a + nota_b + nota_c) / 3

	if media <= 6 {
		fmt.Printf("MEDIA = %.2f\n", media)
		fmt.Println("REPROVADO")
	} else {
		fmt.Printf("MEDIA = %.2f\n", media)
		fmt.Println("APROVADO")
	}
}
