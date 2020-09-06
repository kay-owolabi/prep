package main

import (
	"carl/imagecache"
	"math"
)

func main() {
	imagecache.Run()
}

type Coord struct {
	x, y int
}

type Coords []Coord

type Node struct {
	visited bool
	edges   map[Coord]*Node
}

var graph = map[Coord]*Node{}

func triggers(coords Coords, kIndex int, jIndex int) bool {
	for _, coord := range coords {
		graph[coord] = &Node{false, map[Coord]*Node{}}

	}

	for coord, node := range graph {
		for _, other := range coords {
			if coord == other {
				continue
			}
			if distance(coord, other) < 10 {
				node.edges[other] = graph[other]
			}
		}
	}

	return bfs(graph[coords[kIndex]], graph[coords[kIndex]])
}

func bfs(a, b *Node) bool {
	queue := []*Node{a}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		item.visited = true
		if item == b {
			return true
		}
		for _, edge := range item.edges {
			if !edge.visited {
				edge.visited = true
				queue = append(queue, edge)
			}
		}

	}
	return false
}

func distance(a, b Coord) float64 {
	x := a.x - b.x
	y := a.y - b.y
	return math.Sqrt(float64(x*x + y*y))
}
