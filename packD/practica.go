package main

import (
	"fmt"
	"math"
	"strings"
)

// Para el ejercicio 1
// Función que halla el factorial de un número
func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}

// Para el ejercicio 1
// Función que hallar el coseno de un ángulo
func cosenoT(angulo float64, n int) float64 {
	var cosT = 0.0
	var cont = 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			cosT += (math.Pow(angulo, float64(cont))) / float64(factorial(cont))
		} else {
			cosT -= (math.Pow(angulo, float64(cont))) / float64(factorial(cont))
		}
		cont += 2
	}
	return cosT
}

// Para el ejercicio 2
// Función que hallar si un número es primo
func esPrimo(x int) bool {
	var cont = 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			cont++
		}
	}
	if cont == 2 {
		return true
	} else {
		return false
	}
}

// Para el ejercicio 2
// Función que hallar la secuencia de números primos
func nump(n int) []int {
	var s []int
	for i := 1; i <= n; i++ {
		var flag bool = esPrimo(i)
		if flag == true {
			s = append(s, i)
		}
	}
	return s
}

// Para el ejercicio 2
// Función para invertir un número
func invertirNum(x int) int {
	inv := 0
	for x != 0 {
		inv = inv * 10
		inv = inv + x%10
		x /= 10
	}
	return inv
}

// Para el ejercicio 2
// Función para encontrar su un texto es palíndromo
func palindromo(text string) bool {
	inverso := ""

	for i := len(text) - 1; i >= 0; i-- {
		inverso += string(text[i])
	}

	for i := 0; i < len(text)-1; i++ {
		if string(text[i]) != string(inverso[i]) {
			return false
		}
	}
	return true
}

// Para el ejercicio 2
// Función para encontrar posición de un número en un array
func posArray(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

// Para el ejercicio 2
// Función para encontrar el máximo entre 2 números
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// Para el ejercicio 2
// Función para encontrar números espejo
func espejo() {
	for i := 1; i <= 10000; i++ {
		if i != invertirNum(i) {
			//verificamos su palíndromo
			binI := fmt.Sprintf("%b", i)
			flagPalindromo := palindromo(binI)
			if flagPalindromo == true {
				ip := invertirNum(i)
				if esPrimo(ip) == true && esPrimo(i) == true {
					numMax := max(ip, i)
					arrPrimos := nump(numMax)
					posF := posArray(arrPrimos, i) + 1
					posS := posArray(arrPrimos, ip) + 1
					posFI := invertirNum(posF)
					if posS == posFI {
						fmt.Print("Espejo: ", i, " ", ip, "\n")
					}
				}
			}
		}
	}
}
func findSubtring(a int, b int, text string) string {
	substring := ""
	for j := a; j <= b; j++ {
		substring += string(text[j])
	}
	return substring
}

func secuencias(text1 string, text2 string) string {
	result := ""
	substring := ""
	var i = 0
	var j = 0
	for i = 0; i < len(text1); i++ {
		for j = len(text1) - 1; j >= 0; j-- {
			substring = findSubtring(i, j, text1)
			if strings.Contains(text1, substring) && strings.Contains(text2, substring) && len(substring) > len(result) {
				result = substring
			} else {
				continue
			}
		}
	}
	return result
}

func main() {
	fmt.Print(" -------- Ejercicio 1 -------- \n")
	fmt.Print("Coseno de 30: ", cosenoT(30*math.Pi/180, 30), "\n")

	fmt.Print("\n -------- Ejercicio 2 -------- \n")
	espejo()

	fmt.Print("\n -------- Ejercicio 3 -------- \n")
	fmt.Print("subs: ", secuencias("ctgactga", "actgagc"))
}
