package slide_window

func findSubstring(s string, words []string) []int {
	length, num_of_words, word_length := len(s), len(words), len(words[0])
	res := make([]int, 0)
	// 总共有word_length种分词方式，当i=word_length时候，本质上与i=0的分词方式是一样的，并且i之后有足够的字符计算
	for i := 0; i < word_length && i+num_of_words*word_length-1 < length; i++ {
		differ := make(map[string]int, 0)

		for j := 1; j <= num_of_words; j++ {
			word := s[i+(j-1)*word_length : i+j*word_length]
			differ[word]++
		}

		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		for start := i; start+num_of_words*word_length-1 < length; start += word_length {
			if start != i {
				right_word := s[start+(num_of_words-1)*word_length : start+num_of_words*word_length]
				differ[right_word]++
				if differ[right_word] == 0 {
					delete(differ, right_word)
				}

				left_word := s[start-word_length : start]
				differ[left_word]--
				if differ[left_word] == 0 {
					delete(differ, left_word)
				}
			}

			if len(differ) == 0 {
				res = append(res, start)
			}
		}

	}
	return res
}
