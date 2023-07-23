package arankup

import (
	"fmt"
	"strconv"
)

func HebiStep1() {
	var h, w, sy, sx, m string
	fmt.Scan(&h, &w, &sy, &sx, &m)

	s := make([][]rune, 0)
	for i := 0; i < len(h); i++ {
		var line string
		fmt.Scan(&line)
		s = append(s, []rune(line))
	}

	syInt, _ := strconv.Atoi(sy)
	sxInt, _ := strconv.Atoi(sx)

	switch m {
	case "N":
		syInt -= 1
	case "E":
		sxInt += 1
	case "S":
		syInt += 1
	case "W":
		sxInt -= 1
	}

	if sxInt < 0 || sxInt >= len(w) || syInt < 0 || syInt >= len(h) || s[syInt][sxInt] == '#' {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func HebiStep3() {
	var h, w, sy, sx, n int
	fmt.Scan(&h, &w, &sy, &sx, &n)

	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	directions := [4]byte{'N', 'E', 'S', 'W'}
	now := 0

	for i := 0; i < n; i++ {
		var d string
		fmt.Scan(&d)

		if d == "L" {
			now = (3 + now) % 4
		} else {
			now = (1 + now) % 4
		}

		switch directions[now] {
		case 'N':
			sy--
		case 'E':
			sx++
		case 'S':
			sy++
		case 'W':
			sx--
		}

		if sx < 0 || sx >= w || sy < 0 || sy >= h || s[sy][sx] == '#' {
			fmt.Println("Stop")
			break
		} else {
			fmt.Println(sy, sx)
		}
	}
}
