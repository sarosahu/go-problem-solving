/*
*
* Graph Valid Tree

* Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes),
* write a function to check whether these edges make up a valid tree.

Ex: 1
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]

Output: true

Ex: 2
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]

Output: false
*/
package graph

// { DFS approach
 func validTree(n int, edges [][]int) bool {
	if len(edges) != n - 1 {
		return false
	}
	graph := NewGraph(n)
	for i := range edges {
		graph.addEdge(edges[i][0], edges[i][1])
	}

	seen := make(map[int]struct{})

	return validTreeDfs(0, -1, graph, seen) && len(seen) == n
}

func validTreeDfs(node int, parent int, graph *Graph, seen map[int]struct{}) bool {
	if _, exist := seen[node]; exist {
		return false
	}
	seen[node] = struct{}{}
	for i := range graph.Nodes[node].Edges {
		neiId := graph.Nodes[node].Edges[i].ID
		if neiId != parent {
			if (!validTreeDfs(neiId, node, graph, seen)) {
				return false
			}
		}
	}
	return true
}
// } DFS approach

// {{ Stack approach (iterative) -->
func validTreeUsingStack(n int, edges [][]int) bool {
    graph := NewGraph(n)
	for i := range edges {
		edge := edges[i]
		graph.addEdge(edge[0], edge[1])
	}
	parent := make(map[int]int)
	parent[0] = -1
	stack := Stack[int]{}
	stack.Push(0)

	for !stack.IsEmpty() {
		currNodeId, _ := stack.Pop()

		for i := range graph.Nodes[currNodeId].Edges {
			neiNode := graph.Nodes[currNodeId].Edges[i]
			if pVal, exist := parent[currNodeId]; exist {
				if pVal == neiNode.ID {
					continue
				}
			}
			if _, exist := parent[neiNode.ID]; exist {
				return false
			}
			parent[neiNode.ID] = currNodeId
			stack.Push(neiNode.ID)
		}
	}
	return len(parent) == n
}
// }} Stack approach <-

// {{ BFS approach -->
func validTreeBfs(n int, edges [][]int) bool {
	if len(edges) != n - 1 {
		return false
	}
	graph := NewGraph(n)
	for i := range edges {
		graph.addEdge(edges[i][0], edges[i][1])
	}
	queue := Queue[int]{}
	queue.Enqueue(0)
	parent := make(map[int]int)
	parent[0] = -1

	for !queue.IsEmpty() {
		currNodeId, _ := queue.Dequeue()
		for i := range graph.Nodes[currNodeId].Edges {
			neiId := graph.Nodes[currNodeId].Edges[i].ID
			if neiId == parent[currNodeId] {
				continue
			}
			if _, exist := parent[neiId]; exist {
				return false
			}
			queue.Enqueue(neiId)
			parent[neiId] = currNodeId
		}
	}
	return len(parent) == n
}
// }} BFS approach <--

// {{ Union - Find (Disjoint set) approach -->
type DSU struct {
	Parent []int
	Rank []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		parent[i] = i
		rank[i] = 1
	}
	return &DSU{
		Parent: parent,
		Rank: rank,
	}
}

func (dsu *DSU) Find(node int) int {
	if dsu.Parent[node] != node {
		dsu.Parent[node] = dsu.Find(dsu.Parent[node])
	}
	return dsu.Parent[node]
}

func (dsu *DSU) Union(u, v int)  bool {
	pu, pv := dsu.Find(u), dsu.Find(v)
	// Has Cycle (pu == pv) ? then return false
	if pu == pv {
		return false
	}
	if dsu.Rank[pu] < dsu.Rank[pv] {
		dsu.Parent[pu] = pv
	} else if dsu.Rank[pv] < dsu.Rank[pu] {
		dsu.Parent[pv] = pu
	} else {
		dsu.Parent[pv] = pu
		dsu.Rank[pu] += 1
	}
	return true
}

func validTreeDS(n int, edges [][]int) bool {
	if len(edges) != n - 1 {
		return false
	}
	dsu := NewDSU(n)
	for i := range edges {
		if !dsu.Union(edges[i][0], edges[i][1]) {
			return false
		}
	}
	return true
}
// }} Union - Find (Disjoint set) approach <--

func (graph *Graph) addEdge(n1, n2 int) {
	src, dest := graph.Nodes[n1], graph.Nodes[n2]
	src.Edges = append(src.Edges, dest)
	dest.Edges = append(dest.Edges, src)
}