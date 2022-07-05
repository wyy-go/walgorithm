package prim

import (
	"fmt"
	"sync"
)

type node struct {
	id string
}

type Edge struct {
	src    Node
	dst    Node
	weight float64
}

type prim struct {
	sync.RWMutex
	edge    map[string]map[string]float64
	nodeMap map[string]Node // record all the node in a graph
}

type Node interface {
	NodeID() string
}

func NewNode(id string) Node {
	return &node{id: id}
}

func (n *node) NodeID() string {
	return n.id
}

func NewEdge(src Node, dst Node, w float64) *Edge {
	return &Edge{src: src, dst: dst, weight: w}
}

func New() *prim {
	return &prim{
		edge:    make(map[string]map[string]float64),
		nodeMap: make(map[string]Node),
	}
}

func (g *prim) AddEdge(nodeID1 string, nodeID2 string, w float64) {
	g.Lock()
	defer g.Unlock()

	if nodeID1 == nodeID2 {
		panic("can't add same vertex in one edge")
		return
	}

	if w == 0 {
		panic("weight can't use 0")
		return
	}

	// record each vertex
	g.nodeMap[nodeID1] = NewNode(nodeID1)
	g.nodeMap[nodeID2] = NewNode(nodeID2)

	if _, ok := g.edge[nodeID1]; ok {
		g.edge[nodeID1][nodeID2] = w
	} else {
		tempMap := make(map[string]float64)
		tempMap[nodeID2] = w
		g.edge[nodeID1] = tempMap
	}
}

func (g *prim) Prim(src string) {
	knowNode := make(map[int]string)
	unknowNode := g.nodeMap
	tempEdgeMap := g.edge
	edgeMap := make(map[int]*Edge)
	totalWeight := float64(0)
	key := 0

	knowNode[key] = src
	delete(unknowNode, knowNode[key])
	key++
	var temp string

	// 找到与knowNodeMap里节点权值最小的节点，并将该节点加入nodeMap
	for len(unknowNode) > 0 {
		min := float64(1000000)
		for nodeID := range unknowNode {
			for _, v := range knowNode {
				if tempEdgeMap[nodeID][v] < min && tempEdgeMap[nodeID][v] != 0 {
					min = tempEdgeMap[nodeID][v]
					knowNode[key] = nodeID
					temp = v
				}
			}
		}

		n1 := NewNode(temp)
		n2 := NewNode(knowNode[key])
		edge := NewEdge(n1, n2, min)
		edgeMap[key-1] = edge

		// 从未知节点map删除已经找到的当前权值最小的节点
		delete(unknowNode, knowNode[key])
		totalWeight += min
		key++
	}

	fmt.Println(totalWeight)
	fmt.Println(knowNode)
}
