package array_meduim

import "sort"

//Top K Frequent Elements
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	type pair struct {
		num  int
		freq int
	}

	pairs := []pair{}
	for num, f := range freq {
		pairs = append(pairs, pair{num, f})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}
