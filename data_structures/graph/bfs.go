package graph

import (
	"fmt"
	"log"

	"github.com/crazcalm/interview-prep/data_structures/queue"
)

var (
	//Processed -- Bool slice to mark processed vertices
	Processed = make([]bool, MaxVertices)
	//Discovered -- Bool slice to mark discovered vertices
	Discovered = make([]bool, MaxVertices)
	//Parent -- int slice use to make and trace the parent child relationship
	Parent = make([]int, MaxVertices)
)

//InitializeSearch -- Resets the global variables used to mark graph traversal progress
func InitializeSearch(g *Graph2) {
	for i := 0; i <= g.NVertices; i++ {
		Processed[i] = false
		Discovered[i] = false
		Parent[i] = -1
	}
}

//BFS -- Breadth first search
func BFS(g *Graph2, start int) {
	queue := queue.NewLinkedListQueue()

	queue.Enqueue(start)
	Discovered[start] = true

	for queue.Len() != 0 {
		item, err := queue.Dequeue()
		if err != nil {
			log.Fatalf("Error occurred when dequeuing and item: %s", err.Error())
		}
		vertex := item.(int)
		processVertexEarly(vertex)
		Processed[vertex] = true

		child := g.Edges[vertex]
		for child != nil {
			childID := child.Y
			if Processed[childID] || g.Directed {
				processEdge(vertex, childID)
			}

			if !Discovered[childID] {
				queue.Enqueue(childID)
				Discovered[childID] = true
				Parent[childID] = vertex
			}

			child = child.Next
		}
		processVertexLate(vertex)
	}

}

func processVertexEarly(vertex int) {
	fmt.Printf("process vertex early: %d\n", vertex)
}

func processEdge(vertex, childID int) {
	fmt.Printf("process edge: parent(%d) Edge(%d)\n", vertex, childID)
}

func processVertexLate(vertex int) {
	fmt.Printf("process vertex late: %d\n\n", vertex)
}

//FindPath -- Prints out the path from node a to node b
func FindPath(start, end int, parent []int) {
	if start == end || end == -1 {
		fmt.Printf("%d:", start)
	} else {
		FindPath(start, parent[end], parent)
		fmt.Printf(" %d", end)
	}
}

//FindPath2 -- Returns a slice with the wanted path from node a to node b
func FindPath2(start, end int, parent []int) []int {
	result := []int{end}
	for {
		result = append(result, parent[end])
		end = parent[end]

		if end == start {
			break
		}
		if end == -1 {
			break
		}
	}

	//reverse slice
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

/*ConnectedComponents --
A connected component of an undirected graph is a maximal set of vertices such that
there is a path between every pair of vertices.
*/
func ConnectedComponents(g *Graph2) {
	var counter int //Counts the number of connected components
	for i := 0; i <= g.NVertices; i++ {
		if !Discovered[i] {
			counter++
			fmt.Printf("Component %d:\n", counter)
			BFS(g, i)
		}
	}
}
