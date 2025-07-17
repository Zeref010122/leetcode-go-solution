// Maximum Count of Positive Integer and Negative Integer
package binary_search_easy

func maximumCount(nums []int) int {
	negCount := binarySearch(nums, 0)
	posCount := len(nums) - binarySearch(nums, 1)

	if negCount > posCount {
		return negCount
	}
	return posCount

}

// create a binary search returning the index
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
