package main

import "fmt"

func selectionShort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		i_min := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[i_min] > (*arr)[j] {
				i_min = j
			}
		}
		(*arr)[i_min], (*arr)[i] = (*arr)[i], (*arr)[i_min]
	}
}

func getMaximumAmount(quantity []int, m int) int {
	maximumAmount := 0
	n := len(quantity)

	for i := 0; i < m; i++ {
		selectionShort(&quantity)
		maximumAmount += quantity[n-1]
		quantity[n-1] -= 1
	}

	return maximumAmount
}

func main() {
	quantity := []int{10, 2, 3, 4, 5, 6, 7, 8, 9, 1}

	fmt.Println(getMaximumAmount(quantity, 2))
}
