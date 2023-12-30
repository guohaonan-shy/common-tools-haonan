package classic

import "strings"

func reverseWords(s string) string {
	words := make([]string, 0)

	var word = make([]rune, 0)
	for _, str := range s {
		if str == ' ' {
			if len(word) != 0 {
				words = append(words, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, str)
		}
	}

	if len(word) != 0 {
		words = append(words, string(word))
	}

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}
