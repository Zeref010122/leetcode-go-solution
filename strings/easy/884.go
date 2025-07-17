package string_easy

import "strings"

//Uncommon Words from Two Sentences
func uncommonFromSentences(s1 string, s2 string) []string {
	unCommonMap := make(map[string]int)
	CombinedSentences := strings.Fields(s1 + " " + s2)

	for _, v := range CombinedSentences {
		unCommonMap[v]++
	}
	unCommon := []string{}
	for key, val := range unCommonMap {
		if val == 1 {
			unCommon = append(unCommon, key)
		}
	}

	return unCommon

}
