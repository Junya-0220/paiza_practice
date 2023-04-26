package rankB

import (
    "fmt"
    "sort"
)

func RankB111() {
    var n int
    fmt.Scan(&n)

    weights := make([]int, n)
    var sum int
    for i := 0; i < n; i++ {
        fmt.Scan(&weights[i])
        sum += weights[i]
    }

    sort.Sort(sort.Reverse(sort.IntSlice(weights)))

    var left, right int
    for i := 0; i < n; i++ {
        if left < right {
            left += weights[i]
        } else {
            right += weights[i]
        }
    }

    diff := abs(left - right)
    fmt.Println(diff)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
