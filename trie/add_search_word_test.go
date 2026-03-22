package trie

import "testing"

func TestWordDictionary(t *testing.T) {
	dict := ConstructorDict()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	if dict.Search("pad") {
		t.Errorf("Search(\"pad\") should be false")
	}
	if !dict.Search("bad") {
		t.Errorf("Search(\"bad\") should be true")
	}
	if !dict.Search(".ad") {
		t.Errorf("Search(\".ad\") should be true")
	}
	if !dict.Search("b..") {
		t.Errorf("Search(\"b..\") should be true")
	}
}
