package main

import "fmt"

func main() {
	var l, r int
	var soma, media, contador int
	fmt.Scan(&l, &r)

	contador = 0

	for i := l; i <= r; i++ {
		if i%2 == 0 {
			soma += i
			media += i
			contador++
		}
	}

	if contador > 0 {
		media = media / contador

		fmt.Printf("%d\n", soma)
		fmt.Printf("%d", media)

	} else {
		contador = 0
		fmt.Printf("%d\n", contador)
		fmt.Printf("%d", contador)

	}
}
