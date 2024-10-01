package main

import (
	// "algos/modules/search"
	"algos/modules/sort"
	"fmt"
)

func main() {

	// БИНАРНЫЙ ПОИСК
	// var arr []int = []int{1, 3, 5, 7, 12, 15, 24}
	// var num int = 15
	// ans := search.BinarySearch(arr, num)
	// fmt.Println(ans)

	// СОРТИРОВКА ВЫБОРОМ
	var arr []int = []int{1, -5, 15, 6, 4, 3, 0}
	ans := sort.SelectionSort(arr)
	fmt.Println(ans, "\n", arr)
}
