package arankup

import "fmt"

func ReplaceCharAt(str string, index int, newChar byte) string {
	strBytes := []byte(str)
	strBytes[index] = newChar
	return string(strBytes)
}

/*
マップの判定・横 (paizaランク C 相当)
入力される値
H W
S_0
...
S_(H-1)
*/
func JudgeYoko(H, W int) {

	S := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j == 0 || S[i][j-1] == '#' {
				if j == W-1 || S[i][j+1] == '#' {
					fmt.Println(i, j)
				}
			}
		}
	}
}
