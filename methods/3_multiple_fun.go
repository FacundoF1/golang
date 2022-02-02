package main

import (
	"fmt"
)

func operatorSum(one int, second int) {
	var add, msg = return_multiple(one, second)
	fmt.Println("Sums: ", add, msg)
}

func return_multiple(one int, second int) (int, string) {

	if one+second == 0 {
		err := fmt.Errorf("one: %d, second: %d", one, second)
		fmt.Println(err.Error())
		return 0, err.Error()
	}

	var a = one + second
	const b = "Add exit!"

	return a, b

}
