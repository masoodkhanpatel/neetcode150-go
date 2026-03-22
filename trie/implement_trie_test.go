package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		trie.Search("apple")
		t.Errorf("Search(\"apple\") should be true")
	}
	if trie.Search("app") {
		t.Errorf("Search(\"app\") should be false")
	}
	if !trie.StartsWith("app") {
		t.Errorf("StartsWith(\"app\") should be true")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Search(\"app\") should be true after Insert(\"app\")")
	}
}
