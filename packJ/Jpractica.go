package main

import (
	"fmt"
	"math"
	"strconv"
)

func fact(n float64) (x float64) {
	if n < 2 {
		return 1
	}
	return n * fact(n-1)
}

func coseno(x float64, n float64) (a float64) {
	if n < 0 {
		return 0
	} else {
		return coseno(x, n-1) + math.Pow(-1, n)*(math.Pow(x, 2*n)/fact(2*n))
	}
}

func primos(n int) (rpta bool) {
	contador := 0
	for x := 1; x <= n; x++ {
		if n%x == 0 {
			contador++
		}
	}
	if contador == 2 {
		return true
	} else {
		return false
	}
}

func espejo(x int) (rpta int) {
	rpta1 := strconv.FormatInt(int64(x), 10)
	var aux int
	var nums []byte
	for i := len(rpta1) - 1; i >= 0; i-- {
		z := rpta1[i]
		aux++
		nums = append(nums, z)
	}
	rptaFinal := string(nums[:])
	gg, err := strconv.Atoi(rptaFinal)
	if err == nil {
		//
	}
	return gg
}

func espejo2(x int) (rpta string) {
	rpta1 := strconv.FormatInt(int64(x), 10)
	var aux int
	var nums []byte
	for i := len(rpta1) - 1; i >= 0; i-- {
		z := rpta1[i]
		aux++
		nums = append(nums, z)
	}
	rptaFinal := string(nums[:])
	return rptaFinal
}

func mejorNumero(lim int) {
	contador := 0
	var y int
	var numeros [1000]int
	cont1 := 0
	for i := 10; i < lim; i++ {
		if primos(i) == true {
			numeros[cont1] = i
		}
	}
	for i := 0; i < lim; i++ {
		if primos(numeros[i]) == true {
			numeros[contador] = i
			contador++
			x := contador + 1
			numEspejo := espejo(i)
			for j := 0; j < len(numeros); j++ {
				if numeros[j] == numEspejo {
					y = j + 1
					break
				}
			}
			if espejo(x) == y {
				binario := strconv.FormatInt(int64(i), 2)
				gg, err := strconv.Atoi(binario)
				if err == nil {
					//
				}
				if gg == espejo(gg) {
					fmt.Println(i)
				}
			}
		}

	}
}

func main() {
	//Jose Grados
	//Ejercicio1
	v := strconv.FormatInt(85, 2) // change me!
	fmt.Println(v)
	fmt.Printf("v is of type %T\n", v)
	fmt.Print("Sen 90: ")
	fmt.Println((coseno(90*math.Pi/180, 10)))
	//Ejercicio2
	numero := int64(85)
	fmt.Println(strconv.FormatInt(numero, 2))
	fmt.Print(espejo(85))
	mejorNumero(100)
}
