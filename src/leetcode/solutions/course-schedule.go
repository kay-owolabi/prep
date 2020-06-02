package solutions

import "leetcode/graphs"

func CanFinish(numCourses int, prerequisites [][]int) bool {
	pGraph := graphs.Graph{}
	pGraph.BuildGraphDirected(numCourses, prerequisites)
	_, err := pGraph.TopoSortKahn()
	return err != nil
}

func FindOrder(numCourses int, prerequisites [][]int) []int {
	pGraph := graphs.Graph{}

	pGraph.BuildGraphDirected(numCourses, prerequisites)
	res, _ := pGraph.TopoSortKahn()
	return res
}
