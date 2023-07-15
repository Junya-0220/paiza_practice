package main

import (
	"fmt"
	searchpractice "paiza_pracitce/algo/search"
)

func main() {
	nums := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
	fmt.Println("numsの長さ", len(nums)-1)
	target := 24
	s := searchpractice.BinarySearch(nums, target)
	fmt.Println("index番号は:", s)
	if s == -1 {
		fmt.Println("見つからなかった")
	} else {
		fmt.Println(nums[s])
	}
}
