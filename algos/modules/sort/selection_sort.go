package sort

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_id := 0
	for i, value := range arr {
		if value < smallest {
			smallest = value
			smallest_id = i
		}
	}
	return smallest_id
}

/* Сортировка выбором */
func SelectionSort(arr []int) []int {
	newArr := make([]int, len(arr))
	arr_size := len(arr)
	for i := 0; i < arr_size; i++ {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr[smallest] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
	}
	return newArr
}
