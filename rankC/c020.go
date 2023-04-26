package rankC

import (
	"fmt"
	"math"
)

/*
以下をGoで実装してください。

ある生鮮食品を m[kg] 仕入れました。 m[kg] のうち p[%] を売ることができました。
売れ残った量のうち q[%] が売れました
さて、m[kg] 仕入れたこの食品は最終的に何kg 売れ残ったでしょうか。

入力される値
入力は以下のフォーマットで与えられます。

m p q
ここで、m, p, q はすべて整数です。

売れ残った量 [kg] を小数値で一行に出力してください。ただし、真値との誤差が0.0001 未満である場合にのみ、正答とみなされます。

最後は改行し、余計な文字、空行を含んではいけません。

条件
すべてのテストケースで以下の条件を満たします。

・1 ≦ m ≦ 1000
・0 ≦ p, q ≦ 100
・p = 100 のとき、q = 0

1 80 40
出力例1
0.12
入力例2
10 31 52
出力例2
3.312
*/

func RankC020test() {
	var m, p, q int
	fmt.Scan(&m, &p, &q)

	// 売れた量
	sold := float64(m) * float64(p) / 100.0

	// 売れ残った量
	remaining := float64(m) - sold
	remaining = remaining * float64(100-q) / 100.0
	result := math.Round(remaining*100) / 100
	// 結果を出力
	fmt.Printf("%.2f\n", result)
}

func RankC020A() {
	var m, p, q int
	fmt.Scan(&m, &p, &q)

	// 売れた量
	sold := float64(m) * float64(p) / 100.0

	// 売れ残った量
	remaining := float64(m) - sold
	remaining = remaining * float64(100-q) / 100.0
	result := math.Round(remaining*100) / 100
	// 結果を出力
	fmt.Printf("%.2f\n", result)
}
