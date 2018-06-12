package graph

import (
	"fmt"
)

var (
	//Entry -- Used by DFS to collect the time when the vertex was first seen
	Entry = make([]int, MaxVertices)
	//Exit -- Used by DFS to collect the time when the vertex has been processed
	Exit = make([]int, MaxVertices)
	//Time -- Used by DFS to clock when vertexes enter and exit the DFS search
	Time = 0
	//Finished -- Used by DFS to note when the search is  done (Ex. Finding Cycles)
	Finished = false
)

//InitializeSearchDFS -- Resets the global variables used to mark graph traversal progress
func InitializeSearchDFS(g *Graph2) {
	for i := 0; i <= g.NVertices; i++ {
		Processed[i] = false
		Discovered[i] = false
		Parent[i] = -1
		Entry[i] = -1
		Exit[i] = -1
	}
	Time = 0
}

//DFS -- Depth-first search
func DFS(g *Graph2, vertex int) {
	if Finished { //If search is marked as done, return
		return
	}

	Discovered[vertex] = true //Seeing the vertex for the first time
	Time++
	Entry[vertex] = Time

	fmt.Printf("Entry Time for %d: %d\n", vertex, Time)

	processVertexEarly(vertex)

	child := g.Edges[vertex]
	for child != nil {
		childID := child.Y
		if !Discovered[childID] {
			Parent[childID] = vertex
			processEdge(vertex, childID)
			DFS(g, childID)
		} else if !Processed[childID] || g.Directed {
			processEdge(vertex, childID)
		}

		if Finished { //If search is marked as done, return
			return
		}

		child = child.Next
	}

	processVertexLate(vertex)
	Time++
	Exit[vertex] = Time

	fmt.Printf("Exit Time for %d: %d\n", vertex, Time)

	Processed[vertex] = true
}
