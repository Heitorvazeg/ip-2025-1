package main

import "fmt"

func main() {
	var salariob float32 = 500.00
	var comissao float32 = 9.0 / 100.0
	var venda int
	fmt.Scan(&venda)

	renda := salariob + (float32(venda) * comissao)

	if float32(venda) > 15000.0 {
		renda += 800.0
	}

	fmt.Printf("%f", renda)
}
