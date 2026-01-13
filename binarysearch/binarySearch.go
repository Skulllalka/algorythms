package binarysearch

func BinarySearch(array []int, number int) int {
	left, right := 0, len(array)-1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if array[middle] == number {
			return middle
		} else if array[middle] < number {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
