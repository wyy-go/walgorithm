package prim

import (
	"fmt"
	"testing"
)

func TestPrim(t *testing.T) {
	g := New()
	g.AddEdge("A", "B", 7)
	g.AddEdge("A", "D", 5)
	g.AddEdge("B", "C", 8)
	g.AddEdge("B", "A", 7)
	g.AddEdge("B", "D", 9)
	g.AddEdge("B", "E", 7)
	g.AddEdge("C", "B", 8)
	g.AddEdge("C", "E", 5)
	g.AddEdge("D", "A", 5)
	g.AddEdge("D", "B", 9)
	g.AddEdge("D", "F", 6)
	g.AddEdge("D", "E", 15)
	g.AddEdge("D", "F", 6)
	g.AddEdge("E", "B", 7)
	g.AddEdge("E", "C", 5)
	g.AddEdge("E", "D", 15)
	g.AddEdge("E", "F", 8)
	g.AddEdge("E", "G", 9)
	g.AddEdge("F", "D", 6)
	g.AddEdge("F", "E", 8)
	g.AddEdge("F", "G", 11)
	g.AddEdge("G", "E", 9)
	g.AddEdge("G", "F", 11)

	fmt.Println(g)
	g.Prim("D")
}
