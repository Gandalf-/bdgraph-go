package bdgraph

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Nodes map[int]*Node

type Position int

const (
	Declaration Position = 0
	Options     Position = 1
	Dependency  Position = 2
)

func ParseFile(lines []string) (Graph, error) {

	nodes := Nodes{}
	options := []*Option{}
	position := Declaration

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		if line == "options" {
			position = Options
			continue

		} else if line == "dependencies" {
			position = Dependency
			continue
		}

		switch position {

		case Declaration:
			node, err := ParseDeclaration(line)
			if err != nil {
				return Graph{}, err
			}
			nodes[node.number] = &node

		case Options:
			new_options, err := ParseOption(line)
			if err != nil {
				return Graph{}, err
			}
			for i := range new_options {
				options = append(options, &new_options[i])
			}

		case Dependency:
			err := ParseDependency(line, nodes)
			if err != nil {
				return Graph{}, err
			}
		}
	}

	return Graph{nodes: nodes, options: options}, nil
}

func ParseDeclaration(line string) (Node, error) {
	items := strings.Split(line, ":")

	if len(items) != 2 {
		return Node{}, fmt.Errorf(
			"could not parse declaration: %s", line)
	}

	item := strings.TrimSpace(items[0])
	number, err := strconv.Atoi(item)
	if err != nil {
		return Node{}, fmt.Errorf(
			"could not parse %s as a number", item)
	}

	if number < 1 {
		return Node{}, fmt.Errorf(
			"number in declaration less than 1: %d", number)
	}

	name := strings.TrimSpace(items[1])

	return Node{name: name, number: number}, nil
}

func ParseOption(line string) ([]Option, error) {

	words := strings.Split(line, " ")
	options := make([]Option, 0)

	for i := range words {
		item := strings.TrimSpace(words[i])

		switch item {
		case "":
			// extra whitespace is trimmed to empty string

		case "cleanup":
			options = append(options, OptionCleanup)

		case "circular":
			options = append(options, OptionCircular)

		case "color_next":
			options = append(options, OptionNext)

		case "color_urgent":
			options = append(options, OptionUrgent)

		case "color_complete":
			options = append(options, OptionComplete)

		default:
			return options, fmt.Errorf("unrecongized option: %s", item)
		}
	}

	return options, nil
}

func ParseDependency(line string, nodes Nodes) error {

	require := false
	provide := false
	var items []string

	if strings.Contains(line, "->") {
		provide = true
		items = strings.Split(line, "->")

	} else if strings.Contains(line, "<-") {
		require = true
		items = strings.Split(line, "<-")
	}

	if len(items) != 2 {
		return fmt.Errorf("could not parse dependency: %s\n", line)
	}

	// get the left side
	left, err := strconv.Atoi(strings.TrimSpace(items[0]))
	if err != nil {
		return errors.New(
			"left side of dependency must be a single number")
	}

	// find the target node
	node, ok := nodes[left]
	if !ok {
		return fmt.Errorf("reference to %d not found\n", left)
	}

	// get the right side
	right, err := ParseNumberList(items[1])
	if err != nil {
		return err
	}

	if provide {
		for _, v := range right {
			provision, ok := nodes[v]
			if !ok {
				return fmt.Errorf("reference to %d not found\n", v)
			}

			if node.number == provision.number {
				return fmt.Errorf(
					"%d cannot refer to itself", node.number)
			}
			addProvide(node, provision)
		}
	}

	if require {
		for _, v := range right {
			requirement, ok := nodes[v]
			if !ok {
				return fmt.Errorf("reference to %d not found\n", v)
			}

			if node.number == requirement.number {
				return fmt.Errorf(
					"%d cannot refer to itself", node.number)
			}
			addRequire(node, requirement)
		}
	}

	return nil
}

func ParseNumberList(line string) ([]int, error) {

	items := strings.Split(line, ",")
	numbers := make([]int, 0)

	for _, item := range items {
		if item == "" {
			continue
		}

		number, err := strconv.Atoi(strings.TrimSpace(item))
		if err != nil {
			return numbers, fmt.Errorf(
				"could not parse number %s as a number\n", item)
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	s := "1: @personal skills"
	a := strings.Split(s, ":")

	for i := range a {
		fmt.Println(a[i])
	}
}
