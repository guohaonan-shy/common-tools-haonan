package back_track

// 方法可以跑，但是当word_list变长以后调用栈变大，时间复杂度高
//func findLadders(beginWord string, endWord string, wordList []string) [][]string {
//	dict := make(map[string]int, 0)
//	for _, word := range wordList {
//		dict[word]++
//	}
//	res := make([][]string, 0)
//	temp := []string{}
//	minLength := len(wordList) + 1
//	var dfs func(cur string)
//	dfs = func(cur string) {
//		temp = append(temp, cur)
//		defer func() {
//			temp = temp[:len(temp)-1]
//		}()
//
//		if len(temp) > minLength {
//			return
//		}
//
//		if len(temp) > 0 && cur == endWord && len(temp) <= minLength {
//			if len(temp) < minLength {
//				minLength = len(temp)
//				res = make([][]string, 0)
//			}
//			ans := make([]string, len(temp))
//			copy(ans, temp)
//			res = append(res, ans)
//			return
//		}
//
//		for _, word := range wordList {
//			if !isDiff(cur, word) {
//				continue
//			}
//			if _, ok := dict[word]; !ok {
//				continue
//			}
//
//			dict[word]--
//			if dict[word] == 0 {
//				delete(dict, word)
//			}
//			dfs(word)
//			dict[word]++
//		}
//		return
//	}
//	dfs(beginWord)
//	return res
//}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]bool, 0)
	for _, word := range wordList {
		dict[word] = true
	}
	dict[beginWord] = true // begin word 不一定在dict
	distance := make(map[string]int, 0)
	adjacentMap := make(map[string][]string, 0)

	var expansion func(cur string) []string

	expansion = func(cur string) []string {
		res := make([]string, 0)
		bytes := []byte(cur)
		// 复杂度从len(n) * len(n) * num_of_word_list => len(n) * 26
		for i := range bytes {
			old := bytes[i]
			for str := 'a'; str <= 'z'; str++ {
				bytes[i] = byte(str)
				if dict[string(bytes)] && string(bytes) != cur {
					res = append(res, string(bytes))
				}
				bytes[i] = old
			}
		}
		return res
	}

	var bfs func(start string)
	bfs = func(start string) {
		queue := make([]string, 0)
		queue = append(queue, start)

		distance[start] = 0

		for len(queue) > 0 {
			length := len(queue)
			for i := 0; i < length; i++ {
				cur := queue[0]
				queue = queue[1:]
				words := expansion(cur)
				for _, word := range words {
					if _, ok := distance[word]; !ok {
						distance[word] = distance[cur] + 1
						queue = append(queue, word)
					}
				}

				if _, ok := adjacentMap[cur]; !ok {
					adjacentMap[cur] = words
				}
			}
		}
		return
	}

	bfs(endWord)

	res := make([][]string, 0)
	temp := make([]string, 0)
	minVal := distance[beginWord]
	reached := make(map[string]bool, 0)
	var dfs func(cur string)
	dfs = func(cur string) {

		if len(temp)-1 == minVal && cur == endWord {
			ans := make([]string, len(temp))
			copy(ans, temp)
			res = append(res, ans)
			return
		}

		for _, next := range adjacentMap[cur] {
			if distance[next]+1 != distance[cur] {
				continue
			}

			if reached[next] {
				continue
			}

			reached[next] = true
			temp = append(temp, next)
			dfs(next)
			temp = temp[:len(temp)-1]
			reached[next] = false
		}
		return
	}

	temp = append(temp, beginWord)
	reached[beginWord] = true
	dfs(beginWord)
	return res

}
