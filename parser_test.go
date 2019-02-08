package bdgraph

import (
	// "fmt"
	"reflect"
	"testing"
)

// ParseDeclaration

func TestParseDeclaration1(t *testing.T) {
	// correct parsing

	expected := Node{name: "some name", number: 1}

	var line string = "   1:  some name      "
	node, err := ParseDeclaration(line)

	if err != nil {
		t.Errorf("unexpected error parsing: %s\n", line)
	}

	if !node.equal(expected) {
		t.Error("didn't parse node values correctly")
	}
}

func TestParseDeclaration2(t *testing.T) {
	// no ':'

	var line string = "   1  some name      "
	_, err := ParseDeclaration(line)

	if err == nil {
		t.Error("didn't get expected parse error")
	}
}

func TestParseDeclaration3(t *testing.T) {
	// wrong number of ':'

	var line string = " :  1  some : name      "
	_, err := ParseDeclaration(line)

	if err == nil {
		t.Error("didn't get expected parse error")
	}
}

func TestParseDeclaration4(t *testing.T) {
	// 'number' isn't an int

	var line string = "   apple  : some name      "
	_, err := ParseDeclaration(line)

	if err == nil {
		t.Error("didn't get expected parse error")
	}
}

func TestParseDeclaration5(t *testing.T) {
	// 'number' isn't an int

	var line string = "   -1  : some name      "
	_, err := ParseDeclaration(line)

	if err == nil {
		t.Error("didn't get expected parse error")
	}
}

// ParseOption

func TestParseOption1(t *testing.T) {
	// valid, one item

	line := "  color_next  "
	o, err := ParseOption(line)

	if err != nil {
		t.Error("unexpected error")
	}

	expected := []Option{OptionNext}
	if !reflect.DeepEqual(o, expected) {
		t.Error("got the wrong option")
	}
}

func TestParseOption2(t *testing.T) {
	// valid, two items

	line := "  color_next  color_complete"
	o, err := ParseOption(line)

	if err != nil {
		t.Error("unexpected error")
	}

	expected := []Option{OptionNext, OptionComplete}
	if !reflect.DeepEqual(o, expected) {
		t.Error("got the wrong options")
	}
}

func TestParseOption3(t *testing.T) {
	// invalid

	line := "  junk"
	_, err := ParseOption(line)

	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseOption4(t *testing.T) {
	// mix of valid and invalid

	line := "  circular   junk"
	o, err := ParseOption(line)

	expected := []Option{OptionCircular}
	if !reflect.DeepEqual(o, expected) {
		t.Error("got the wrong options")
	}

	if err == nil {
		t.Error("didn't get expected error")
	}
}

// ParseNumberList

func TestParseNumberList1(t *testing.T) {
	// single value

	line := "1"
	ns, err := ParseNumberList(line)

	if err != nil {
		t.Error("unexpected error")
	}

	es := []int{1}
	if !reflect.DeepEqual(ns, es) {
		t.Error("didn't get expected value")
	}
}

func TestParseNumberList2(t *testing.T) {
	// multiple values

	line := "1,22"
	ns, err := ParseNumberList(line)

	if err != nil {
		t.Error("unexpected error")
	}

	es := []int{1, 22}
	if !reflect.DeepEqual(ns, es) {
		t.Error("didn't get expected value")
	}
}

func TestParseNumberList3(t *testing.T) {
	// multiple values with spaces

	line := "  1, 22  "
	ns, err := ParseNumberList(line)

	if err != nil {
		t.Error("unexpected error")
	}

	es := []int{1, 22}
	if !reflect.DeepEqual(ns, es) {
		t.Error("didn't get expected value")
	}
}

func TestParseNumberList4(t *testing.T) {
	// mixed valid and invalid

	line := "1,junk"
	_, err := ParseNumberList(line)

	if err == nil {
		t.Error("didn't get expected error")
	}
}

// ParseDependency

func TestParseDependency1(t *testing.T) {
	// invalid type, not provide or require

	var line string = " junk "
	var nodes Nodes

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency2(t *testing.T) {
	// too many elements

	var line string = " a -> b -> c "
	var nodes Nodes

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency3(t *testing.T) {
	// left not a number

	var line string = " junk -> 1,  2, 3,4"
	var nodes Nodes

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency4(t *testing.T) {
	// left not found

	var line string = " 99 -> 1,2 "
	nodes := Nodes{
		1: &Node{name: "a", number: 1},
		2: &Node{name: "b", number: 2},
	}

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency5(t *testing.T) {
	// right side can't be parsed

	var line string = " 1 -> 2,junk "
	nodes := Nodes{
		1: &Node{name: "a", number: 1},
		2: &Node{name: "b", number: 2},
	}

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency6(t *testing.T) {
	// reference in right not found

	var line string = " 1 -> 2, 99 "
	nodes := Nodes{
		1: &Node{name: "a", number: 1},
		2: &Node{name: "b", number: 2},
	}

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("didn't get expected error")
	}
}

func TestParseDependency7(t *testing.T) {
	// add multiple provides

	var line string = " 1 -> 2,3 "
	a := &Node{name: "a", number: 1}
	b := &Node{name: "b", number: 2}
	c := &Node{name: "c", number: 3}

	nodes := Nodes{
		1: a,
		2: b,
		3: c,
	}

	err := ParseDependency(line, nodes)
	if err != nil {
		t.Error("unexpected error")
	}

	if !reflect.DeepEqual(a.provides, []*Node{b, c}) {
		t.Error("a provides")
	}

	if !reflect.DeepEqual(b.requires, []*Node{a}) {
		t.Error("b requires")
	}

	if !reflect.DeepEqual(c.requires, []*Node{a}) {
		t.Error("c requires")
	}
}

func TestParseDependency8(t *testing.T) {
	// add multiple requires

	var line string = " 1 <- 2,3 "
	a := &Node{name: "a", number: 1}
	b := &Node{name: "b", number: 2}
	c := &Node{name: "c", number: 3}

	nodes := Nodes{
		1: a,
		2: b,
		3: c,
	}

	err := ParseDependency(line, nodes)
	if err != nil {
		t.Error("unexpected error")
	}

	if !reflect.DeepEqual(a.requires, []*Node{b, c}) {
		t.Error("a requires")
	}

	if !reflect.DeepEqual(b.provides, []*Node{a}) {
		t.Error("b provides")
	}

	if !reflect.DeepEqual(c.provides, []*Node{a}) {
		t.Error("c provides")
	}
}

func TestParseDependency9(t *testing.T) {
	// invalid self reference

	var line string = " 1  <- 1 "
	nodes := Nodes{
		1: &Node{name: "a", number: 1},
		2: &Node{name: "b", number: 2},
	}

	err := ParseDependency(line, nodes)
	if err == nil {
		t.Error("unexpected error")
	}
}

// ParseFile

func TestParseFile1(t *testing.T) {

	var lines = []string{
		"# comment",
		"  ",
		"  1: apple",
		"  2: blueberry",
		"  5: cranberry ",
		"",
		"options",
		" circular color_next",
		" color_complete  ",
		"",
		"dependencies",
		" 1 -> 2",
		" 2 -> 5",
	}

	graph, err := ParseFile(lines)
	if err != nil {
		t.Error("unexpected error")
	}

	expected := []*Option{&OptionCircular, &OptionNext, &OptionComplete}
	if !reflect.DeepEqual(graph.options, expected) {
		t.Error("options not parsed correctly")
	}

	first, ok := graph.nodes[1]
	if !ok {
		t.Error("a was not parsed")
	}
	if first.name != "apple" {
		t.Error("a parsed incorrectly")
	}

	second, ok := graph.nodes[2]
	if !ok {
		t.Error("b was not parsed")
	}
	if second.name != "blueberry" {
		t.Error("b parsed incorrectly")
	}

	third, ok := graph.nodes[5]
	if !ok {
		t.Error("c was not parsed")
	}
	if third.name != "cranberry" {
		t.Error("c parsed incorrectly")
	}

	if first.provides[0] != second {
		t.Error("a dependencies not parsed")
	}
	if second.provides[0] != third {
		t.Error("b dependencies not parsed")
	}
}
