package standardin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 半角スペース区切りの 1,000 個の入力
// s_1 s_2 ... s_999 s_1000
func ReadStringSpace() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")
	for _, word := range words {
		fmt.Println(word)
	}
}

// 【整数の行入力】1,000行の整数の入力
/*
a_1
a_2
...
a_999
a_1000
*/
func ReadIntSpaceLine() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 1000; i++ {
		scanner.Scan()
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		fmt.Println(num)
	}
}

// 1,000個の整数の半角スペース区切りの入力
// a_1 a_2 ... a_999 a_1000

func ReadIntSpace() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, " ")
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		fmt.Println(num)
	}
}

//【N 個の整数の入力】1 行目で与えられる N 個の整数の入力 (large)
// N a_1 ... a_N
/*
期待する出力
a_1
...
a_N
*/
func ReadIntSpace2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, " ")
	for i := 1; i < len(nums); i++ {
		num, _ := strconv.Atoi(nums[i])
		fmt.Println(num)
	}
}

// 改行区切りでの N 個の文字列の入力
/*
N
s_1
...
s_N

output
s_1
...
s_N
*/
func ReadStringFromN() {
	// 標準入力から値を受け取る
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
}

/*
1 行目で、整数 N と、続けて N 個の整数 a_1, ... , a_N が半角スペース区切りで与えられます。
a_1, ... , a_N を改行区切りで出力してください。
以下の形式で標準入力によって与えられます。
1 行目に整数 N と N 個の整数 a_1, ... , a_N が半角スペース区切りで与えられます。
N a_1 ... a_N
*/
func ReadNandNum() {
	// 標準入力から値を受け取る
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	nums := strings.Split(input, " ")

	// 先頭の値を取り出して、ループ回数として使用する
	n, _ := strconv.Atoi(nums[0])

	// ループ回数分だけ、元の文字列の配列の中身を出力する
	for i := 0; i <= n; i++ {
		if i == 0 {
			continue
		}
		fmt.Println(nums[i])
	}
}

/*
1 行目に整数 N が与えられます。
2 行目以降に、N 組の文字列 s_i と整数 a_i が N 行で与えられます。
8 組目の s_i と a_i を出力してください。
*/
func ReadAndOutByConditon() {
	// 標準入力から値を受け取る
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var s string
	var a int
	for i := 1; i <= n; i++ {
		scanner.Scan()
		input := scanner.Text()
		if i == 8 {
			// 8組目の場合は、sとaを取り出して出力する
			split := strings.Split(input, " ")
			s = split[0]
			a, _ = strconv.Atoi(split[1])
			break
		}
	}

	fmt.Printf("%s %d\n", s, a)
}

func ReadJissu() {
	var n int
	fmt.Scan(&n) // 実数の個数を入力

	nums := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i]) // 実数を入力
	}

	for _, num := range nums {
		fmt.Println(num) // 実数を出力
	}
}

func ReadNrowMcolumn() {

	var n, m int
	fmt.Scan(&n, &m) // 行列の行数と列数を入力

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j]) // 行列の各要素を入力
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(matrix[i][j])
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// N * N の九九表の出力
// 自然数 N が入力されます。N 行 N 列の九九表を出力してください。
func NkakeruN() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == n {
				fmt.Printf("%d", i*j)
				continue
			}
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}

//【行によって長さが違う二次元配列の表示】すべての行の長さと値が不定な 2 次元配列の出力

func TwoArray() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < b[i]; j++ {
			if j == b[i]-1 {
				fmt.Print(a[index])
				index++
				continue
			}
			fmt.Print(a[index], " ")
			index++
		}
		fmt.Println()
	}
}
