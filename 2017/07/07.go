package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type node struct {
	name     string
	weight   int
	children []node
}

func main() {
	lines := getDataFromFile("input.txt")
	nodes := createNodesFromData(lines)
	tree := createTreeFromNodes(nodes)
	balanceTree(tree)
}

func balanceTree(tree node) int {
	if len(tree.children) == 0 {
		return tree.weight
	}

	weight := 0
	weight += tree.weight
	weights := make([]int, len(tree.children))
	for i, branch := range tree.children {
		weights[i] = balanceTree(branch)
	}

	balanced := true
	for _, w := range weights {
		if weights[0] != w {
			balanced = false
		}
		weight += w
	}

	if !balanced {
		for i, w := range weights {
			fmt.Println(w, tree.children[i].name, tree.children[i].weight)
		}
	}

	return weight
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDataFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 1)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	check(err)
	return lines
}

func createLeaf(line string) node {
	namePattern := regexp.MustCompile(`[a-z]*`)
	weightPattern := regexp.MustCompile(`[0-9]+`)
	name := namePattern.FindString(line)
	weightString := weightPattern.FindString(line)
	weight, _ := strconv.Atoi(weightString)
	var children []node
	return node{name, weight, children}
}
func createNode(line string) node {
	parentSeparator := regexp.MustCompile("-> ")
	str := parentSeparator.Split(line, -1)
	node := createLeaf(str[0])
	node.children = createChildrenFromLine(str[1])
	return node
}

func createChildrenFromLine(line string) []node {
	childSeparator := regexp.MustCompile(", ")
	children := childSeparator.Split(line, -1)
	nodes := make([]node, 0)
	var emptyChildren []node
	for _, child := range children {
		nodes = append(nodes, node{name: child, weight: 0, children: emptyChildren})
	}
	return nodes
}

func createNodesFromData(lines []string) []node {
	parentSeparator := regexp.MustCompile("-> ")
	nodes := make([]node, 0)
	for _, line := range lines {
		var newNode node
		if parentSeparator.MatchString(line) {
			newNode = createNode(line)
		} else {
			newNode = createLeaf(line)
		}
		nodes = append(nodes, newNode)
	}
	return nodes
}

func getChildData(curr node, rest []node) (node, []node) {
	for i, r := range rest {
		if curr.name == r.name {
			rest = append(rest[:i], rest[i+1:]...)
			return r, rest
		}
	}
	return curr, rest
}

func createTree(curr node, rest []node) (node, []node) {
	var n []node
	if len(rest) == 0 {
		return curr, n
	}

	for i, child := range curr.children {
		curr.children[i], rest = getChildData(child, rest)
		curr.children[i], rest = createTree(curr.children[i], rest)
	}

	return curr, rest
}

func createTreeLoop(nodes []node) node {
	counter := 0
	var res node
	for len(nodes) > 1 {
		counter++
		res = nodes[0]
		res, nodes = createTree(res, nodes[1:])
		nodes = append(nodes, res)
	}
	return res
}

func createTreeFromNodes(nodes []node) node {
	nodeCopy := append([]node(nil), nodes...)
	return createTreeLoop(nodeCopy[1:])
}
