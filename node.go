package bdgraph

import (
	"fmt"
)

type Node struct {
	name     string
	number   int
	requires []*Node
	provides []*Node
}

func (node Node) show() {
	fmt.Printf("%s %d", node.name, node.number)

	fmt.Printf("\n  -> ")
	for _, n := range node.provides {
		fmt.Printf("%d ", n.number)
	}

	fmt.Printf("\n  <- ")
	for _, n := range node.requires {
		fmt.Printf("%d ", n.number)
	}

	fmt.Printf("\n")
}

func (node Node) equal(other Node) bool {
	return node.name == other.name &&
		node.number == other.number
}

func (node Node) toGraphviz() string {

	result := fmt.Sprintf("\"%s\"\n", node.name)

	for _, other := range node.provides {
		result += fmt.Sprintf(
			"\"%s\" -> \"%s\"\n",
			node.name,
			other.name)
	}

	return result
}

func addRequire(left, right *Node) {

	for _, v := range left.requires {
		if v.number == right.number {
			return
		}
	}

	left.requires = append(left.requires, right)
	right.provides = append(right.provides, left)
}

func addProvide(left, right *Node) {

	for _, v := range left.provides {
		if v.number == right.number {
			return
		}
	}

	left.provides = append(left.provides, right)
	right.requires = append(right.requires, left)
}
