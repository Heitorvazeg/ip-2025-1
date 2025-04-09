package main

import "fmt"

func main() {
	var valorfinal, valorfabrica float64
	var ar, pintura, vidro, hid bool
	fmt.Scan(&valorfabrica)

	valorfinal = valorfabrica

	ar = true

	if ar == true {
		valorfinal += 1750.0
	}

	if pintura == true {
		valorfinal += 800.0
	}

	if vidro == true {
		valorfinal += 1200.0
	}

	if hid == true {
		valorfinal += 2000.0
	}

	fmt.Printf("%.2f", valorfinal)
}
