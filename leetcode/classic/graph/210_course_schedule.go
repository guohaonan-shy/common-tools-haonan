package graph

func findOrder(numCourses int, prerequisites [][]int) []int {

	var (
		courseGraph = make(map[int][]int, 0)
		init        func()
		dfs         func(course int)
		reached     = make([]int, numCourses)
		valid       = true

		stack = make([]int, 0)
	)

	init = func() {
		for _, prerequisite := range prerequisites {
			pre, cur := prerequisite[1], prerequisite[0]

			if _, ok := courseGraph[pre]; ok {
				courseGraph[pre] = append(courseGraph[pre], cur)
			} else {
				courseGraph[pre] = []int{cur}
			}
		}
	}

	dfs = func(course int) {
		reached[course] = 1
		nexts := courseGraph[course]

		for _, next := range nexts {
			if reached[next] == 0 {
				dfs(next)
				if !valid {
					return
				}
			} else if reached[next] == 1 {
				valid = false
				return
			}
		}

		reached[course] = 2
		stack = append(stack, course)
	}

	init()

	for i := 0; i < numCourses && valid; i++ {
		if reached[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}

	for i := 0; i < len(stack)/2; i++ {
		stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
	}

	return stack
}
