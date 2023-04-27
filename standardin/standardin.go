package standardin
import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"strconv"
)

// 半角スペース区切りの 1,000 個の入力
// s_1 s_2 ... s_999 s_1000
func ReadStringSpace() {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    words := strings.Split(line, " ")
    for _, word := range words {
        fmt.Println(word)
    }
}

// 【整数の行入力】1,000行の整数の入力
/*
a_1
a_2
...
a_999
a_1000
*/
func ReadIntSpaceLine() {
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < 1000; i++ {
        scanner.Scan()
        line := scanner.Text()
        num, _ := strconv.Atoi(line)
        fmt.Println(num)
    }
}

// 1,000個の整数の半角スペース区切りの入力
// a_1 a_2 ... a_999 a_1000

func ReadIntSpace() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    nums := strings.Split(line, " ")
    for _, numStr := range nums {
        num, _ := strconv.Atoi(numStr)
        fmt.Println(num)
    }
}

//【N 個の整数の入力】1 行目で与えられる N 個の整数の入力 (large)
// N a_1 ... a_N
/*
期待する出力
a_1
...
a_N
*/
func ReadIntSpace2() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    nums := strings.Split(line, " ")
    for i := 1; i < len(nums); i++ {
        num, _ := strconv.Atoi(nums[i])
        fmt.Println(num)
    }
}
