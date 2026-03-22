package trie

// FindWords returns all words on the board.
// Time Complexity: O(m * n * 4^len(word))
// Space Complexity: O(total number of characters in words)
func FindWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		curr := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if curr.children[c] == nil {
				curr.children[c] = &TrieNode{}
			}
			curr = curr.children[c]
		}
		curr.isEnd = true
	}

	rows, cols := len(board), len(board[0])
	res := make(map[string]bool)
	visit := make(map[[2]int]bool)

	var dfs func(int, int, *TrieNode, string)
	dfs = func(r, c int, node *TrieNode, word string) {
		if r < 0 || c < 0 || r >= rows || c >= cols ||
			visit[[2]int{r, c}] || node.children[board[r][c]-'a'] == nil {
			return
		}

		visit[[2]int{r, c}] = true
		node = node.children[board[r][c]-'a']
		word += string(board[r][c])
		if node.isEnd {
			res[word] = true
		}

		dfs(r+1, c, node, word)
		dfs(r-1, c, node, word)
		dfs(r, c+1, node, word)
		dfs(r, c-1, node, word)

		visit[[2]int{r, c}] = false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root, "")
		}
	}

	var result []string
	for w := range res {
		result = append(result, w)
	}
	return result
}
