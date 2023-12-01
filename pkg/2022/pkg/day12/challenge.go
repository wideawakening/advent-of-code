package day12

import (
	"fmt"
	"os"
	"strings"
)

/*
area elevation from a..z lowest to highest
nodeValue position : S
best signal: E

go from S to E with least steps (up,down,left,right) always to at most 1 level higher (all lower options) than n nodeValue S lvl.

get fewest steps required
*/
func Star1(inputFileName string) int {
	dataMap, _ := LoadMap(inputFileName)

	startPosition, _ := GetPositionFor(dataMap, "S")
	endPosition, _ := GetPositionFor(dataMap, "E")
	graph := CreateGraph(dataMap)

	graph.GetCountToFrom(startPosition, endPosition)
	return 0
}

func CreateGraph(dataMap [][]Node) *Graph {
	graph := &Graph{}

	for row, rowNodes := range dataMap {
		for column, _ := range rowNodes {
			currentNode := dataMap[row][column]
			graph.AddVertex(currentNode)

			currentPosition := currentNode.position
			// up
			if currentPosition[0]-1 >= 0 {
				node := dataMap[currentPosition[0]-1][currentPosition[1]]
				graph.AddVertex(node)
				if currentNode.IsJumpable(node) {
					graph.AddEdge(currentNode, node)
				}
			}

			// down
			if currentPosition[0]+1 < len(dataMap) {
				node := dataMap[currentPosition[0]+1][currentPosition[1]]
				graph.AddVertex(node)
				if currentNode.IsJumpable(node) {
					graph.AddEdge(currentNode, node)
				}
			}

			// left
			if currentPosition[1]-1 >= 0 {
				node := dataMap[currentPosition[0]][currentPosition[1]-1]
				graph.AddVertex(node)
				if currentNode.IsJumpable(node) {
					graph.AddEdge(currentNode, node)
				}
			}

			// right
			if currentPosition[1]+1 < len(dataMap[0]) {
				node := dataMap[currentPosition[0]][currentPosition[1]+1]
				graph.AddVertex(node)
				if currentNode.IsJumpable(node) {
					graph.AddEdge(currentNode, node)
				}
			}
		}
	}

	graph.Print()
	return graph
}

//func FindPaths(dataMap [][]Node) error {
//	//currentNode, err := GetStartPosition(dataMap)
//	if err != nil {
//		fmt.Errorf(err.Error())
//	}
//
//	graph := &Graph{}
//	//graph.AddVertex(*currentNode)
//	//RecursivePathSearch(graph, dataMap, *currentNode, GetJumpableNodes(dataMap, currentNode.position))
//
//	graph.Print()
//	return nil
//}

//
//func RecursivePathSearch(graph *Graph, dataMap [][]Node, currentNode Node, jumpableNodes []Node) {
//	fmt.Println()
//	fmt.Printf("-- node position {%v} with value %s\n", currentNode.position, string(currentNode.nodeValue))
//	for len(jumpableNodes) != 0 && string(currentNode.nodeValue) != "E" &&
//		!(len(jumpableNodes) == 1 && string(jumpableNodes[0].nodeValue) == "E") {
//		for _, childNode := range jumpableNodes {
//			graph.AddVertex(childNode)
//			added, _ := graph.AddEdge(childNode, currentNode)
//			if added {
//				fmt.Printf("node position {%v} with value %s\n", childNode.position, string(childNode.nodeValue))
//				RecursivePathSearch(graph, dataMap, childNode, GetJumpableNodes(dataMap, childNode.position))
//			}
//		}
//		//fmt.Println("finished child nodes")
//	}
//	if string(currentNode.nodeValue) == "E" {
//		fmt.Println("found E!")
//	} else {
//		fmt.Println("no jumpable nodes")
//	}
//}

func LoadMap(inputFileName string) ([][]Node, error) {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("could not load file")
	}

	var dataMap [][]Node
	for row, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)

		var dataRow []Node
		for column, character := range dataLine {
			node := NewNode(character, [2]int{row, column})
			dataRow = append(dataRow, *node)
		}
		dataMap = append(dataMap, dataRow)
	}

	//PrintMap(dataMap, "nodeValue")
	return dataMap, nil
}

func GetPositionFor(dataMap [][]Node, value string) (*Node, error) {
	for row, node := range dataMap {
		for column, _ := range node {
			if string(dataMap[row][column].nodeValue) == value {
				return &dataMap[row][column], nil
			}
		}
	}
	return nil, fmt.Errorf("no value position found")
}

func GetJumpableNodes(dataMap [][]Node, currentPosition [2]int) []Node {

	currentNode := dataMap[currentPosition[0]][currentPosition[1]]

	var node Node
	var jumpableNodes []Node

	// up
	if currentPosition[0]-1 >= 0 {
		node = dataMap[currentPosition[0]-1][currentPosition[1]]
		if currentNode.IsJumpable(node) {
			jumpableNodes = append(jumpableNodes, node)
		}
	}

	// down
	if currentPosition[0]+1 < len(dataMap) {
		node = dataMap[currentPosition[0]+1][currentPosition[1]]
		if currentNode.IsJumpable(node) {
			jumpableNodes = append(jumpableNodes, node)
		}
	}

	// left
	if currentPosition[1]-1 >= 0 {

		node = dataMap[currentPosition[0]][currentPosition[1]-1]
		if currentNode.IsJumpable(node) {
			jumpableNodes = append(jumpableNodes, node)
		}
	}

	// right
	if currentPosition[1]+1 < len(dataMap[0]) {

		node = dataMap[currentPosition[0]][currentPosition[1]+1]
		if currentNode.IsJumpable(node) {
			jumpableNodes = append(jumpableNodes, node)
		}
	}

	return jumpableNodes
}

func PrintMap(dataMap [][]Node, printType string) {
	for row, node := range dataMap {
		for column, _ := range node {
			switch printType {
			case "nodeValue":
				{
					fmt.Print(string(dataMap[row][column].nodeValue))
					break
				}
			}
		}
		fmt.Println()
	}
}

type Node struct {
	position     [2]int
	nodeValue    rune
	signal       int
	up           rune
	down         rune
	left         rune
	right        rune
	isUsed       bool
	nextMovement rune
}

func (t *Node) IsJumpable(node Node) bool {

	// position OK
	if !(t.position[0]+1 == node.position[0] || t.position[1]+1 == node.position[1] ||
		t.position[0]-1 == node.position[0] || t.position[1]-1 == node.position[1]) {
		return false
	}

	// signal OK
	if string(node.nodeValue) == "E" {
		return true
	}
	if string(t.nodeValue) == "S" && string(node.nodeValue) == "S" {
		return false
	}

	if string(t.nodeValue) == "S" {
		return true
	}
	return t.signal+1 == node.signal || t.signal == node.signal
}

func NewNode(current rune, position [2]int) *Node {
	node := new(Node)
	node.nodeValue = current
	node.position = position
	if string(current) == "S" {
		node.signal = 0
	} else {
		node.signal = int(current)
	}
	return node
}
