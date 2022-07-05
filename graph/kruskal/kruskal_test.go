package kruskal

import (
	"log"
	"testing"
)

func buildExampleGraph() Graph {
	var g = New()
	var vertexes = []string{"a", "b", "c", "d"}
	for _, vertex := range vertexes {
		g.AddVertex(vertex)
	}

	g.AddEdge("a", "b", 1)
	g.AddEdge("b", "c", 2)
	g.AddEdge("c", "a", 2)
	g.AddEdge("c", "d", 8)
	g.AddEdge("d", "a", 1)

	return g
}

func TestKruskal(t *testing.T) {
	g := buildExampleGraph()
	edges := Kruskals(&g)
	for _, edge := range edges {
		log.Printf("Edge from %s to %s with cost %v", edge.source, edge.sink, edge.weight)
	}
}
