package array_string

import (
	"strings"
)

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

func fullJustify_20250509(words []string, maxWidth int) []string {

	temp := make([][]string, 1)
	temp[0] = make([]string, 0)
	curIdx := 0
	curLength := 0

	for cur := 0; cur < len(words); {
		word := words[cur]
		if curLength+len(word) > maxWidth {
			curLength = 0
			curIdx++
			temp = append(temp, make([]string, 0))
			continue
		}

		// at least one space
		curLength += len(word) + 1
		temp[curIdx] = append(temp[curIdx], word)
		cur++
	}

	res := make([]string, 0)
	for i := range temp {

		wordLength := 0

		for j := range temp[i] {
			wordLength += len(temp[i][j])
		}

		space := maxWidth - wordLength
		var base, extra int
		if len(temp[i]) > 1 {
			base, extra = space/(len(temp[i])-1), space%(len(temp[i])-1)
		} else {
			base = space
		}

		ans := ""
		if i < len(temp)-1 {
			for j := range temp[i] {
				ans += temp[i][j]
				if len(ans) == maxWidth {
					break
				}
				ans += strings.Repeat(" ", base)
				if extra > 0 {
					ans += " "
					extra--
				}
			}
		} else {
			for j := range temp[i] {
				ans += temp[i][j]
				ans += " "
				if j == len(temp[i])-1 {
					ans += strings.Repeat(" ", maxWidth-len(ans))
				}
			}
		}

		res = append(res, ans)
	}
	return res
}
