package main

import (
	"fmt"
	"strings"
)

func analizarADN(a string, b string) string {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	base := ""
	mayorLen := 0
	for i := 0; i <= len(a); i++ {
		for j := 0; j < i; j++ {
			esSubstr := strings.Contains(b, a[j:i])

			if esSubstr && mayorLen < len(a[j:i]) {
				base = a[j:i]
				mayorLen = len(base)
			}
		}
	}

	return base
}

func main() {
	str1 := "ATGTCTTCCTCGA"
	str2 := "TGCTTCCTATGAC"
	fmt.Println(analizarADN(str1, str2))

	str1 = "ctgactga"
	str2 = "actgagc"
	fmt.Println(analizarADN(str1, str2)) //ok

	str1 = "cgtaattgcgat"
	str2 = "cgtacagtagc"
	fmt.Println(analizarADN(str1, str2)) //ok

	str1 = "ctgggccttgaggaaaactg"
	str2 = "gtaccagtactgatagt"
	fmt.Println(analizarADN(str1, str2))
}
