package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	cond1 := A + B
	cond2 := B + C
	cond3 := C + A

	if cond1 > C && cond2 > A && cond3 > B {
		if A != B && B != C {
			fmt.Println("Escaleno")

		} else if A == B && B == C {
			fmt.Println("Equilatero")

		} else {
			fmt.Println("Isosceles")
		}
	} else {
		fmt.Println("Nao forma triangulo")
	}
}
