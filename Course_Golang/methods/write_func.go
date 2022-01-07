package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var first_var = 45
	var second_var = 4
	res := sums(first_var, second_var)
	fmt.Println("Declaring variables", rand.Intn(1))
	fmt.Println("Sums: ", res, sums(0, 0))
}

func sums(one int, second int) int {

	response := one + second

	if response == 0 {
		err := fmt.Errorf("one %q (second %d) not found", one, second)
		fmt.Println(err.Error())
		return 0
	}

	return response

}
