package sortpractice

import (
	"math/rand"
	"time"
)

func BubbleSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	lennums := len(nums)
	for i := 0; i < lennums; i++ {
		for j := 0; j < lennums-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func CocktailSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	lennums := len(nums)
	swapped := true
	start := 0
	end := lennums - 1
	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end = end - 1

		for i := end - 1; i >= start; i-- {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}

		start = start + 1
	}

	return nums
}

func CombSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}

	lennums := len(nums)
	gap := lennums
	swapped := true

	for gap != 1 || swapped {
		gap = int(float64(gap) / 1.3)
		if gap < 1 {
			gap = 1
		}

		swapped = false

		for i := 0; i < lennums-gap; i++ {
			if nums[i] > nums[i+gap] {
				nums[i], nums[i+gap] = nums[i+gap], nums[i]
				swapped = true
			}
		}
	}

	return nums
}

func SelectionSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}

		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

func GnomeSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	lennums := len(nums)
	index := 0
	for index < lennums {
		if index == 0 {
			index += 1
		}
		if nums[index] >= nums[index-1] {
			index += 1
		} else {
			nums[index], nums[index-1] = nums[index-1], nums[index]
			index -= 1
		}
	}

	return nums
}

func insertionSort(nums []int) []int {
	lennums := len(nums)
	for i := 1; i < lennums; i++ {
		temp := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j -= 1
		}

		nums[j+1] = temp
	}

	return nums
}

func BucketSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}

	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	lennums := len(nums)
	size := maxNum / lennums

	buckets := make([][]int, size)
	for i := 0; i < size; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, num := range nums {
		i := num / size
		if i != size {
			buckets[i] = append(buckets[i], num)
		} else {
			buckets[size-1] = append(buckets[size-1], num)
		}
	}

	for i := 0; i < size; i++ {
		insertionSort(buckets[i])
	}

	result := make([]int, 0)
	for i := 0; i < size; i++ {
		result = append(result, buckets[i]...)
	}

	return result
}

func ShellSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}

	lennums := len(nums)
	gap := lennums / 2
	for gap > 0 {
		for i := gap; i < lennums; i++ {
			temp := nums[i]
			j := i
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = temp
		}
		gap /= 2
	}

	return nums
}

func CountingSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	counts := make([]int, maxNum+1)
	result := make([]int, len(nums))

	for _, num := range nums {
		counts[num] += 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		index := nums[i]
		result[counts[index]-1] = nums[i]
		counts[index] -= 1
	}

	return result
}

func partition(nums []int, low int, high int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func QuickSort() []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(1001)
	}
	var _quickSort func(nums []int, low int, high int)
	_quickSort = func(nums []int, low int, high int) {
		if low < high {
			partitionIndex := partition(nums, low, high)
			_quickSort(nums, low, partitionIndex-1)
			_quickSort(nums, partitionIndex+1, high)
		}
	}

	_quickSort(nums, 0, len(nums)-1)
	return nums
}
