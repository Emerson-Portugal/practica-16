package main

import (
	"fmt"
	"strings"
)

func patron(cadena1 string, n int) (rpta string) {
	var aux []byte
	var auxiliar string
	for i := 0; i < n; i++ {
		aux = append(aux, cadena1[i])
		auxiliar = string(aux[:])
	}
	return auxiliar
}

func patronFinal(cadena1 string, cadena2 string) (cadena3 string) {
	if cadena1 == cadena2 {
		return cadena1
	} else {
		for i := len(cadena1); i > 0; i-- {
			auxiliar := patron(cadena1, i)
			contain := strings.Contains(cadena2, auxiliar)
			if contain == true {
				return auxiliar
			}
		}
		return "no hay coincidencias"
	}
}

func main() {
	fmt.Println(patronFinal("cgtaattgcgat", "cgtacagtagc"))
	fmt.Println(patronFinal("abababab", "xdxdxdxdxx"))
	fmt.Println(patronFinal("abcde", "bjervbjabc"))
}
