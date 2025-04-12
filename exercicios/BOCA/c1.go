package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	soma := 0

	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Printf("%d", i)
			soma += i
			break
		}
		fmt.Printf("%d ", i)
		soma += i
	}

	fmt.Printf("\n%d", soma)
}
