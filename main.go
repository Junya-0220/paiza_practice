package main

import "fmt"

func main() {
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
