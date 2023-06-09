package rankC

/*
あなたは生徒にしりとりの課題を出しました。生徒は英字小文字または英字大文字からなる N 個の単語を順に提出します。N 個の単語が与えられた順でしりとりが成立するかどうか判定してください。

ここで、N 個の単語が与えられた順でしりとりが成立するとは、全ての i (1 ≦ i ≦ N-1) において、i 番目の単語の末尾と i+1 番目の単語の先頭が同じ文字となっていることを言います。このとき、小文字と大文字を区別することに注意してください。

しりとりが成立しない場合、初めてしりとりが成立しなかった 2 つの連続する単語の 1 つめの単語の末尾の文字と 2 つめの単語の先頭の文字を出力してください。しりとりが成立する場合は、Yes と出力してください。

入力例 1 の場合、はじめの二つの単語、apPle と error は、apPle の末尾と error の先頭が同じ e であるため、しりとりが成立しています。しかし、続く 2 つの単語 error と suBway は、error の末尾の文字と suBway の先頭の文字が等しくないので、しりとりは成立していません。したがって、期待される出力は r と s になります。

入力される値
入力は以下のフォーマットで与えられます。

N
s_1
s_2
...
s_N
・1 行目に入力される単語の数 N が与えられます。
・続く N 行のうちの i 行目 (1 ≦ i ≦ N) には英字小文字および、英字大文字からなる文字列 s_i が与えられます。
・入力は合計で N + 1 行となり、入力値最終行の末尾に改行が 1 つ入ります。

期待する出力
しりとりが成立する場合、以下の形式で出力してください。
Yes
しりとりが成立しない場合、以下の形式で出力してください。
x y
期待する出力は 1 行からなります。初めてしりとりが成立しなかった 2 つの連続する単語の 1 つ目の単語の末尾の文字を表す英字 x と 2 つ目の単語の先頭の文字を表す英字 y をこの順で半角スペース区切りで出力してください。
末尾に改行を入れ、余計な文字、空行を含んではいけません。

入力例1
4
apPle
error
suBway
Zb

出力例2
r s

入力例2
3
adijA
Akq
qZR
出力例3
Yes
*/

import "fmt"

func RankC144() {
	var N int
	fmt.Scan(&N)

	var prev string
	fmt.Scan(&prev)

	for i := 1; i < N; i++ {
		var current string
		fmt.Scan(&current)
		if prev[len(prev)-1] != current[0] {
			fmt.Printf("%c %c\n", prev[len(prev)-1], current[0])
			return
		}
		prev = current
	}

	fmt.Println("Yes")
}
