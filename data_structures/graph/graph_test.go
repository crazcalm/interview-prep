package graph

import (
	"fmt"
	"testing"
)

var data = map[int][]int{
		0: []int{1,2,5,6},
		1: []int{0},
		2: []int{0},
		3: []int{4,5},
		4: []int{3,5,6},
		5: []int{0,3,4},
		6: []int{0,4},
		7: []int{8},
		8: []int{7},
		9: []int{10, 11, 12},
		10: []int{9},
		11: []int{9,12},
		12: []int{9,11},
	}

func TestAdj(t *testing.T) {
	graph := NewGraph()
	graph.ReadIn(data)

	tests := []struct{
		Node int
		Adj []int
	}{
		{12, []int{9, 11}},
		{13, []int{}},
	}

	for _, test := range tests {
		result := graph.Adj(test.Node)

		if len(result) != len(test.Adj) {
			t.Errorf("Expected %d elements, but got %d", len(test.Adj), len(result))
		}

		for _, node := range test.Adj {
			found := false
			for _, v := range result {
				if node == v {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected to find node %d, but did not see it", node)
			}
		}
		
	}
}

func TestReadIn(t *testing.T) {
	graph := NewGraph()
	graph.ReadIn(data)
	fmt.Println(graph)
}

func TestNumOfVertices(t *testing.T) {
	graph := NewGraph()
	graph.ReadIn(data)
	if len(data) != graph.NumOfVertices(){
		t.Errorf("Expected %d, but got %d", len(data), graph.NumOfVertices())
	}
}

func TestNumOfEdges(t *testing.T){
	graph := NewGraph()
	graph.ReadIn(data)
	answer := 13

	if answer != graph.NumOfEdges(){
		t.Errorf("Expected %d, but got %d", answer, graph.NumOfEdges())
	}
}