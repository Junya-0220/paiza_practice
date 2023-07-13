package arankup

import "fmt"

func ZahyoStep1() {
	var H, W int
	fmt.Scan(&H, &W)

	S := make([][]rune, H)
	for y := 0; y < H; y++ {
		S[y] = make([]rune, W)
		var line string
		fmt.Scan(&line)
		for x, char := range line {
			S[y][x] = char
		}
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if S[y][x] == '#' {
				fmt.Println(y, x)
			}
		}
	}
}

func ZahyoStep2() {
	var y, x, n int
	fmt.Scan(&y, &x, &n)

	for i := 0; i < n; i++ {
		var direction string
		fmt.Scan(&direction)

		switch direction {
		case "N":
			y -= 1
		case "E":
			x += 1
		case "S":
			y += 1
		case "W":
			x -= 1
		}

		fmt.Println(y, x)
	}
}

func ZahyoStep3() {
	var y, x int
	var nowDirection string
	fmt.Scan(&y, &x, &nowDirection)

	var d string
	fmt.Scan(&d)

	lr := 1
	if d == "L" {
		lr = -1
	}

	if nowDirection == "N" {
		x += lr
	} else if nowDirection == "E" {
		y += lr
	} else if nowDirection == "S" {
		x -= lr
	} else if nowDirection == "W" {
		y -= lr
	}

	fmt.Println(y, x)
}

func ZahyoStep4() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	directions := []string{"E", "S", "W", "N"}
	nowDirection := 0
	count := 0
	maxCount := 1
	first := true

	for i := 0; i < n; i++ {
		x, y = move(directions[nowDirection], x, y)
		count++
		if first && count == maxCount {
			first = false
			count = 0
			nowDirection = (1 + nowDirection) % 4
		} else if count == maxCount {
			first = true
			count = 0
			maxCount++
			nowDirection = (1 + nowDirection) % 4
		}
	}

	fmt.Println(x, y)
}

func move(direction string, x, y int) (int, int) {
	if direction == "N" {
		y -= 1
	} else if direction == "E" {
		x += 1
	} else if direction == "S" {
		y += 1
	} else if direction == "W" {
		x -= 1
	}
	return x, y
}
