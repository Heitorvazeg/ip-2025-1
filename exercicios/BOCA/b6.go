package main

import "fmt"

func main() {
	var (
		conta     int
		limite    float32
		saldoi    float32
		depositos float32
		retiradas float32
	)

	fmt.Scan(&conta, &limite, &saldoi, &depositos, &retiradas)

	saldo := (saldoi + depositos) - retiradas

	if saldo > limite {
		fmt.Printf("Credito excedido: %f", saldo)

	} else {
		fmt.Printf("Saldo: %f", saldo)
	}
}
