package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	expected := [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}
	result := PacificAtlantic(heights)

	sortInner := func(arr [][]int) {
		sort.Slice(arr, func(i, j int) bool {
			if arr[i][0] != arr[j][0] {
				return arr[i][0] < arr[j][0]
			}
			return arr[i][1] < arr[j][1]
		})
	}

	sortInner(result)
	sortInner(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PacificAtlantic() = %v; want %v", result, expected)
	}
}
