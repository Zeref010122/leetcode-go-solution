package array_easy

//Intersection of Two Arrays II

func intersect(nums1 []int, nums2 []int) []int {
	frequencyMap := make(map[int]int)
	output := []int{}

	for _, v := range nums1 {
		frequencyMap[v]++
	}

	for _, v := range nums2 {
		val, exist := frequencyMap[v]
		if exist && val != 0 {
			frequencyMap[v]--
			output = append(output, v)

		}
	}
	return output
}
