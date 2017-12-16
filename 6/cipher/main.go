package main

import (
	"fmt"
	"regexp"
)

func main() {
	var N, K int
	var S string
	fmt.Scanf("%d\n", &N)
	fmt.Scanf("%s\n", &S)
	fmt.Scanf("%d\n", &K)

	k := K%26
	encrypted := ""
	for _, ch := range S {
		encrypted += string(ciper(ch, k))
	}

	fmt.Println(encrypted)
}

func ciper(ch rune, k int) int {
	encryptedCh := int(ch) + k
	if ok, _ := regexp.MatchString(`[a-z]`, string(ch)); ok {
		if encryptedCh > 'z' {
			encryptedCh -= 26
		}
		return encryptedCh
	} else if ok, _ := regexp.MatchString(`[A-Z]`, string(ch)); ok {
		if encryptedCh > 'Z' {
			encryptedCh -= 26
		}
		return encryptedCh
	} else {
		return ch
	}
}