package graph

// CanFinish returns true if you can finish all courses.
// Time Complexity: O(n + e) where n is nodes and e is edges
// Space Complexity: O(n + e)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	visit := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(course int) bool {
		if visit[course] {
			return false
		}
		if len(preMap[course]) == 0 {
			return true
		}

		visit[course] = true
		for _, pre := range preMap[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visit, course)
		preMap[course] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
