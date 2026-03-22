package graph

// LadderLength returns the number of words in the shortest transformation sequence.
// Time Complexity: O(n * m^2) where n is word count and m is word length
// Space Complexity: O(n * m^2)
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	res := 1

	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return res
			}

			for j := 0; j < len(word); j++ {
				for c := 'a'; c <= 'z'; c++ {
					newWord := word[:j] + string(c) + word[j+1:]
					if wordSet[newWord] {
						queue = append(queue, newWord)
						delete(wordSet, newWord)
					}
				}
			}
		}
		res++
	}

	return 0
}
