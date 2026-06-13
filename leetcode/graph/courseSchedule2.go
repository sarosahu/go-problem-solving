package graph

/*
 *210. Course Schedule II
 There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
 You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you 
 must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid 
answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]

Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.
 */

// { First approach -- recursive 
type State int
const (
    Unvisited State = iota
    Visiting
    Visited
)

func findOrderTopologicalOrderR(numCourses int, prerequisites [][]int) []int {
    graph := NewGraph(numCourses)
    states := make([]State, numCourses)
    addPrerequisites3(graph, prerequisites)

    stack := &Stack[int]{}
    for i := 0; i < len(graph.getNodes()); i++ {
        curr := graph.getNodes()[i]
        if states[curr.ID] == Unvisited {
            hasCycle := topologicalOrderUtil(curr, stack, states)
            if hasCycle {
                return []int{}
            }
        }
    }
    topoOrder := make([]int, numCourses)
    idx := 0
    for !stack.IsEmpty() {
        topoOrder[idx], _ = stack.Pop()
        idx++
    }
    return topoOrder
}

func addPrerequisites3(graph *Graph, prerequisites [][]int) {
    for i := 0; i < len(prerequisites); i++ {
        srcId, destId := prerequisites[i][1], prerequisites[i][0]
        srcNode, destNode := graph.Nodes[srcId], graph.Nodes[destId]
        srcNode.Edges = append(srcNode.Edges, destNode)
    }
}

func topologicalOrderUtil(curr *Node, stack *Stack[int], states []State) bool {
    states[curr.ID] = Visiting
    cycleFound := false
    for _, edgeNode := range curr.Edges {
        if states[edgeNode.ID] == Unvisited {
            cycleFound = topologicalOrderUtil(edgeNode, stack, states)
        } else if states[edgeNode.ID] == Visiting {
            cycleFound = true
        }
        if cycleFound {
            return cycleFound
        }
    }
    states[curr.ID] = Visited
    stack.Push(curr.ID)
    
    return cycleFound
}
// } First approach -- recursive 

// { Second approach -- Kahn's algorithm
func findOrderTopologicalOrderKahn(numCourses int, prerequisites [][]int) []int {
    graph := NewGraph(numCourses)
    inDegree := make([]int, numCourses)
    addPrerequisites2(graph, prerequisites, inDegree)
    orders := make([]int, numCourses)

    queue := Queue[*Node]{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue.Enqueue(graph.Nodes[i])
        }
    }
    count := 0
    for !queue.IsEmpty() {
        curr, exist := queue.Dequeue()
        if exist {
            orders[count] = curr.ID
            count++
        }
        for i := 0; i < len(curr.Edges); i++ {
            edgeNode := curr.Edges[i]
            inDegree[edgeNode.ID]--
            if inDegree[edgeNode.ID] == 0 {
                queue.Enqueue(edgeNode)
            }
        }
    }
    if count != numCourses {
        return []int{}
    }
    return orders
}

func addPrerequisites2(graph *Graph, prerequisites [][]int, inDegree []int) {
    for i := 0; i < len(prerequisites); i++ {
        srcId, destId := prerequisites[i][1], prerequisites[i][0]
        srcNode, destNode := graph.Nodes[srcId], graph.Nodes[destId]
        srcNode.Edges = append(srcNode.Edges, destNode)
        inDegree[destId] += 1
    }
}
// } Second approach -- Kahn's algorithm