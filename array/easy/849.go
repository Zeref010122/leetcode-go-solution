package array_easy

//Intersection of Two Arrays
func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	intersection := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}

	for _, v := range nums2 {
		if set1[v] {
			intersection[v] = true
		}
	}
	output := []int{}
	for key := range intersection {
		output = append(output, key)
	}

	return output
}
