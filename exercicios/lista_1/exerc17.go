package main

import "fmt"

func main() {
	var x, y, y2 int

	fmt.Println("Digite os valores: ")
	fmt.Scan(&x, &y)

	if x%2 == 0 {
		y2 = x + (y * 2)
		for i := x; i < y2; i++ {
			if i%2 == 0 {
				fmt.Printf("%d ", i)
			}
		}

	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}
