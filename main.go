package main

import "fmt"

func RankB091() {
	var N int
	fmt.Scan(&N)

	// 地図の読み込み
	heights := make([][]int, N+2)
	for i := 0; i < N+2; i++ {
		heights[i] = make([]int, N+2)
		for j := 0; j < N+2; j++ {
			if i == 0 || j == 0 || i == N+1 || j == N+1 {
				heights[i][j] = 0
			} else {
				fmt.Scan(&heights[i][j])
			}
		}
	}

	// 山頂を探す
	peaks := make([]int, 0)
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if heights[i][j] > heights[i-1][j] &&
				heights[i][j] > heights[i+1][j] &&
				heights[i][j] > heights[i][j-1] &&
				heights[i][j] > heights[i][j+1] {
				peaks = append(peaks, heights[i][j])
			}
		}
	}

	// 標高の高い順にソートして出力
	for i := 0; i < len(peaks); i++ {
		for j := i + 1; j < len(peaks); j++ {
			if peaks[i] < peaks[j] {
				peaks[i], peaks[j] = peaks[j], peaks[i]
			}
		}
	}
	for _, peak := range peaks {
		fmt.Println(peak)
	}
}
