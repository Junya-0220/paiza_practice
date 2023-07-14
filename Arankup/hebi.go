package arankup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
入力例1
3 3 1 1 E
..#
..#
...

出力例1
No

入力例2
9 2 4 0 S
#.
#.
..
##
..
..
.#
..
.#

出力例2
Yes
*/
func HebiStep1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputs := strings.Split(input, " ")

	// 入力値の m 以外を整数に変換
	h, _ := strconv.Atoi(inputs[0])
	w, _ := strconv.Atoi(inputs[1])
	sy, _ := strconv.Atoi(inputs[2])
	sx, _ := strconv.Atoi(inputs[3])
	m := inputs[4]

	board := make([][]string, h)
	for i := 0; i < h; i++ {
		scanner.Scan()
		line := scanner.Text()
		board[i] = strings.Split(line, "")
	}

	move := map[string][2]int{
		"N": {-1, 0},
		"E": {0, 1},
		"S": {1, 0},
		"W": {0, -1},
	}

	ny := sy + move[m][0]
	nx := sx + move[m][1]

	if 0 <= ny && ny <= h-1 && 0 <= nx && nx <= w-1 && board[ny][nx] != "#" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
