package bdgraph

import (
	"fmt"
	"sort"
)

const graphvizHeader string = `digraph g{
rankdir=LR;
ratio=fill;
node [style=filled];
`

const graphvizFooter string = "}\n"

type Graph struct {
	nodes   Nodes
	options []*Option
}

func (graph Graph) show() {
	fmt.Println("nodes")
	for _, node := range graph.nodes {
		node.show()
	}
	fmt.Println("options")
	for _, option := range graph.options {
		option.show()
	}
}

func (graph Graph) orderNodes() []int {
	/*
		storing the nodes in a hash map means that iteration over the keys is
		non-determinisitic. we don't always want that, particularly for writing
		out files
	*/
	keys := make([]int, 0)
	for key, _ := range graph.nodes {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func (graph Graph) toGraphviz() string {

	body := ""
	ordering := graph.orderNodes()

	for _, i := range ordering {
		body += graph.nodes[i].toGraphviz()
	}

	return graphvizHeader + body + graphvizFooter
}

func (graph Graph) toFile() string {

	return ""
}
