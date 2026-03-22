# Neetcode 150 — Go Solutions

Solutions to all [Neetcode 150](https://neetcode.io/practice) problems implemented in Go, with tests for every problem.

## Structure

Each category is its own Go package with an implementation file and a corresponding test file.

| Category | Problems |
|---|---|
| arrays_hashing | 9 |
| two_pointers | 5 |
| sliding_window | 6 |
| stack | 7 |
| binary_search | 7 |
| linked_list | 13 |
| trees | 17 |
| trie | 3 |
| backtracking | 9 |
| heap | 7 |
| graph | 13 |
| intervals | 6 |
| greedy | 8 |
| advanced_graphs | 6 |
| bit_manipulation | 7 |
| math_geometry | 8 |
| dynamic_programming_1d | 12 |
| dynamic_programming_2d | 11 |
| **Total** | **154** |

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific category:
```bash
go test ./dynamic_programming_1d/...
```

Run a specific test:
```bash
go test ./arrays_hashing/... -run TestTwoSum
```

## Requirements

- Go 1.18+
