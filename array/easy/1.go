package array_easy

//Two Sum
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range nums {
		sum := target - v
		// 9 - 2  = 7
		value, exist := numMap[sum]
		if exist {
			return []int{i, value}
		} else {
			numMap[v] = i
		}
	}

	return []int{}
}
