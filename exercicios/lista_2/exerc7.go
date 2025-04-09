package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if A < B {
		A, B = B, A
	}

	if B < C {
		B, C = C, B
	}

	if A < B {
		A, B = B, A
	}

	menor := C
	inter := B
	maior := A

	fmt.Printf("%d %d %d", menor, inter, maior)
}
