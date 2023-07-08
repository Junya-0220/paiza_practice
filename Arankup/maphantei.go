package arankup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReplaceCharAt(str string, index int, newChar byte) string {
	strBytes := []byte(str)
	strBytes[index] = newChar
	return string(strBytes)
}

/*
マップの判定・横 (paizaランク C 相当)
入力される値
H W
S_0
...
S_(H-1)
*/
func JudgeYoko(H, W int) {

	S := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j == 0 || S[i][j-1] == '#' {
				if j == W-1 || S[i][j+1] == '#' {
					fmt.Println(i, j)
				}
			}
		}
	}
}

func NijigenHairetsu {
	scanner := bufio.NewScanner(os.Stdin)

	// H, W, N の入力を受け取る
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	H, _ := strconv.Atoi(values[0])
	_, _ = strconv.Atoi(values[1]) // Wは使用しないため、読み飛ばす
	N, _ := strconv.Atoi(values[2])

	// S を入力して二次元スライスに格納する
	S := make([][]byte, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		S[i] = []byte(scanner.Text())
	}

	// N 回のクエリを処理する
	for i := 0; i < N; i++ {
		// y, x の入力を受け取る
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		y, _ := strconv.Atoi(values[0])
		x, _ := strconv.Atoi(values[1])

		// S[y][x] の値を出力する
		fmt.Println(string(S[y][x]))
	}
}

func NijigenHairetsu2 {
	scanner := bufio.NewScanner(os.Stdin)

	// h, w の入力を受け取る
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	h, _ := strconv.Atoi(values[0])
	w, _ := strconv.Atoi(values[1])

	// s を入力して二次元スライスに格納する
	s := make([][]byte, h)
	for i := 0; i < h; i++ {
		scanner.Scan()
		s[i] = []byte(scanner.Text())
	}

	// 条件に合致する座標を出力する
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