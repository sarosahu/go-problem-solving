package graph

/*
 * 207. Course Schedule
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
 * You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must
 * take course bi first if you want to take course ai.

 * For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
 * Return true if you can finish all courses. Otherwise, return false.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.

Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := NewGraph(numCourses)
	inDegree := make([]int, numCourses)
	addPrerequisites(graph, prerequisites, inDegree)

	queue := Queue[*Node]{}
	for i := 0; i < numCourses; i++ {
		curr := graph.Nodes[i]
		if inDegree[curr.ID] == 0 {
			queue.Enqueue(curr)
		}
	}

	//visited := make(map[int]struct{})
	visitedCount := 0
	for !queue.IsEmpty() {
		curr, exist := queue.Dequeue()
		if !exist {
			continue
		}
		//visited[curr.ID] = struct{}{}
		visitedCount++

		/*if visitedCount == numCourses {
			return true
		}*/
		for i := 0; i < len(curr.Edges); i++ {
			edgeNode := curr.Edges[i]
			inDegree[edgeNode.ID]--
			if inDegree[edgeNode.ID] == 0 {
				queue.Enqueue(edgeNode)
			}
		}
	}
	return visitedCount == numCourses
}

func addPrerequisites(graph *Graph, prerequisites [][]int, inDegree []int) {
	for i := range prerequisites {
		srcId, destId := prerequisites[i][1], prerequisites[i][0]
		src := graph.Nodes[srcId]
		dest := graph.Nodes[destId]
		src.Edges = append(src.Edges, dest)
		inDegree[dest.ID] += 1
	}
}