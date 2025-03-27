// 1584. Min Cost to Connect All Points
// This is just a Minimum Spanning Tree problem
/*
    Here, I am using Prim's algorithm to create the minimum spanning tree and find the minimum cost. 

    Prim's algorithm is a greedy algorithm which will start at an arbitrary vertex and it will grow the MST by greedily choosing the edge to the next node that has the smallest weight. 

    The goal is to visit all nodes through edges that have the smallest weight

    So, in the case of this problem, we will traverse the graph to create a MST by visiting each node once, using the 'cheapest' edges. I will keep track of the total 'cost' spent to visit each node once.

    I create a min heap, starting with a random node that has a weight of 0 and goes to 0 (first node). I then 'visit' this node (if it hasn't been visited yet), add the edge weight to my total, and then look through the edges of this current node. 

    When looking through the edges of the node, if the next connected node hasn't been visited, I calculate the 'manhattan distance' (weight) for the edge and push it onto the min heap

    By doing this, the min heap will auto heapify itself, which will ensure the edge with the smallest 'manhattan distance' (wieght) will be at the top to be popped off. Again, the challenge here, is just ensuring we aren't visiting something more than once, creating a cycle.

*/
type Edge struct {
    weight int
    to     int
}

// Implement golangs min heap interface
type MinHeap []Edge

func (minHeap MinHeap) Len() int {
	return len(minHeap)
}

func (minHeap MinHeap) Less(i, j int) bool {
    return minHeap[i].weight < minHeap[j].weight
}

func (minHeap MinHeap) Swap(i, j int) {
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}
func (minHeap *MinHeap) Push(x any) {
	*minHeap = append(*minHeap, x.(Edge))
}

func (minHeap *MinHeap) Pop() any {
	placeholder := *minHeap
	length := len(placeholder)
	value := placeholder[length-1]
	*minHeap = placeholder[0 : length-1]

	return value
}

func minCostConnectPoints(points [][]int) int {
    // Create min heap, visited array
    minHeap := &MinHeap{}
    visited := make([]bool, len(points))
    used := 0
    cost := 0

    // initialize, push starting edge 0,0
    heap.Init(minHeap)
    heap.Push(minHeap, Edge{0, 0})

    // Prims Algorithm
    // Pop edge with min weight off the min heap, set visited, process it's edges
    for used < len(points){
        curr := heap.Pop(minHeap).(Edge)
        if visited[curr.to] {
            continue
        }

        visited[curr.to] = true
        cost += curr.weight
        used++

        // loop and push edges (manhattan distance being the weight)
        for next := 0; next < len(points); next++ {
            if !visited[next] {
                distance := int(math.Abs(float64(points[curr.to][0]-points[next][0])) + math.Abs(float64(points[curr.to][1]-points[next][1])))
                heap.Push(minHeap, Edge{distance, next})
            }
        }
    }

    return cost
}
