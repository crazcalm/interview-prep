package graph

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const (
	//MaxVertices -- Defines the Max number of node/vertices a graph can have
	MaxVertices = 1000
)

//Data -- Used to decode map input data from json
type Data struct {
	TotalNodes int              `json:"total_nodes"`
	Directed   bool             `json:"directed"`
	Mapping    map[string][]int `json:"mapping"`
}

//Node -- Node will be used as the linked list (like) structure to make the list portion of
//the adjacency lists
type Node struct {
	Y       int
	Weight  int
	Next    *Node
	Payload interface{}
}

//Graph2 -- This is the graph structure depicted in The Algorithm Design Manual
type Graph2 struct {
	Edges     []*Node
	NVertices int
	NEdges    int
	Directed  bool
	Degree    []int
}

//New2 -- Initializes a new Graph2
func New2(directed bool) *Graph2 {

	graph := Graph2{Edges: make([]*Node, MaxVertices), NVertices: 0, NEdges: 0, Directed: directed, Degree: make([]int, MaxVertices)}
	return &graph
}

//ReadIn -- Reads in data from json and adds it to the graph
func ReadIn(rawData []byte) *Graph2 {
	data := new(Data)
	err := json.Unmarshal(rawData, data)
	if err != nil {
		log.Fatal(err)
	}

	//Initialize the graph
	graph := New2(data.Directed)
	graph.NVertices = data.TotalNodes

	for node, edges := range data.Mapping {
		nodeInt, err := strconv.Atoi(node)
		if err != nil {
			log.Fatalf("Error occurred when node string (%s) to int: %s", node, err.Error())
		}
		for _, edge := range edges {
			InsertEdge(graph, nodeInt, edge, data.Directed)
		}
	}

	return graph
}

//InsertEdge -- Adds and edge to Graph2
func InsertEdge(g *Graph2, x, y int, directed bool) {
	//Creating the new edge
	edgeNode := new(Node)
	edgeNode.Weight = 0
	edgeNode.Y = y
	edgeNode.Next = g.Edges[x]

	//insert as the head of the linked list
	g.Edges[x] = edgeNode

	//Counts the number of edges of node x
	g.Degree[x]++

	if !directed {
		InsertEdge(g, y, x, true)
	} else {
		g.NEdges++
	}
}

//Print -- Prints out a Graph
func Print(g *Graph2) {
	fmt.Printf("Graph: \n")
	for i := 0; i <= g.NVertices; i++ {
		fmt.Printf("%d: ", i)
		node := g.Edges[i] //  Head of linked list of edge nodes

		for node != nil {
			fmt.Printf(" %d", node.Y)
			node = node.Next
		}
		fmt.Printf("\n")
	}
}
