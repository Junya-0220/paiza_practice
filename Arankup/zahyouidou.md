# 座標系での向きの変わる移動

## Step1 マップからの座標取得

マップの行数 H と列数 W とマップを表す H 行 W 列の文字列 S_1 ...S_H が与えられます。
要素が '#' になっているマスが 1 つあるので、その y , x 座標 を答えてください。

なお、マスの座標系は左上端のマスの座標を ( y , x ) = ( 0 , 0 ) とし、
下方向が y 座標の正の向き、右方向が x 座標の正の向きとします。



### 方針

各行各列の要素について for 文などを用いて全て調べます。
map_name[i][j] == '#' のときに i と j を出力すれば良いです。
二重ループを用いてすべてのマスを判定しています。
もし S[y][x] が # ならば座標を出力します。

```Go
func ZahyoStep1() {
	var H, W int
	fmt.Scan(&H, &W)

	S := make([][]rune, H)
	for y := 0; y < H; y++ {
		S[y] = make([]rune, W)
		var line string
		fmt.Scan(&line)
		for x, char := range line {
			S[y][x] = char
		}
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if S[y][x] == '#' {
				fmt.Println(y, x)
			}
		}
	}
}

```

## Step2 座標系での移動・方角

開始時点の y , x 座標 と移動の回数 N が与えられます。
続く N 行で移動の方角 d_1 ... d_N が与えられるので、与えられた順に移動をしたときの各移動後の y , x 座標 を答えてください。

ただし、図の通り、上側（ y 軸の負の向き）を北とします。

なお、マスの座標系は左上端のマスの座標を ( y , x ) = ( 0 , 0 ) とし、
下方向が y 座標の正の向き、右方向が x 座標の正の向きとします。

<img width="500" src="./image/zahyo2.png">


### 方針

各方角への移動は次の通り言い換えができます。（座標軸の向きに注意）
北(N)への移動 → y座標を -1
南(S)への移動 → y座標を +1
東(E)への移動 → x座標を +1
西(W)への移動 → x座標を -1
この操作を現在の座標について行えば、各移動後の座標が分かります。

```Go
func ZahyoStep2() {
	var y, x, n int
	fmt.Scan(&y, &x, &n)

	for i := 0; i < n; i++ {
		var direction string
		fmt.Scan(&direction)

		switch direction {
		case "N":
			y -= 1
		case "E":
			x += 1
		case "S":
			y += 1
		case "W":
			x -= 1
		}

		fmt.Println(y, x)
	}
}

```
## Step3 座標系での移動・向き

### 方針

```Go
```
## Step4 座標系での規則的な移動

### 方針

```Go
```
