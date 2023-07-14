package searchpractice

type IndexNum int

func linearSearch(numbers []int, value int) IndexNum {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == value {
			return IndexNum(i)
		}
	}
	return -1
}

func BinarySearch(numbers []int, value int) IndexNum {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == value {
			return IndexNum(mid)
		} else if numbers[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
