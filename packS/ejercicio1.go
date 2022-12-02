package main

import (
	"fmt"
	"math"
	"strconv"
)

func factorial(n int) uint64 {
	var rp uint64 = 1
	for i := 1; i <= n; i++ {
		rp *= uint64(i)
	}
	return rp
}

func myCos(x float64) float64 {
	acum := 0.0
	for i := 0; i < 30; i++ {
		idx := float64(i * 2)
		//fmt.Printf("idx: %g, signo: %g\n", idx, math.Pow(-1, float64(i) ))
		//fmt.Printf("Fact(i): %d\n", factorial(i*2))
		acum += math.Pow(-1, float64(i)) * math.Pow(x, idx) / float64(factorial(i*2))
	}
	return acum
}

func esPrimo(a int) bool {
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

// Codigo obtenido de los ejemplos de google
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func listaPrimos10k() []int {
	primos := []int{-1}
	for i := 2; i < 10000; i++ {
		if esPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func getReverse(n int) int {
	strn := strconv.Itoa(n)
	strnr := reverse(strn)
	revn, _ := strconv.Atoi(strnr)
	return revn
}

func mejorN() {
	lprimos := listaPrimos10k()

	for i := 1; i < len(lprimos); i++ {
		if lprimos[i] < 70 {
			continue
		}
		n := lprimos[i]
		revn := getReverse(n)
		revi := getReverse(i)
		if esPrimo(revn) {
			//fmt.Printf("Primos espejos: %d, %d\n", n, revn)
			if len(lprimos)-1 <= revi {
				//fmt.Printf("in: %d, revin: %d\n",in, revin)
				continue
			}
			if revn == lprimos[revi] {
				//Sus posiciones tambien son espejo
				fmt.Printf("El numero %d es mejor numero que 73\n", n)
				fmt.Printf("Posicion: %d\nPosicion espejo: %d\n", i, revi)
			}
		}

	}
}

func main() {
	fmt.Println("Ej1\n\n")
	mejorN()
	fmt.Println("\n\nEj2\n\n")
	var a float64
	fmt.Println("Ingresa: ")
	_, err := fmt.Scanf("%g", &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myCos(a))
	fmt.Println(myCos(a * math.Pi / 180.0))
}
