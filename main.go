package main

import (
	"fmt"
	arankup "paiza_pracitce/Arankup"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	arankup.JudgeYoko(H, W)
}
