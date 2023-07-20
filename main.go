package main

import (
	"fmt"
)

var sx, sy int

func move(D, M byte) {
	LR := 1

	if M == 'L' {
		LR = -1
	}

	switch D {
	case 'N':
		sx += LR
	case 'S':
		sx -= LR
	case 'E':
		sy += LR
	default:
		sy -= LR
	}
}

func main() {
	var H, W int
	var d, m byte
	fmt.Scan(&H, &W, &sy, &sx, &d, &m)

	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	move(d, m)

	if 0 <= sx && sx < W && 0 <= sy && sy < H && S[sy][sx] != '#' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
