package rankC

import "fmt"

/*
go言語で下記のプログラムを作成してください。
入力は以下のフォーマットで与えられます。

N D
d_2
...
d_N
・1 行目には、整数 N および整数 D が半角スペース区切りで与えられます。
・1 + i 行目 (1 ≦ i ≦ N - 1) には整数 d_{i + 1} が与えられます。
これは、 i + 1 枚目の折り紙が i 枚目の折り紙に d_{i + 1} cm 重なっていることを表します。

NとDを1行で取得します。
N-1回文ループで回します。
ループの中で標準出力から受け取った数値を変数に入れます。


入力例1
3 4
2
1
出力例1
36

4 * 4 + (4-2) + (4-1)
4 * 9 = 36

入力例2
4 10
3
4
5
出力例2
280

10 * 10 + (10-3) + (10-4) + (10-5)
10 * 10 + 7 + 6 +5
10 * 28

*/

func RankC099() {
	var n, d int
	fmt.Scan(&n, &d)
	var loop []int

	for i := 0; i < n-1; i++ {
		var x int
		fmt.Scan(&x)
		loop = append(loop, x)

	}
	var calc int
	for _, n := range loop {
		r := d - n
		calc += r
	}

	fmt.Println(d * (calc + d))

}
