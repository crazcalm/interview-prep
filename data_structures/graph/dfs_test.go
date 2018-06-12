package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	graph := ReadIn(testData)

	InitializeSearchDFS(graph)
	DFS(graph, 0)
	FindPath(0, 4, Parent)
	FindPath2(0, 4, Parent)
}
