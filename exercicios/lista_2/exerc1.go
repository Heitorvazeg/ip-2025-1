package main

import "fmt"

func main() {
	var inteiro int
	fmt.Scan(&inteiro)

	if inteiro%2 == 0 {
		fmt.Println("Par")

	} else {
		fmt.Println("Impar")
	}
}
