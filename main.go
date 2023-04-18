package main

import "fmt"

func main() {
	// 時給を格納する変数を宣言
	var x, y, z int
	// 出勤日数を格納する変数を宣言
	var n int
	// 時刻のペアを格納するスライスを宣言
	var timePairs []struct{ start, end int }

	// 入力値を読み込む
	fmt.Scan(&x, &y, &z)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		timePairs = append(timePairs, struct{ start, end int }{s, t})
	}
	fmt.Println(timePairs)
	// 合計給与を計算するための変数を宣言
	var total int
	// 各日の給与を計算する
	for _, tp := range timePairs {
		// 1日の給与を計算するための変数を宣言
		fmt.Println(tp.start)
		fmt.Println(tp.end)
		var salary int
		for i := tp.start; i < tp.end; i++ {
			if i < 9 || i >= 22 {
				// 深夜
				salary += z
			} else if i < 17 {
				// 通常
				salary += x
			} else {
				// 夜
				salary += y
			}
		}
		total += salary
	}
	// 結果を出力
	fmt.Println(total)
}
