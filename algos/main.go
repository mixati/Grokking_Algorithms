package main

import (
	"algos/modules/search"
	"fmt"
)

func main() {
	var arr []int = []int{1, 3, 5, 7, 12, 15, 24}
	var num int = 15

	ans := search.Binary_search(arr, num)
	fmt.Println(ans)
}
