package graphs

import (
	"errors"
	"leetcode/lib"
)

type Nodes map[int]bool
type Graph map[int]Nodes

func (g *Graph) BuildGraphDirected(n int, connections [][]int) {
	res := g.initNodes(n)
	for _, connection := range connections {
		res[connection[0]][connection[1]] = true
	}
	*g = res
}

func (g *Graph) BuildGraphUnDirected(n int, connections [][]int) {
	res := g.initNodes(n)
	for _, connection := range connections {
		res[connection[0]][connection[1]] = true
		res[connection[1]][connection[0]] = true
	}
	*g = res
}

func (g *Graph) initNodes(n int) Graph {
	res := make(Graph, n)
	for i := 0; i < n; i++ {
		res[i] = make(Nodes)
	}
	return res
}

func (g Graph) BFS(queue []int) {
	seen := make(Nodes)
	for len(queue) > 0 {
		items := queue
		queue = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for node, _ := range g[item] {
					queue = append(queue, node)
				}
			}
		}
	}
}

func Find(slice []int, x int) int {
	if slice[x] != x {
		slice[x] = Find(slice, slice[x])
	}
	return slice[x]
}

func Union(slice []int, x, y int) {
	xRoot := Find(slice, x)
	yRoot := Find(slice, y)
	if xRoot == yRoot {
		return
	}
	slice[yRoot] = xRoot
}

func (g Graph) DFS() {
	seen := make(Nodes)
	var dfs func(nodes Nodes)
	dfs = func(nodes Nodes) {
		for nodeId, _ := range nodes {
			if !seen[nodeId] {
				seen[nodeId] = true
				dfs(g[nodeId])
			}
		}
	}

	for nodeID, connectedNodes := range g {
		if !seen[nodeID] {
			dfs(connectedNodes)
		}
	}
}

func (g Graph) Tarjans() [][]int {
	index := 0
	stack := []int{}
	onStack := make(Nodes)
	indexes := make([]int, len(g))
	lowlink := make([]int, len(g))
	var strongconnect func(n int, connections Nodes)

	result := [][]int{}

	strongconnect = func(fromNode int, connections Nodes) {
		indexes[fromNode] = index
		lowlink[fromNode] = index
		index++
		stack = append([]int{fromNode}, stack...)
		onStack[fromNode] = true

		for toNode, _ := range connections {
			delete(g[toNode], fromNode)
			onstack, ok := onStack[toNode]
			if !ok {
				strongconnect(toNode, g[toNode])
				lowlink[fromNode] = lib.MinInt(lowlink[fromNode], lowlink[toNode])
			} else if onstack {
				lowlink[fromNode] = lib.MinInt(lowlink[fromNode], indexes[toNode])
			}
		}

		if lowlink[fromNode] == indexes[fromNode] {
			connected := []int{}
			w := -1
			var ok bool
			for ok = true; ok; ok = w != fromNode {
				w, stack = stack[0], stack[1:]
				onStack[w] = false
				connected = append(connected, w)
			}
			result = append(result, connected)
		}
	}

	for node, connections := range g {
		_, ok := onStack[node]
		if !ok {
			strongconnect(node, connections)
		}
	}

	return result
}

func (g Graph) TopoSortKahn() ([]int, error) {
	numOfNodes := len(g)
	incomingEdges := make([]int, numOfNodes)
	for sourceNode, _ := range g {
		for destNode, _ := range g[sourceNode] {
			incomingEdges[destNode] += 1
		}
	}

	var queue []int
	for node, _ := range g {
		if incomingEdges[node] == 0 {
			queue = append(queue, node)
		}
	}

	var visitedCount int
	var topOrder []int
	for visitedCount = 0; len(queue) > 0; visitedCount++ {
		var node int
		node, queue = queue[0], queue[1:]
		topOrder = append([]int{node}, topOrder...)

		for destNode, _ := range g[node] {
			incomingEdges[destNode] -= 1
			if incomingEdges[destNode] == 0 {
				queue = append(queue, destNode)
			}
		}
	}

	if visitedCount != numOfNodes {
		return nil, errors.New("one or more cycles detected")
	}
	return topOrder, nil
}

func (g Graph) TopoSortDFS() ([]int, error) {
	var order []int
	seen := make(Nodes)
	var visitAll func(items Nodes) error

	visitAll = func(items Nodes) error {
		for item, _ := range items {
			if !seen[item] {
				_, ok := seen[item]
				if ok {
					order = nil
					return errors.New("not a directed acyclic graph")
				}
				seen[item] = false
				if err := visitAll(g[item]); err != nil {
					return err
				}
				seen[item] = true
				order = append(order, item)
			}
		}
		return nil
	}

	keys := make(Nodes)
	for node, _ := range g {
		keys[node] = true
	}

	err := visitAll(keys)
	return order, err
}
