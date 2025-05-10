package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	prev := 1
	count := 1
	auxiliar := 0
	for j := 1; j < len(array); j++ {
		if array[j] > prev {
			count++
			prev = array[j]

		} else {
			auxiliar = count
			count = 1
			prev = array[j]
		}
	}
	if count > auxiliar {
		fmt.Println(count)
		return
	}
	fmt.Println(auxiliar)
}
