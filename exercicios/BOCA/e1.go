package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	arrayAlt := make([]float64, n)
	soma := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&arrayAlt[i])
		soma += arrayAlt[i]
	}

	media := float64(soma) / float64(n)

	for j := 0; j < len(arrayAlt); j++ {
		if arrayAlt[j] > media {
			fmt.Printf("%.2f\n", arrayAlt[j])
		}
	}
}
