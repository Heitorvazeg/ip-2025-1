package main

import "fmt"

func main() {
	var (
		array [10]int
	)
	pares := make([]int, 0)
	impares := make([]int, 0)

	for i := 0; i < 10; i++ {
		fmt.Scan(&array[i])
	}

	soma := 0
	qtd := 0
	for j := 0; j < 10; j++ {
		if array[j]%2 == 0 {
			pares = append(pares, array[j])
			soma += array[j]

		} else {
			impares = append(impares, array[j])
			qtd += 1
		}
	}
	fmt.Printf("%v\n", pares)
	fmt.Printf("%d\n", soma)
	fmt.Printf("%v\n", impares)
	fmt.Printf("%d\n", qtd)

}
