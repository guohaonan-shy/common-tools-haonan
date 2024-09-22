package graph

import "strconv"

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := make(map[int][]int, 0)
	edges := len(prerequisites)
	preGraph := make(map[int]bool, 0)
	for _, p := range prerequisites {
		if p[0] == p[1] { // self cycle
			return false
		}

		if _, ok := courseGraph[p[1]]; ok {
			courseGraph[p[1]] = append(courseGraph[p[1]], p[0])
		} else {
			courseGraph[p[1]] = []int{p[0]}
		}

		preGraph[p[0]] = true
	}

	if len(courseGraph) == numCourses {
		return false
	}

	reached := make(map[string]bool, 0)
	var dfs func(course int)
	dfs = func(course int) {

		nexts, ok := courseGraph[course]
		if !ok {
			return
		}

		for _, next := range nexts {
			if !reached[strconv.Itoa(course)+"-"+strconv.Itoa(next)] {
				reached[strconv.Itoa(course)+"-"+strconv.Itoa(next)] = true
				dfs(next)
			}

		}
		return
	}
	for i := 0; i < numCourses; i++ {

		if preGraph[i] { // not start point of graph
			continue
		}

		dfs(i)
		if len(reached) == edges {
			return true
		}
	}

	return false

}

// 图遍历中，通常有三种状态，分别是没遍历，正在遍历(指子节点还没有遍历完成)，遍历完成(即该节点所有的子节点均以遍历完成)
func canFinish_Standard(numCourses int, prerequisites [][]int) bool {
	courseGraph := make(map[int][]int, 0)
	for _, p := range prerequisites {
		if p[0] == p[1] { // self cycle
			return false
		}

		if _, ok := courseGraph[p[1]]; ok {
			courseGraph[p[1]] = append(courseGraph[p[1]], p[0])
		} else {
			courseGraph[p[1]] = []int{p[0]}
		}
	}

	reached := make([]int, numCourses)
	valid := true
	var dfs func(course int)
	dfs = func(course int) {
		nexts := courseGraph[course]
		reached[course] = 1
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
		return
	}

	for i := 0; i < numCourses && valid; i++ {
		if reached[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// time complexity high
/*
	this type of practice start to apply deep first search for each point, and find if this graph exists cycle or not
	1. if there is a cycle when we are iterating, it means we cannot major all courses on cycle. We have to check whether this graph is a DAG
	eg. a => b => c => a;

	extra question: Can we find another path or solution to cover all courses?
	- no, in this case, if we iterate to a node that we have covered in current path, it means these courses are cycle dependencies.
*/
func canFinishV2(numCourses int, prerequisites [][]int) bool {
	// analysis the dependency between courses
	adjacentMap := make(map[int][]int, 0)
	for _, pair := range prerequisites {
		target, prerequisite := pair[0], pair[1]

		if _, ok := adjacentMap[prerequisite]; ok {
			adjacentMap[prerequisite] = append(adjacentMap[prerequisite], target)
		} else {
			adjacentMap[prerequisite] = []int{target}
		}
	}

	scheduled := make(map[int]bool, 0)
	var dfs func(cur int, remain int) bool
	dfs = func(cur int, remain int) bool {

		if scheduled[cur] {
			return false
		}

		adjacentList, _ := adjacentMap[cur]
		if len(adjacentList) == 0 && remain == 0 {
			return true
		}

		scheduled[cur] = true
		for _, next := range adjacentList {

			if !dfs(next, remain-1) {
				return false
			}
		}
		scheduled[cur] = false
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, numCourses-1) {
			return false
		}
	}
	return true
}
