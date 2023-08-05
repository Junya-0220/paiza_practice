package chap02

import (
	"fmt"
	"math"
)

func Code02_1() {
	var N int
	fmt.Scan(&N)

	count := 0
	for i := 0; i < N; i++ {
		count++
	}
	fmt.Println(count)

}

func Code02_2() {
	var N int
	fmt.Scan(&N)

	count := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			count++
		}
	}
	fmt.Println(count)
}

func Code02_3() {
	var N int
	fmt.Scan(&N)

	for i := 2; i <= N; i += 2 {
		fmt.Println(i)
	}
}

// 2点 (x1, y1) と (x2, y2) との距離を求める関数
func calcDist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func Code02_4() {
	// 入力データを受け取る
	var N int
	fmt.Scan(&N)
	x := make([]float64, N)
	y := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	// 求める値を、十分大きな値で初期化しておく
	minimumDist := 100000000.0

	// 探索開始
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			// (x[i], y[i]) と (x[j], y[j]) との距離
			distIJ := calcDist(x[i], y[i], x[j], y[j])

			// 暫定最小値 minimumDist を distIJ と比べる
			if distIJ < minimumDist {
				minimumDist = distIJ
			}
		}
	}

	// 答えを出力する
	fmt.Println(minimumDist)
}
