package trie

type WordDictionary struct {
	root *TrieNode
}

func ConstructorDict() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.children[c] == nil {
			curr.children[c] = &TrieNode{}
		}
		curr = curr.children[c]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(j int, root *TrieNode) bool {
		curr := root
		for i := j; i < len(word); i++ {
			c := word[i]
			if c == '.' {
				for _, child := range curr.children {
					if child != nil && dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if curr.children[c-'a'] == nil {
					return false
				}
				curr = curr.children[c-'a']
			}
		}
		return curr.isEnd
	}
	return dfs(0, this.root)
}
