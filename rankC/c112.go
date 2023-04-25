package rankC

/*
Go言語で下記のプログラムを実装してください。

入力される値
入力は以下のフォーマットで与えられます。

N
s_1 f_1 t_1
s_2 f_2 t_2
...
s_N f_N t_N
・1 行目にループ回数の N が与えられます。
・続く N 行のうちの i 行目 (1 ≦ i ≦ N) には、整数 s_i, f_i, t_i がこの順で半角スペース区切りで与えられ、
これは i 日目の予定が「出発地の現地時間 s_i 時に出発し、飛行機に f_i 時間乗ったのち、到着地の現地時間 t_i 時に到着する」というものであるということを表します。
・入力は合計で N + 1 行となり、N + 1 行目の末尾に改行が 1 つ入ります。

期待する出力
min
max
・1 行目に、N 日間で最も短い「1日の時間」を出力してください。
・2 行目に、N 日間で最も長い「1日の時間」を出力してください。
・末尾に改行を入れ、余計な文字、空行を含んではいけません。

(出発地の現地時間 s_i )+ (飛行機に f_i) + (24 - 到着地の現地時間 t_i ) = 「1日の時間」
N回分「1日の時間」を求めて、その求めた「1日の時間」の中から1番大きいものをmaxに入れる。1番小さいものをminに入れる。
最終的にminとmaxを出力する。


条件
すべてのテストケースにおいて、以下の条件をみたします。

・1 ≦ N ≦ 100
・1 ≦ s_i, f_i, t_i ≦ 23 (1 ≦ i ≦ N)
・ケース中に出てくる任意の 2 つのタイムゾーンのペアに対して、時差が 24 時間以上となることはない

入力例1
3
7 5 12
10 6 20
12 3 8
出力例1
20
31
*/

import (
	"fmt"
)

func RankC112() {
	// 旅程の日数を表す整数 N を宣言
	var n int
	fmt.Scan(&n)

	var result []int

	for i := 0; i < n; i++ {
		var f, fl, e int
		fmt.Scan(&f, &fl, &e)
		var r int
		r = f + fl + (24 - e)
		result = append(result, r)
	}
	min := result[0]
	max := result[0]
	for _, number := range result {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}
