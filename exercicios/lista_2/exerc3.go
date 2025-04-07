package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	soma := A + B

	if soma > 20 {
		soma += 8
		fmt.Printf("%d", soma)

	} else {
		soma -= 5
		fmt.Printf("%d", soma)
	}
}
