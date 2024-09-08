package search

func BinarySearch(arr []int, num int) int {
	low := 0
	high := len(arr)
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == num {
			return mid
		}
		if guess > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
