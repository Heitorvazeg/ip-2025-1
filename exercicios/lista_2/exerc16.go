package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	var delta float64

	if A == 0 {
		fmt.Print("Nao e equacao do segundo grau")

	} else {

		delta = float64(B*B - 4*A*C)

		if delta < 0 {
			fmt.Print("Raízes imaginárias")

		} else {
			raizdelta := math.Sqrt(delta)
			raiz_1 := (float64(-B) + raizdelta) / (2.0 * float64(A))
			raiz_2 := (float64(-B) - raizdelta) / (2.0 * float64(A))

			if raiz_1 != raiz_2 {
				fmt.Printf("%f %f Raízes distintas\n", raiz_1, raiz_2)

			} else {
				fmt.Printf("%f Raíz única\n", raiz_1)
			}
		}
	}
}
