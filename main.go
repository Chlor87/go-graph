package main

import (
	"fmt"
	"strings"
)

func main() {
	g := NewGraph()
	v0 := &Vertex{"a"}
	v1 := &Vertex{"b"}
	v2 := &Vertex{"c"}
	v3 := &Vertex{"d"}

	g.AddVertex(v0)
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)

	g.AddEdge(v0, v1)
	g.AddEdge(v1, v2)
	g.AddEdge(v2, v3)
	g.AddEdge(v3, v0)

	fmt.Println(g.mat.String())
	fmt.Println()

	g.DFS(v1, func(v *Vertex, depth int) {
		fmt.Printf("%s%v\n", strings.Repeat(".", depth), v.Value)
	})

	fmt.Println()

	g.BFS(v1, func(v *Vertex, depth int) {
		fmt.Printf("%s%v\n", strings.Repeat(".", depth), v.Value)
	})

	var out []string
	for _, v := range g.ShortestPath(v3, v0) {
		out = append(out, v.Value.(string))
	}

	fmt.Println(strings.Join(out, " => "))

}
