# マップの判定・縦横

## Step1 盤面情報の取得

```
行数 H , 列数 W の盤面があり、各マスには文字が 1 つだけ書かれています。盤面と N 個の y , x 座標 が与えられるので、与えられた座標の文字を順に出力してください。

なお、マスの座標系は左上端のマスの座標を ( y , x ) = ( 0 , 0 ) とし、
下方向が y 座標の正の向き、右方向が x 座標の正の向きとします
```

### 方針

まず入力を受け取り盤面を表す二重配列をつくります。
その後ループで n 個の座標を受け取り盤面の配列のインデックスを指定して出力します。


```Go
func Step1() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)

	S := make([][]rune, H)
	for i := 0; i < H; i++ {
		S[i] = make([]rune, W)
		var line string
		fmt.Scan(&line)
		for j, char := range line {
			S[i][j] = char
		}
	}

	for i := 0; i < N; i++ {
		var y, x int
		fmt.Scan(&y, &x)
		fmt.Println(string(S[y][x]))
	}
}
```


## Step2 盤面情報の変更

行数 H , 列数 W の盤面があり、各マスには文字が 1 つだけ書かれています。盤面と N 個 の y , x 座標 が与えられるので、盤面の与えられた座標の文字を "#" に書き換えた後の盤面を出力してください。

なお、マスの座標系は左上端のマスの座標を ( y , x ) = ( 0 , 0 ) とし、
下方向が y 座標の正の向き、右方向が x 座標の正の向きとします。

### 方針

マップは文字列の配列として受け取ると良いです。
マップの i 行 j 列の要素は、 S[i][j] で受け取れます。
S[i][j]='#' とすることで要素を書き換えられます。

```Go
```

## Step3 マップの判定・横

## Step4 マップの判定・縦

## Step5 マップの判定・縦横