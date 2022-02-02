package main

import (
	"fmt"
)

func _array(nums ...int) {
	fmt.Println("Arrays: ", nums)
	result := sum_array(nums...)
	printResults(result)
}

func printResults(data int) {
	fmt.Println("Result: ", data)
	fmt.Println("")
}

func sum_array(nums ...int) int {
	fmt.Println("Nums º: ", nums)
	total := 0
	for _, num := range nums {
		fmt.Println("ite: ", num)
		total += num
	}
	return total
}

func fun_arrays() {

	fmt.Println("5_naked_fun")
	_array(3, 7)
	_array(3, 7, 1, 2, 1, 1, 2, 3)

}
