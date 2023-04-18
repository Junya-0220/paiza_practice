package rankC

/*
給与計算のプログラムをgolangで書いてください。
時間帯ごとに
・9 時から 17 時まで : 時給 X 円 (通常の時給)
・17 時から 22 時まで : 時給 Y 円 (夜の時給)
・それ以外の時間 : 時給 Z 円 (深夜の時給)

入力は以下のフォーマットで与えられます。

```
X Y Z
N
S_1 T_1
S_2 T_2
...
S_N T_N
```

・1 行目には、通常の時給 X、夜の時給 Y、深夜の時給 Z がこの順に整数で半角スペース区切りで与えられます。
・2 行目には、出勤日数 N が整数で与えられます。
・続く N 行の i 番目 (1 ≦ i ≦ N) には、i 日目の出勤時刻 S_i と退勤時刻 T_i がこの順に整数で半角スペース区切りで与えられます。
・入力は合計で N + 2 行となり、入力値最終行の末尾に改行が 1 つ入ります。

期待する出力

・N 日間の給料の合計金額を整数で出力してください。
・末尾に改行を入れ、余計な文字、空行を含んではいけません。

条件
すべてのテストケースにおいて、以下の条件をみたします。

・0 ≦ X, Y, Z ≦ 3,000
・1 ≦ N ≦ 100
・0 ≦ S < T ≦ 23

入力例1
1000 1300 1500
4
0 9
9 17
17 22
22 23

出力例1
29500
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func _main() {
// 	var (
// 		x, y, z int //時給
// 		n       int //ループ回数
// 		sum     int //金額
// 	)
// 	fmt.Scan(&x, &y, &z)
// 	fmt.Scan(&n)

// 	for i := 0; i < n; i++ {
// 		a, b := Read()
// 		fmt.Println(a, "a")
// 		fmt.Println(b, "b")
// 	}
// 	var time int
// 	var a int
// 	var r int

// 	if time >= 9 {
// 		r = a * x
// 	} else if time >= 17 {
// 		r = a * y
// 	} else {
// 		r = a * z
// 	}
// 	fmt.Println(sum)
// }

func Read() (str1, str2 string) {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	return arr[0], arr[1]
}

func RankC076() {
	// 時給を格納する変数を宣言
	var x, y, z int
	// 出勤日数を格納する変数を宣言
	var n int
	// 時刻のペアを格納するスライスを宣言
	// [{0 9} {9 17} {17 22} {22 23}] こんな感じのデータになる
	var timePairs []struct{ start, end int }
	// 入力値を読み込む
	fmt.Scan(&x, &y, &z)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		timePairs = append(timePairs, struct{ start, end int }{s, t})
	}
	// 合計給与を計算するための変数を宣言
	var total int
	// 各日の給与を計算する
	for _, tp := range timePairs {
		// 1日の給与を計算するための変数を宣言
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
