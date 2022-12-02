package main

import (
	"fmt"
	"math"
)

func es_primo(a int) bool {
	i := 2
	var temp_raiz float64 = math.Sqrt(float64(a))
	var raiz_num int = int(temp_raiz)
	for i <= raiz_num {
		if a%i == 0 {
			return false
		}
		i++
	}
	return true
}

func main() {
	var a int
	fmt.Println("Ingresa: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Los numeros primos menores son: ")
	for i := 1; i < a; i++ {
		if es_primo(i) {
			fmt.Println(i)
		}
	}

}
