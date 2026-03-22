package graph

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	Solve(board)
	if !reflect.DeepEqual(board, expected) {
		t.Errorf("Solve() = %v; want %v", board, expected)
	}
}
