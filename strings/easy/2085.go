package string_easy

//Count Common Words With One Occurrence
func countWords(words1 []string, words2 []string) int {
	freq := make(map[string]int)
	freq2 := make(map[string]int)
	count := 0
	for _, w := range words1 {
		freq[w]++
	}

	for _, w := range words2 {
		freq2[w]++
	}

	for key, val := range freq {
		if val == 1 && freq2[key] == 1 {
			count++
		}
	}

	return count
}
