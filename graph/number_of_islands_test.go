package graph

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	result := NumIslands(grid)
	if result != expected {
		t.Errorf("NumIslands() = %d; want %d", result, expected)
	}

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expected2 := 3
	result2 := NumIslands(grid2)
	if result2 != expected2 {
		t.Errorf("NumIslands() = %d; want %d", result2, expected2)
	}
}
