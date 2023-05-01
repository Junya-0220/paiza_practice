package dataset

import "fmt"

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
