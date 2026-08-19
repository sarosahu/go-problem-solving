package graph

// Node represents a vertex in the graph.
type Node struct {
	ID int
	Edges []*Node // Slices of pointers to other Node instances
}

// NewNode is a constructor for creating a single Node.
func NewNode(id int) *Node {
	return &Node{
		ID: id,
		Edges: []*Node{},  // Initializes an empty slice
	}
}

// Graph represents the collection of nodes.
type Graph struct {
	Nodes []*Node
}

// NewGraph is a constructor that initalizes a graph with 'n' nodes.
func NewGraph(n int) *Graph {
	graph := &Graph{
		Nodes: make([]*Node, n),  // Pre-allocates a slice of size
	}
	for i := range n {
		graph.Nodes[i] = NewNode(i)
	}

	return graph
}

// Fetch the nodes
func (g *Graph) getNodes() []*Node {
	return g.Nodes
}