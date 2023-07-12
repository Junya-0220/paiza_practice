package arankup

import "fmt"

func MapStep1() {
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

func MapStep2() {
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

func MapStep3() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([][]rune, h)
	for i := 0; i < h; i++ {
		s[i] = make([]rune, w)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			s[i][j] = char
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == 0 || s[y][x-1] == '#' {
				if x == w-1 || s[y][x+1] == '#' {
					fmt.Println(y, x)
				}
			}
		}
	}
}

func MapStep4() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([][]rune, h)
	for i := 0; i < h; i++ {
		s[i] = make([]rune, w)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			s[i][j] = char
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 || s[y-1][x] == '#' {
				if y == h-1 || s[y+1][x] == '#' {
					fmt.Println(y, x)
				}
			}
		}
	}
}

func MapStep5() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([][]rune, h)
	for i := 0; i < h; i++ {
		s[i] = make([]rune, w)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			s[i][j] = char
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			flagRow := false
			flagColumn := false

			if x == 0 || s[y][x-1] == '#' {
				if x == w-1 || s[y][x+1] == '#' {
					flagRow = true
				}
			}

			if y == 0 || s[y-1][x] == '#' {
				if y == h-1 || s[y+1][x] == '#' {
					flagColumn = true
				}
			}

			if flagColumn && flagRow {
				fmt.Println(y, x)
			}
		}
	}
}
