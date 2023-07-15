package searchpractice

type IndexNum int

/*
nums := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
s := searchpractice.LinearSearch(nums, 15)
*/
func LinearSearch(numbers []int, value int) IndexNum {
	//見つかった時のIndex番号を返す
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == value {
			return IndexNum(i)
		}
	}
	//見つからない場合
	return -1
}

/*
nums := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
s := searchpractice.BinarySearch(nums, 15)
*/
func BinarySearch(numbers []int, value int) IndexNum {
	//leftは0,rightは配列の末尾を初期値にする
	left, right := 0, len(numbers)-1
	//leftがrightよりも大きくなるまでループをする
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == value {
			//半分の地点で探したい値が見つかったらそれを返す
			return IndexNum(mid)
		} else if numbers[mid] < value {
			/*
				探したい値が見つからなかった時、numbsers[mid]と探したい値を比較して、
				numbsers[mid]よりも探したい値の値が大きいなら、
				leftにmid + 1をする
			*/
			left = mid + 1
		} else {
			/*
				探したい値が見つからなかった時、numbsers[mid]と探したい値を比較して、
				numbsers[mid]よりも探したい値の値が小さいなら、
				rightにmid - 1をする
			*/
			right = mid - 1
		}
	}
	return -1
}
