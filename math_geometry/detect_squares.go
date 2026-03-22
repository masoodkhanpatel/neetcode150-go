package math_geometry

// DetectSquares counts axis-aligned squares using added points.
type DetectSquares struct {
	pointCount map[[2]int]int
	points     [][2]int
}

func DetectSquaresConstructor() DetectSquares {
	return DetectSquares{pointCount: make(map[[2]int]int)}
}

// Add adds a point to the data structure.
func (ds *DetectSquares) Add(point []int) {
	p := [2]int{point[0], point[1]}
	if ds.pointCount[p] == 0 {
		ds.points = append(ds.points, p)
	}
	ds.pointCount[p]++
}

// Count returns the number of ways to choose 3 points such that they form an axis-aligned square with point.
// Time Complexity: O(n) per query
// Space Complexity: O(n)
func (ds *DetectSquares) Count(point []int) int {
	px, py := point[0], point[1]
	count := 0
	for _, p := range ds.points {
		// p is the diagonal point
		if p[0] == px || p[1] == py {
			continue
		}
		dx := p[0] - px
		dy := p[1] - py
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx != dy {
			continue
		}
		// Two other corners
		count += ds.pointCount[[2]int{px, p[1]}] * ds.pointCount[[2]int{p[0], py}]
	}
	return count
}
