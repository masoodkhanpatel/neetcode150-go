package graph

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	INF := 2147483647
	rooms := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	expected := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}
	WallsAndGates(rooms)
	if !reflect.DeepEqual(rooms, expected) {
		t.Errorf("WallsAndGates() = %v; want %v", rooms, expected)
	}
}
