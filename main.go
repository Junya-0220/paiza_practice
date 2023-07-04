package main

import (
	"fmt"
	arankup "paiza_pracitce/Arankup"
)

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)
	S := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	for i := 0; i < N; i++ {
		var y, x int
		fmt.Scan(&y, &x)
		S[y] = arankup.ReplaceCharAt(S[y], x, '#')
	}

	for i := 0; i < H; i++ {
		fmt.Println(S[i])
	}
}
