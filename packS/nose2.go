package main

import (
	"fmt"
)

func mcd(a int, b int) int {
	for (a % b) != 0 {
		old_a := a
		old_b := b
		a = old_b
		b = old_a % old_b
	}
	return b
}

func main() {
	var a, b int
	fmt.Println("Ingresa 2 enteros: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El MCD es : ", mcd(a, b))
}
