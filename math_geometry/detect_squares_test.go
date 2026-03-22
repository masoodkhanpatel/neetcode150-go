package math_geometry

import "testing"

func TestDetectSquares(t *testing.T) {
	ds := DetectSquaresConstructor()
	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})

	if got := ds.Count([]int{11, 10}); got != 1 {
		t.Errorf("Count([11,10]) = %d; want 1", got)
	}
	if got := ds.Count([]int{14, 8}); got != 0 {
		t.Errorf("Count([14,8]) = %d; want 0", got)
	}

	ds.Add([]int{11, 2})
	if got := ds.Count([]int{11, 10}); got != 2 {
		t.Errorf("Count([11,10]) after dup = %d; want 2", got)
	}
}
