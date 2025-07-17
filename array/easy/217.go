package array_easy

// Contains Duplicate

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)

	for _, v := range nums {
		if !numMap[v] {
			numMap[v] = true
		} else {
			return true
		}

	}

	return false
}
