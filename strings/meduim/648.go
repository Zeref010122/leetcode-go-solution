package string_meduim

import "strings"

//Replace Words

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		shortest := word
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) && len(root) < len(shortest) {
				shortest = root
			}
		}
		words[i] = shortest
	}

	return strings.Join(words, " ")
}
