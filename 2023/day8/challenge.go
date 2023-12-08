package challenge

import (
	"fmt"
	"strings"
	"sync"
)

type Instruction string

const GoRight Instruction = "R"
const GoLeft Instruction = "L"

const RightNode = 1
const LeftNode = 0

type Map struct {
	Instructions []Instruction
	Nodes        map[string][2]string
}

func GetElfMap(mapContent string) Map {
	elfMap := Map{Instructions: []Instruction{}, Nodes: make(map[string][2]string)}

	for lineNumber, dataLine := range strings.Split(string(mapContent), "\n") {
		//fmt.Printf("%s\n", dataLine)

		if lineNumber == 0 {
			for _, instruction := range dataLine {
				elfMap.Instructions = append(elfMap.Instructions, Instruction(string(instruction)))
			}

		} else {
			if len(strings.TrimSpace(dataLine)) == 0 {
				continue
			}

			nodeDataTxt := strings.Split(dataLine, "=")
			if nodeDataTxt == nil || len(nodeDataTxt) != 2 {
				continue
			}

			nodeInstructionsTxt := strings.Split(nodeDataTxt[1], ",")
			if nodeDataTxt == nil || len(nodeDataTxt) != 2 {
				continue
			}

			nodeName := strings.TrimSpace(nodeDataTxt[0])
			instructionLeft := strings.TrimPrefix(nodeInstructionsTxt[0], " (")
			instructionRight := strings.TrimSpace(strings.TrimSuffix(nodeInstructionsTxt[1], ")"))

			elfMap.Nodes[nodeName] = [2]string{instructionLeft, instructionRight}
		}
	}

	return elfMap
}

func GetStepsToExit(elfMap Map) int {
	steps := 0

	numInstructions := len(elfMap.Instructions)
	currentPos := "AAA"

	for ; currentPos != "ZZZ"; steps++ {
		instruction := elfMap.Instructions[steps%numInstructions]
		//fmt.Printf("%s\n", instruction)
		if instruction == GoRight {
			currentPos = elfMap.Nodes[currentPos][RightNode]
		} else {
			currentPos = elfMap.Nodes[currentPos][LeftNode]
		}
	}
	return steps
}

func GetStartNodes(elfMap Map) []string {
	var startNodes []string
	for nodeName, _ := range elfMap.Nodes {
		if nodeName[2] == 'A' {
			startNodes = append(startNodes, nodeName)
		}
	}
	return startNodes
}

func InneficientGetGhostStepsToExit(elfMap Map) int {
	steps := 0

	numInstructions := len(elfMap.Instructions)
	allCurrentPositions := GetStartNodes(elfMap)
	fmt.Printf("current positions: %d\n", len(allCurrentPositions))

	startpointWaitGroup := sync.WaitGroup{}
	nodeWaitGroup := sync.WaitGroup{}

	//PrintAllCurrentPositions(allCurrentPositions)
	for ; !AllNodesEnded(allCurrentPositions); steps++ {
		steps := steps
		startpointWaitGroup.Add(1)
		go func() {
			defer startpointWaitGroup.Done()

			for i, currentPosition := range allCurrentPositions {
				//fmt.Printf("evaluating %d %s\n", i, currentPosition)
				i := i
				steps := steps
				currentPosition := currentPosition

				nodeWaitGroup.Add(1)
				go func() {
					defer nodeWaitGroup.Done()

					instruction := elfMap.Instructions[steps%numInstructions]
					//fmt.Printf("%s\n", instruction)
					if instruction == GoRight {
						allCurrentPositions[i] = elfMap.Nodes[currentPosition][RightNode]
					} else if instruction == GoLeft {
						allCurrentPositions[i] = elfMap.Nodes[currentPosition][LeftNode]
					} else {
						fmt.Printf(" fuckers")
					}
				}()
				nodeWaitGroup.Wait()
			}
		}()
		startpointWaitGroup.Wait()
		//PrintAllCurrentPositions(allCurrentPositions)
		fmt.Printf("%d\n", steps)
	}
	return steps
}

func GetGhostStepsToExit(elfMap Map) int {
	numInstructions := len(elfMap.Instructions)
	allCurrentPositions := GetStartNodes(elfMap)
	fmt.Printf("current positions: %d\n", len(allCurrentPositions))

	reachedEndOnStep := make(map[string][]int)
	for i, currentPos := range allCurrentPositions {

		steps := 0 // TODO double?
		for ; !NodeEnded(currentPos); steps++ {
			instruction := elfMap.Instructions[steps%numInstructions]
			//fmt.Printf("%s\n", instruction)
			if instruction == GoRight {
				currentPos = elfMap.Nodes[currentPos][RightNode]
			} else {
				currentPos = elfMap.Nodes[currentPos][LeftNode]
			}
		}
		if NodeEnded(currentPos) {
			fmt.Printf("%d node reached end %s on step %d\n", i, currentPos, steps)
			reachedEndOnStep[currentPos] = append(reachedEndOnStep[currentPos], steps)
		}
	}
	return 0
}

func AllNodesEnded(nodePositions []string) bool {
	for _, nodePosition := range nodePositions {
		if !NodeEnded(nodePosition) {
			return false
		}
	}
	return true
}

func NodeEnded(nodeName string) bool {
	return string(nodeName[2]) == "Z"
}

func PrintAllCurrentPositions(allCurrentPositions []string) {
	for _, currentPosition := range allCurrentPositions {
		fmt.Printf("%s ", currentPosition)
	}
	fmt.Println()
}
