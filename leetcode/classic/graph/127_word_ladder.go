package graph

type Transformation struct {
	TargetStr string
	Cnt       int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	dict := make(map[string][]string, 0)

	init := func() {
		for i, word := range wordList {

			for j := i + 1; j < len(wordList); j++ {

				if differ(word, wordList[j]) {

					if dict[word] == nil {
						dict[word] = []string{wordList[j]}
					} else {
						dict[word] = append(dict[word], wordList[j])
					}

					if dict[wordList[j]] == nil {
						dict[wordList[j]] = []string{word}
					} else {
						dict[wordList[j]] = append(dict[wordList[j]], word)
					}
				}
			}
		}
		return
	}
	init()
	reached := make(map[string]bool, 0)
	queue := []*Transformation{{
		TargetStr: beginWord,
		Cnt:       1,
	}}

	for len(queue) > 0 {

		str := queue[0]
		queue = queue[1:]

		for _, next := range dict[str.TargetStr] {
			if next == endWord {
				return str.Cnt + 1
			}

			if !reached[next] {
				reached[next] = true
				queue = append(queue, &Transformation{
					TargetStr: next,
					Cnt:       str.Cnt + 1,
				})
			}
		}
	}
	return 0
}

func differ(a, b string) bool {
	cnt := 0
	for i := range a {
		if a[i] != b[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}
