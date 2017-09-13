package graph

import (
	"fmt"
)

///// ///// ///// ///// ///// ///// GRAPH ///// ///// ///// ///// ///// /////

// Graph TODO:Write
type Graph struct {
	meta graphMeta
	vert map[int]vertex
	edge map[int]edge
}

// NewGraph TODO:Write
func NewGraph() *Graph {
	s := new(Graph)

	return s
}

func (g Graph) Test() {
	fmt.Println("Struct method test")
}

///// ///// ///// ///// ///// ///// GRAPH META ///// ///// ///// ///// ///// /////
type graphMeta struct {
	nodeID int
	maxID  int // Total number of worker nodes in complete graph
}

///// ///// ///// ///// ///// ///// VERTEX ///// ///// ///// ///// ///// /////
type vertex struct {
	id    int
	vType vertexType
	feats map[string]string
	edges []int
}

type vertexType int

const (
	person vertexType = iota
	featPlace
	featJobTitle
)

///// ///// ///// ///// ///// ///// EDGE ///// ///// ///// ///// ///// /////
type edge struct {
	id    int
	eType edgeType
	from  int
	to    int
}

type edgeType int

const (
	follows edgeType = iota
	livesIn
	hasTitle
)
