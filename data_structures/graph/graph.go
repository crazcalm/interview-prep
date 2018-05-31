package graph

import (
	"fmt"
)

type Graph struct {
	Nodes	map[int][]int
}

//NewGraph -- returns a new instance of Graph
func NewGraph() Graph {
	nodes := make(map[int][]int)
	return Graph{Nodes: nodes}
}

//ReadIn -- Reads in data to create a graph
func (g Graph) ReadIn(data map[int][]int){
	for key, values := range data {
		_, ok := g.Nodes[key]
		if !ok {
			g.Nodes[key] = values
		}else {
			for _, value := range values {
				found := false
				for _, node := range g.Nodes[key] {
					if value == node {
						found = true
						break
					}
				}
				if !found {
					g.Nodes[key] = append(g.Nodes[key], value)
				}
			}
		}
	}
}

//NumOfVertices -- Returns the number of vertices in the graph
func (g Graph) NumOfVertices () int {
	return len(g.Nodes)
}

//NumOfEdges -- returns the number of edges in the graph
func (g Graph) NumOfEdges() int {
	var count int
	for _, values := range g.Nodes {
		count += len(values)
	} 
	return count /2
}

//AddEdge -- Adds an edge to the graph
func (g Graph) AddEdge(edge1, edge2 int){
	_, ok := g.Nodes[edge1]
	if !ok {
		g.Nodes[edge1] =  []int{edge2}
	}else {
		found := false
		for _, adj := range g.Nodes[edge1] {
			if adj == edge2 {
				found = true
				break
			}
		}
		if !found {
			g.Nodes[edge1] = append(g.Nodes[edge1], edge2)	
		}
		
	}

	//Do the other side
	g.AddEdge(edge2, edge1)
}

//Adj -- returns all edges connected to the passed in edge with distance one
func (g Graph) Adj (edge int) []int {
	adj, ok := g.Nodes[edge]
	if !ok {
		return []int{}
	}
	return adj
}

//String -- returns a string representation of the graph
func (g Graph) String() string {
	result := "Graph:\n"
	for key, value := range g.Nodes {
		result += fmt.Sprintf("node %d: value -- %v\n",key, value)
	}
	return result
}