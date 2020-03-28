package main

// Vertex ...
type Vertex struct {
	Value interface{}
}

// Graph ...
type Graph struct {
	mat      Mat
	vertices []*Vertex
}

type cbFunc func(*Vertex, int)

func noopCB(*Vertex, int) {}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{NewMat(0, 0), []*Vertex{}}
}

// AddVertex ...
func (g *Graph) AddVertex(v *Vertex) {
	if g.vertexIdx(v) != -1 {
		return
	}
	g.mat.Grow()
	g.vertices = append(g.vertices, v)
}

// AddEdge ...
func (g *Graph) AddEdge(from, to *Vertex) {
	fromIdx, toIdx := g.vertexIdx(from), g.vertexIdx(to)
	if fromIdx == -1 || toIdx == -1 {
		panic("failed to locate vertex")
	}
	// unidirectional
	g.mat[fromIdx][toIdx] = 1
}

// Len ...
func (g *Graph) Len() int {
	return len(g.vertices)
}

// DFS ...
func (g *Graph) DFS(start *Vertex, cb cbFunc) {
	g.dfs(g.vertexIdx(start), cb, 0, make([]bool, g.Len()))
}

func (g *Graph) dfs(curr int, cb cbFunc, depth int, visited []bool) {
	visited[curr] = true
	cb(g.vertices[curr], depth)
	depth++
	for i, v := range g.mat[curr] {
		if v == 1 && !visited[i] {
			g.dfs(i, cb, depth, visited)
		}
	}
}

// BFS ...
func (g *Graph) BFS(start *Vertex, cb cbFunc) {
	g.bfs(g.vertexIdx(start), cb)
}

func (g *Graph) bfs(start int, cb cbFunc) (parents []int, visited []bool) {
	q, l := []int{start}, g.Len()
	var depth int
	parents, visited = make([]int, l), make([]bool, l)
	visited[start] = true
	cb(g.vertices[start], depth)

	for len(q) > 0 {
		depth++
		var curr int
		curr, q = q[0], q[1:]
		for i, v := range g.mat[curr] {
			if v == 0 || visited[i] {
				continue
			}
			cb(g.vertices[i], depth)
			visited[i], parents[i] = true, curr
			q = append(q, i)
		}
	}
	return
}

// ShortestPath ...
func (g *Graph) ShortestPath(from, to *Vertex) (path []*Vertex) {
	fromIdx, toIdx := g.vertexIdx(from), g.vertexIdx(to)
	if fromIdx == -1 || toIdx == -1 {
		panic("failed to locate vertex")
	}
	parents, visited := g.bfs(fromIdx, noopCB)
	if !visited[toIdx] {
		return
	}
	for i := toIdx; i != fromIdx; i = parents[i] {
		path = append(path, g.vertices[i])
	}
	path = append(path, g.vertices[fromIdx])
	reversePath(path)
	return
}

func (g *Graph) vertexIdx(v *Vertex) int {
	for i, vv := range g.vertices {
		if vv == v {
			return i
		}
	}
	return -1
}

func reversePath(in []*Vertex) {
	for i := len(in)/2 - 1; i >= 0; i-- {
		opp := len(in) - 1 - i
		in[i], in[opp] = in[opp], in[i]
	}
}
