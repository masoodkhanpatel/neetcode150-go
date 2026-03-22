package trie

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.children[c] == nil {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if curr.children[c] == nil {
			return false
		}
		curr = curr.children[c]
	}
	return true
}
