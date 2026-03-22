package graph

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "Example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			name:      "Example 2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LadderLength(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("LadderLength() = %d; want %d", result, tt.expected)
			}
		})
	}
}
