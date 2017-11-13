package bargi

import (
	"github.com/jlintvedt/bargi/graph"
	log "github.com/sirupsen/logrus"
)

type Node struct {
	st *graph.Graph
}

func NewNode() *Node {
	n := new(Node)
	n.st = graph.NewGraph()

	return n
}

func (n Node) Test() {
	log.Info("Node Test")
	n.st.Test()
}
