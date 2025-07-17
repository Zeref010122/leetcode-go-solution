package string_easy

// Valid Anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequency := make(map[rune]int)

	for _, ch := range s {
		frequency[ch]++
	}

	for _, ch := range t {
		frequency[ch]--
		if frequency[ch] < 0 {
			return false
		}
	}

	return true
}
