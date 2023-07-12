package arankup

import "fmt"

func Step1() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)

	S := make([][]rune, H)
	for i := 0; i < H; i++ {
		S[i] = make([]rune, W)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			S[i][j] = char
		}
	}

	for i := 0; i < N; i++ {
		var y, x int
		fmt.Scan(&y, &x)
		fmt.Println(string(S[y][x]))
	}
}

func Step2() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)

	s := make([][]rune, h)
	for i := 0; i < h; i++ {
		s[i] = make([]rune, w)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			s[i][j] = char
		}
	}

	for i := 0; i < n; i++ {
		var y, x int
		fmt.Scan(&y, &x)
		s[y][x] = '#'
	}

	for y := 0; y < h; y++ {
		fmt.Println(string(s[y]))
	}
}
