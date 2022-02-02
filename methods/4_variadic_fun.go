package main

import (
	"fmt"
)

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

func multiple() (int, int) {
	return 3, 7
}

func fun_implement() {
	fmt.Println(split(17))
	fmt.Println("")
	var a, b = multiple()
	fmt.Println(a, b)
}
