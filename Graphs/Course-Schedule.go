// 207. Course Schedule
// Graph DFS problem -- Essentially a topological sort to find cycles
// Approach here involves creating a 'visited' list for each node
// Then, creating an adjacency list to represent the graph
// If, when looking at each node in the graph, it was visited, return false
// Otherwise, mark it as 'seen / being processed' and continue
// When done looking at the children, mark as 2, or 'fully processed'
// Need to loop over all courses and call dfs in order to handle any unconnected nodes
func canFinish(numCourses int, prerequisites [][]int) bool {

    // visited list: index = value of the course
    nodes := make([]int, numCourses)

    // Need to make the adj list to represent the graph here. 
    // Loop over the prerequisites and do this
    adjList := make([][]int, numCourses)

    for i, val := range(prerequisites) {
        fmt.Println(i, val)
        adjList[val[1]] = append(adjList[val[1]], val[0])
    }

    var dfs func(visited []int, graph [][]int, node int) bool 
    dfs = func(visited []int, graph [][]int, node int) bool {
        // current node is already 'being processed', return false (cycle).
        if visited[node] == 1 {
            return false
        }

        // current node has been fully processed.
        if visited[node] == 2 {
            return true
        }

        // now processing current node.
        visited[node] = 1

        for _, neighbor := range graph[node] {
            if dfs(visited, graph, neighbor) == false {
                return false
            }
        }

        // node is done being processed.
        visited[node] = 2
        return true
    }

    // Loop and call dfs on each course number. 
    // NOTE: needed to ensure we visit even unconnected nodes
    for i := range numCourses {
        if nodes[i] == 0 {
            if dfs(nodes, adjList, i) == false {
                return false
            }
        }
    }

    return true
}
