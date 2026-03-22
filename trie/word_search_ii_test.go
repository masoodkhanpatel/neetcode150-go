package trie

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	expected := []string{"eat", "oath"}

	result := FindWords(board, words)
	sort.Strings(result)
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FindWords() = %v; want %v", result, expected)
	}
}
