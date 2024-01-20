package trie_tree

type Trie struct {
	Nexts  [26]*Trie
	HasEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {

	if len(word) == 0 {
		return
	}

	cur := this
	for i := 0; i < len(word); i++ {
		if cur.Nexts[word[i]-'a'] == nil {
			cur.Nexts[word[i]-'a'] = &Trie{}
		}
		cur = cur.Nexts[word[i]-'a']

		if i == len(word)-1 {
			cur.HasEnd = true
		}
	}

}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := this

	for i := 0; i < len(word); i++ {
		if cur.Nexts[word[i]-'a'] == nil {
			return false
		} else {
			cur = cur.Nexts[word[i]-'a']
		}
	}

	return cur.HasEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	cur := this

	for i := 0; i < len(prefix); i++ {
		if cur.Nexts[prefix[i]-'a'] == nil {
			return false
		} else {
			cur = cur.Nexts[prefix[i]-'a']
		}
	}
	return true
}
