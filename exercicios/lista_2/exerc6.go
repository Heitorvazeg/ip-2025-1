package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A%B == 0 {
		fmt.Printf("%d é divisivel por %d", A, B)

	} else {
		fmt.Printf("%d não é divisivel por %d", A, B)
	}
}
