package graph

// FindOrder returns the ordering of courses you should take to finish all courses.
// Time Complexity: O(n + e)
// Space Complexity: O(n + e)
func FindOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int)
	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	res := []int{}
	visit, cycle := make(map[int]bool), make(map[int]bool)

	var dfs func(int) bool
	dfs = func(course int) bool {
		if cycle[course] {
			return false
		}
		if visit[course] {
			return true
		}

		cycle[course] = true
		for _, pre := range preMap[course] {
			if !dfs(pre) {
				return false
			}
		}
		cycle[course] = false
		visit[course] = true
		res = append(res, course)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}
	return res
}
