package hashtablepractice

/*
numbers := []int{11, 2, 5, 9, 10, 3}
target := 12
pair1, pair2 := hashtablepractice.GetPair(numbers, target)
fmt.Printf("(%v, %v)\n", pair1, pair2)
*/
func GetPair(numbers []int, target int) (int, int) {
	cache := make(map[int]bool)
	for _, num := range numbers {
		val := target - num
		if cache[val] {
			return val, num
		}
		cache[num] = true
	}
	return 0, 0
}

/*
numbers = []int{11, 2, 5, 9, 11, 3}
pair1, pair2 = hashtablepractice.GetPairHalfSum(numbers)
fmt.Printf("(%v, %v)\n", pair1, pair2)
*/
func GetPairHalfSum(numbers []int) (int, int) {
	sumNumbers := 0
	for _, num := range numbers {
		sumNumbers += num
	}
	halfSum, remainder := sumNumbers/2, sumNumbers%2
	if remainder != 0 {
		return 0, 0
	}

	cache := make(map[int]bool)
	for _, num := range numbers {
		cache[num] = true
		val := halfSum - num
		if cache[val] {
			return val, num
		}
	}
	return 0, 0
}
