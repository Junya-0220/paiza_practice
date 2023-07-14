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
