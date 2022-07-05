package graph

import (
	"testing"
)

func TestNewUndirected(t *testing.T) {
	g := NewUndirected()
	//增加顶点
	for i := 0; i < 10; i++ {
		g.AddVertex(VertexId(i))
	}

	if g.VerticesCount() != 10 {
		t.Error("count err")
		return
	}

	//增加边
	for i := 0; i < 10; i++ {
		_ = g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}
	if !g.CheckEdge(2, 0) {
		t.Error("err")
		return
	}
	if g.CheckEdge(2, 1) {
		t.Error("err")
		return
	}
	if g.GetEdgeWeight(2, 0) != 1 {
		t.Error("err")
		return
	}

	//删除顶点
	if err := g.RemoveVertex(VertexId(2)); err != nil {
		t.Error(err)
		return
	}
	if g.CheckVertex(VertexId(2)) {
		t.Error("err")
		return
	}
	if g.CheckEdge(2, 0) {
		t.Error("err")
		return
	}

	//增加边，存在修改
	if err := g.AddEdge(3, 0, 1); err != nil {
		t.Error("err")
		return
	}
	if !g.CheckEdge(3, 0) {
		t.Error("err")
		return
	}

	//删除边
	if err := g.RemoveEdge(3, 0); err != nil {
		t.Error("err")
		return
	}
	if g.CheckEdge(3, 0) {
		t.Error("err")
		return
	}

	//统计边
	c := g.EdgesIter()
	countEdge := 0
	for _ = range c {
		countEdge++
	}

	if g.EdgesCount() != countEdge {
		t.Error("err")
		return
	}
}

func TestNewDirected(t *testing.T) {
	g := NewDirected()
	//增加顶点
	for i := 0; i < 10; i++ {
		g.AddVertex(VertexId(i))
	}
	if g.VerticesCount() != 10 {
		t.Error("count err")
		return
	}
	//增加边
	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}

	if !g.CheckEdge(2, 0) {
		t.Error("err")
		return
	}
	if g.CheckEdge(2, 1) {
		t.Error("err")
		return
	}
	if g.GetEdgeWeight(2, 0) != 1 {
		t.Error("err")
		return
	}
	//删除顶点
	if err := g.RemoveVertex(VertexId(2)); err != nil {
		t.Error(err)
		return
	}
	if g.CheckVertex(VertexId(2)) {
		t.Error("err")
		return
	}
	if g.CheckEdge(2, 0) {
		t.Error("err")
		return
	}
	//增加边，存在修改
	if err := g.AddEdge(3, 0, 1); err != nil {
		t.Error("err")
		return
	}
	if !g.CheckEdge(3, 0) {
		t.Error("err")
		return
	}
	//删除边
	if err := g.RemoveEdge(3, 0); err != nil {
		t.Error("err")
		return
	}
	if g.CheckEdge(3, 0) {
		t.Error("err")
		return
	}

	//统计边
	c := g.EdgesIter()
	countEdge := 0
	for _ = range c {
		countEdge++
	}

	if g.EdgesCount() != countEdge {
		t.Error("err")
		return
	}

	//查看
	//for p := range g.EdgesIter() {
	//	t.Log(p)
	//}
	//入度
	gp := g.GetPredecessors(VertexId(1)).VerticesIter()
	for p := range gp {
		if p != 3 && p != 5 && p != 7 && p != 9 {
			t.Error("err")
			return
		}
	}

	for p := range g.GetSuccessors(VertexId(4)).VerticesIter() {
		if p != VertexId(0) {
			t.Error("err")
			return
		}
	}
}
