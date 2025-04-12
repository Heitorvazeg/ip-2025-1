package main

import (
	"fmt"
	"math"
)

func main() {
	var b, e int
	var potencia int

	fmt.Scan(&b, &e)

	if e == 0 {
		potencia = 1

	} else {
		potencia = int(math.Pow(float64(b), float64(e)))
	}
	fmt.Print(potencia)
}
