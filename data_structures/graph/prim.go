package graph

var (
	//InTree -- used by Prim's Algorithm to check if the vertex is in the Prim tree
	InTree = make([]bool, MaxVertices)
	//Distance -- Used to mark the cost of going from one vertex to another
	Distance = make([]int, MaxVertices)
	//ArbitrarlyBigNumber -- A reallt big number
	ArbitrarlyBigNumber = 9999999999
)

//Prim -- Prim's Algorithm for minimum spanning tree
func Prim(g *Graph2, start int) {
	for i := 0; i <= g.NVertices; i++ {
		InTree[i] = false
		Distance[i] = ArbitrarlyBigNumber //Set the initial distance for each vertex really high so that
		//it does not get confused for the "actual" value before it gets updated
		Parent[i] = -1
	}

	Distance[start] = 0
	v := start

	for !InTree[v] {
		InTree[v] = true
		child := g.Edges[v]
		for child != nil {
			w := child.Y
			weight := child.Weight

			if Distance[w] > weight && !InTree[w] { // Updates the distance for child w
				Distance[w] = weight
				Parent[w] = v
			}
			child = child.Next
		}

		v = 1
		dist := ArbitrarlyBigNumber
		for i := 0; i <= g.NVertices; i++ {
			if !InTree[i] && dist > Distance[i] { // Selects the next vertex to be put in the InTree
				dist = Distance[i]
				v = i
			}
		}
	}
}
