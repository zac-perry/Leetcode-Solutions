// 210: Course Schedule 2
// Topological sort problem 
/*
    Its asking to return the ordering of courses you should take to finish all course. In other words, return the topologically sorted order of the nodes in the graph (visit this before this before this etc).

    My approach here was using Kahns Algorithm, which utilizes BFS and an 'in degree' list. I first create an adj list to track the neighbors for each node. Then, I also make this in degree list, which tracks the number of 'incoming edges' to the node. 

    For BFS, to start, I add any nodes with an indegree of 0 to the queue. This indicates the "starting" nodes that have no 'prerequisites' to visit. Then, when going through normal BFS (popping from queue), I check the adj nodes to the current node (the neighbors) using the adj list and decrement the 'in degrees' for each one. 

    By decrementing, I am indicating that the node has 'lost' an incoming edge, or that the 'prerequisite' to visit the node has been satisfied. If any node now has an indegree of 0 after this, they are added to the queue. The curr node is also added to the result's slice, as it has been processed (in degree of 0) and it's neighbors have all been 'decremented'

    This will ensure that I am properly visiting all nodes in their intended order. Or in the case of this problem, I am completing each prerequisite class needed before I take the next one. 

    If I get to the end and the number of 'classes' that were stored in the results slice DOES NOT equal the number of courses, then there is a cycle. Otherwise, the answer is correct.

    Kahns Algorithm worked great here. I think the most confusing part of this problem for me is just tracking the different courses based on index within each slice. High level, it's not too hard but during the implementation, it's just a little difficult to keep track of. 
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
    queue := make([]int, 0)
    degree := make([]int, numCourses)
    results := make([]int, 0)
    adjList := make([][]int, numCourses)
    
    // Create adj list and indegree list
    for _, val := range(prerequisites) {
        adjList[val[1]] = append(adjList[val[1]], val[0])
        degree[val[0]] += 1
    }

    // add to queue, anything with a 0 indegree to start
    for i, _ := range degree {
       if degree[i] == 0 {
        queue = append(queue, i)
       } 
    }

    // BFS
    for len(queue) > 0 {
        // dequeue current node with 0 indegree
        curr := queue[0]
        queue = queue[1:] 

        // add to results
        results = append(results, curr)

        // go through the adj list for this current node
        // For each neighbor, decrement indegree
        // If the neighbor now has an indegree of 0, add to the queue
        for _, val := range adjList[curr] {
            degree[val] -= 1
            if degree[val] == 0 {
                queue = append(queue, val)
            }
        }
    }

    // cycle was detected
    if len(results) != numCourses {
        return []int{}
    }

    return results
}
