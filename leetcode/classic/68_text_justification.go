package classic

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	cur_length := 0
	cur_words := make([]string, 0)
	res := make([]string, 0)
	for _, word := range words {

		space := len(cur_words)

		if space+cur_length+len(word) > maxWidth {

			space_length := maxWidth - cur_length

			if len(cur_words) == 1 {
				res = append(res, cur_words[0]+strings.Repeat(" ", maxWidth-cur_length))
			} else {
				base, remain := space_length/(len(cur_words)-1), space_length%(len(cur_words)-1)

				temp := cur_words[0]
				cur := 1
				for i := 1; i <= remain; i++ {
					temp += strings.Repeat(" ", base+1)
					temp += cur_words[cur]
					cur++
				}

				for ; cur < len(cur_words); cur++ {
					temp += strings.Repeat(" ", base)
					temp += cur_words[cur]
				}
				res = append(res, temp)
			}

			cur_words = []string{word}
			cur_length = len(word)

		} else {
			cur_length = cur_length + len(word)
			cur_words = append(cur_words, word)
		}

	}

	// last line
	last_word := cur_words[0]
	for i := 1; i <= len(cur_words)-1; i++ {
		last_word += " "
		last_word += cur_words[i]
	}

	last_word += strings.Repeat(" ", maxWidth-len(last_word))
	res = append(res, last_word)
	return res
}
