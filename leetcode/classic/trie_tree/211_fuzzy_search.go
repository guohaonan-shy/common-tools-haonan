package trie_tree

type WordDictionary struct {
	Nexts [26]*WordDictionary
	IsEnd bool
}

func WordConstructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	cur := this
	for i := 0; i < len(word); i++ {
		if cur.Nexts[word[i]-'a'] == nil {
			cur.Nexts[word[i]-'a'] = &WordDictionary{}
		}
		cur = cur.Nexts[word[i]-'a']
	}
	cur.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {

	if len(word) == 0 {
		return this.IsEnd
	}

	if word[0] != '.' {
		if this.Nexts[word[0]-'a'] == nil {
			return false
		}
		return this.Nexts[word[0]-'a'].Search(word[1:])
	}

	for _, dict := range this.Nexts {
		if dict == nil {
			continue
		}

		if dict.Search(word[1:]) {
			return true
		}
	}
	return false

}
