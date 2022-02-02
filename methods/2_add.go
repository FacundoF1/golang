package main

import (
	"fmt"
)

func multipleOperator(one int, second int) {
	res := sums(one, second)
	fmt.Println("Sums: ", res)
}

func sums(one int, second int) int {

	response := one + second

	if response == 0 {
		err := fmt.Errorf("one: %d, second: %d", one, second)
		fmt.Println(err.Error())
		return 0
	}

	return response

}
