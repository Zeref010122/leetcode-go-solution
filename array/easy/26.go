package array_easy

//Remove Duplicates from Sorted Array

func removeDuplicates(nums []int) int {
	numMap := make(map[int]bool)
	var count int

	for _, v := range nums {
		if !numMap[v] {
			numMap[v] = true
			nums[count] = v
			count++
		}
	}

	return count
}
