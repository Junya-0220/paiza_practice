package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	// スライスを作成
	var slice []int

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	for i := 0; i < q; i++ {
		var qr, val int
		fmt.Scan(&qr, &val)
		if qr == 0 {
			PushBackSlice(&slice, val)
		} else if qr == 1 {
			PopBackSlice(&slice)
		} else if qr == 2 {
			// スライスの中身を表示して確認
			fmt.Println(slice)
		}

	}

}

func PushBackSlice(s *[]int, n int) {
	*s = append(*s, n)

}

func PopBackSlice(s *[]int) {
	*s = (*s)[:len(*s)-1]
}
