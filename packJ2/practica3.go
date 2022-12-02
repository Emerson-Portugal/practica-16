package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(bioinformatico("ctgactga", "actgagc"))
}

func bioinformatico(adn1, adn2 string) string {
	var resultado string
	for i := len([]rune(adn1)); i >= 0; i-- {
		for j := 0; j < i; j++ {
			if strings.Contains(adn2, adn1[j:i]) == true {
				if len([]rune(resultado)) < len([]rune(adn1[j:i])) {
					resultado = adn1[j:i]
				}
			}
		}
	}
	return resultado
}
