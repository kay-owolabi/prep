package solutions

import (
	"leetcode/graphs"
	"leetcode/lib"
)

func CriticalConnections(n int, connections [][]int) [][]int {
	var result [][]int
	network := graphs.Graph{}
	network.BuildGraphUnDirected(n, connections)

	index := 1
	var stack []int
	onStack := make(graphs.Nodes)
	indexes := make([]int, len(network))
	lowLink := make([]int, len(network))
	var tarjans func(fromNode int, connections graphs.Nodes)

	tarjans = func(fromNode int, connections graphs.Nodes) {
		indexes[fromNode] = index
		lowLink[fromNode] = index
		index++
		stack = append([]int{fromNode}, stack...)
		onStack[fromNode] = true

		var toNode int
		for toNode, _ = range connections {
			delete(network[toNode], fromNode)
			onstack, ok := onStack[toNode]
			if !ok {
				tarjans(toNode, network[toNode])
				lowLink[fromNode] = lib.MinInt(lowLink[fromNode], lowLink[toNode])
			} else if onstack {
				lowLink[fromNode] = lib.MinInt(lowLink[fromNode], indexes[toNode])
			}

			if !onStack[toNode] {
				result = append(result, []int{fromNode, toNode})
			}
		}

		if lowLink[fromNode] == indexes[fromNode] {
			w := -1
			var ok bool
			for ok = true; ok; ok = w != fromNode {
				w, stack = stack[0], stack[1:]
				onStack[w] = false
			}
		}
	}

	for node, connections := range network {
		_, ok := onStack[node]
		if !ok {
			tarjans(node, connections)
		}
	}
	return result
}
