package dataset

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DispNarray() {
	// N,Mを読み込む
	var n, m int
	fmt.Scan(&n, &m)

	// Aを読み込む
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// AのM番目の値を出力する
	fmt.Println(a[m-1])
}

func DispNandQarray() {
	// Nを読み込む
	var n int
	fmt.Scan(&n)

	// Aを読み込む
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var q int
	fmt.Scan(&q)

	b := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&b[i])
	}
	for _, v := range b {
		fmt.Println(a[v-1])
	}
}

/*
N 個の要素からなる数列 A が与えられます。数列 A に対し、次の 3 つの操作を行うプログラムを作成してください。
・ push_back x : A の末尾に x を追加する
・ pop_back : A の末尾を削除する
・ print : A を半角スペース区切りで1行に出力する
1 行目に A の要素数 N と、A に対する操作回数 Q が与えられます。2 行目に A の 各要素の値が与えられます。
続くQ行のうち i 行目に、各操作の情報 query_i が与えられます。query_i は
0 x
または
1
または
2
の形式で与えられ、最初の数値が 0 x は push_back x を、 1 は pop_back を、 2 は print をそれぞれ表します。
入力例1
3 5
1 2 3
0 10
0 12
2
1
2
出力例1
1 2 3 10 12
1 2 3 10
*/
func DoutekiHairetsu() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// read N, Q
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	// read A
	A := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}

	// perform Q operations
	for i := 0; i < q; i++ {
		scanner.Scan()
		cmd, _ := strconv.Atoi(scanner.Text())

		if cmd == 0 {
			// push_back
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			A = append(A, x)
		} else if cmd == 1 {
			// pop_back
			A = A[:len(A)-1]
		} else {
			// print
			fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
		}
	}

}

/*
入力例1
5
1 2 3 3 6
出力例1
0 1 1 2 0 0 1 0 0 0
*/
func Shutsugenritsu() {
	scanner := bufio.NewScanner(os.Stdin)

	// N の入力
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// A の入力
	scanner.Scan()
	aStr := strings.Split(scanner.Text(), " ")
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(aStr[i])
	}

	// count の計算
	count := make([]int, 10)
	for _, num := range A {
		count[num]++
	}

	// 結果の出力
	fmt.Println(strings.Trim(fmt.Sprint(count), "[]"))
}

func ChofukuJudge(){
	nums := make(map[int]int)
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		if i > 0 {
			if nums[A] > 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		nums[A]++
	}
}
