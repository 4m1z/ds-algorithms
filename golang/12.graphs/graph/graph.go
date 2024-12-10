package graph

type Graph struct {
	vertices []int
	edges    [][]int
}

func NewGraph() Graph {
	return Graph{
		vertices: []int{},
		edges:    [][]int{},
	}
}

func (g Graph) AddVertex(vertex int) Graph {
	g.vertices = append(g.vertices, vertex)

	if vertex < len(g.edges) {
		g.edges[vertex] = []int{}
	} else {
		g.edges = make([][]int, vertex+1)
	}

	return g
}

func (g Graph) AddEdge(from, to int) Graph {
	g.edges[from] = append(g.edges[from], to)
	g.edges[to] = append(g.edges[to], from)

	return g
}

func (g Graph) BfsTraverse() []int {
	queue := []int{0}
	result := []int{0}
	seen := make(map[int]bool)
	seen[0] = true
	edges := g.edges

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		connections := edges[vertex]
		for i := 0; i < len(connections); i++ {
			connection := connections[i]
			if !seen[connection] {
				queue = append(queue, connection)
				result = append(result, connection)
			}
			seen[connection] = true
		}
	}

	return result
}

func dfs(edges [][]int, seen map[int]bool, vertex int, res *[]int) {
	// kill condition
	if seen[vertex] {
		return
	}

	// mark as seen
	seen[vertex] = true
	// logic
	*res = append(*res, vertex)

	// visit other nodes
	conns := edges[vertex]
	for _, conn := range conns {
		if !seen[conn] {
			dfs(edges, seen, conn, res)
		}
	}
}

func (g Graph) DfsTraverse() (res []int) {
	seen := make(map[int]bool, len(g.vertices))

	dfs(g.edges, seen, 0, &res)

	return res
}
