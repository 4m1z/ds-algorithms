package main

import "fmt"

const maxN = 1e5

func Problem8() {
	///////////////////// input

	// n: number of cities
	// m: number of roads
	var n, m int
	var adj [maxN][]int

	fmt.Scan(&n, &m)

	for i := 0; i < m; i++ {
		var v, u int

		fmt.Scan(&v, &u)

		// just to be zerobase
		v--
		u--

		// fill adj list
		adj[v] = append(adj[v], u)
		adj[u] = append(adj[u], v)
	}
	//////////////////////////

	seen := make(map[int]bool, n)

	var conns [][]int

	// i = city
	for i := 0; i < n; i++ {
		fmt.Println(i, seen)
		if !seen[i] {
			var conn []int
			dfs(seen, adj, i, &conn)

			fmt.Println(conn)

			conns = append(conns, conn)
		}
	}

	fmt.Println(len(conns) - 1)
	fmt.Println(adj[:4])

	//if len(conns)-1 > 0 {
	//for i := range conns {
	//fmt.Printf("%d %d", conns[i][0], conns[i+1][0])
	//}
	//}

}

func dfs(seen map[int]bool, edges [maxN][]int, v int, res *[]int) {
	seen[v] = true

	/////////////
	*res = append(*res, v)
	////////////

	for conn := range edges[v] {
		if !seen[conn] {
			dfs(seen, edges, v, res)
		}
	}

}

func main() {
	Problem8()
}
