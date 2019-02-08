package bdgraph

import (
	"fmt"
	"testing"
)

func TestWriteNode1(t *testing.T) {

	node := Node{name: "some name", number: 1}
	res := node.toGraphviz()

	if res != "\"some name\"\n" {
		t.Error("formatted node incorrectly")
	}
}

func TestWriteNode2(t *testing.T) {

	var lines = []string{
		"  1: apple",
		"  2: blueberry",
		"  5: cranberry ",
		"",
		"dependencies",
		" 1 -> 2",
		" 2 -> 5",
	}

	graph, err := ParseFile(lines)
	if err != nil {
		t.Error("unexpected error")
	}

	var res string = graph.nodes[1].toGraphviz()
	var expected string = "\"apple\"\n" +
		"\"apple\" -> \"blueberry\"\n"

	if res != expected {
		fmt.Printf("%s vs %s\n", res, expected)
		t.Error("formatted node incorrectly")
	}

	res = graph.nodes[2].toGraphviz()
	expected = "\"blueberry\"\n" +
		"\"blueberry\" -> \"cranberry\"\n"

	if res != expected {
		fmt.Printf("%s vs %s\n", res, expected)
		t.Error("formatted node incorrectly")
	}
}

func TestWriteGraph1(t *testing.T) {

	var lines = []string{
		"  1: apple",
		"  2: blueberry",
		"  5: cranberry ",
		"",
		"dependencies",
		" 1 -> 2",
		" 2 -> 5",
	}

	graph, err := ParseFile(lines)
	if err != nil {
		t.Error("unexpected error")
	}

	res := graph.toGraphviz()
	expected := graphvizHeader + `"apple"
"apple" -> "blueberry"
"blueberry"
"blueberry" -> "cranberry"
"cranberry"
` + graphvizFooter

	if res != expected {
		fmt.Printf("%s vs %s\n", res, expected)
		t.Error("formatted node incorrectly")
	}
}
