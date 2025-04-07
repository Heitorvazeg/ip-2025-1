package main

import "fmt"

func main() {
	var inteiro int
	fmt.Scan(&inteiro)

	if inteiro%5 == 0 {
		fmt.Printf("%d é divisivel por 5", inteiro)

	} else {
		fmt.Printf("%d não é divisivel por 5", inteiro)
	}
}
