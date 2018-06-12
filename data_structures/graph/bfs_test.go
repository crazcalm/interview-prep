package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph := ReadIn(testData)

	InitializeSearch(graph)
	BFS(graph, 0)
	FindPath(0, 4, Parent)
	FindPath2(0, 4, Parent)

}

func TestConnectedComponents(t *testing.T) {
	graph := ReadIn(testData)

	InitializeSearch(graph)
	ConnectedComponents(graph)
}
