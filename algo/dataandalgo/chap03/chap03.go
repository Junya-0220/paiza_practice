package chap03

import "fmt"

func Code03_1() {
	// 入力を受け取る
	var N, v int
	fmt.Scan(&N, &v)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	// 線形探索
	exist := false // 初期値は false に
	for i := 0; i < N; i++ {
		if a[i] == v {
			exist = true // 見つかったらフラグを立てる
		}
	}

	// 結果出力
	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
