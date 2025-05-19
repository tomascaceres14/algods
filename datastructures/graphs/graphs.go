package graphs

import (
	"fmt"
	"slices"
)

type Graph[T comparable] struct {
	edges    map[T][]T
	directed bool
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		edges:    make(map[T][]T),
		directed: directed,
	}
}

func (g *Graph[T]) AddVertex(vertex T) {
	if _, exists := g.edges[vertex]; !exists {
		g.edges[vertex] = []T{}
	}
}

func (g *Graph[T]) RemoveVertex(vertex T) {
	delete(g.edges, vertex)

	for key := range g.edges {
		g.RemoveEdge(key, vertex)
	}
}

func (g *Graph[T]) AddEdge(from, to T) {
	g.AddVertex(from)
	g.AddVertex(to)

	g.edges[from] = append(g.edges[from], to)

	if !g.directed {
		g.edges[to] = append(g.edges[to], from)
	}
}

func (g *Graph[T]) RemoveEdge(from, to T) {

	if list, exists := g.edges[from]; exists {
		for i, v := range list {
			if v == to {
				g.edges[from] = slices.Delete(list, i, i+1)
				break
			}
		}
	}

	if !g.directed {
		if list, exists := g.edges[to]; exists {
			for i, v := range list {
				if v == from {
					g.edges[to] = slices.Delete(list, i, i+1)
					break
				}
			}
		}
	}
}

func (g *Graph[T]) String() string {
	result := ""
	for node, neighbours := range g.edges {
		result += fmt.Sprintf("%v: %v\n", node, neighbours)
	}
	return result
}
