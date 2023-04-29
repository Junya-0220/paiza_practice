package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < b[i]; j++ {
			if j == b[i]-1 {
				fmt.Print(a[index])
				index++
				continue
			}
			fmt.Print(a[index], " ")
			index++
		}
		fmt.Println()
	}
}
