package graphs

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	graph := NewGraph[int](false)

	graph.AddEdge(0, 1)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(4, 0)

	graph.AddVertex(5)

	graph.RemoveVertex(0)

	fmt.Println(graph)
}
