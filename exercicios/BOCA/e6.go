package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for j := 0; j < n; j++ {
		for k := 0; k < n; k++ {
			if k == n-1 {
				continue
			} else {
				if array[k] > array[k+1] {
					array[k], array[k+1] = array[k+1], array[k]
				}
			}
		}
	}
	for l := 0; l < len(array); l++ {
		fmt.Printf("%d ", array[l])
	}
}
