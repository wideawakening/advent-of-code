package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EvaluateStar(inputFileName string, start1 bool) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		return 0
	}

	treeHeights := [][]int{}
	for _, dataLine := range strings.Split(string(data), "\n") {
		dataX := []int{}
		//fmt.Printf("%s", dataLine)
		if dataLine == "" {
			continue
		}
		for j := 0; j < len(dataLine); j++ {
			num, err := strconv.Atoi(string(dataLine[j]))
			if err != nil {
				fmt.Printf("error")
				return 0
			}
			dataX = append(dataX, num)
		}
		treeHeights = append(treeHeights, dataX)
	}

	if start1 {
		return GetNumberVisibleTress(treeHeights)
	}

	return GetMaxScenicScore(treeHeights)

}

func GetMaxScenicScore(treeHeights [][]int) int {
	maxScore := 0

	for i := 0; i < len(treeHeights[0]); i++ {
		for j := 0; j < len(treeHeights); j++ {
			treeScore := DetermineScenicScore(treeHeights, i, j)
			if treeScore > maxScore {
				maxScore = treeScore
			}
		}
	}
	return maxScore
}

func GetNumberVisibleTress(treeHeights [][]int) int {
	numVisibleTrees := 0

	for i := 0; i < len(treeHeights[0]); i++ {
		for j := 0; j < len(treeHeights); j++ {
			if CheckVisibleTree(treeHeights, i, j) {
				numVisibleTrees++
			}
		}
	}
	return numVisibleTrees
}

func CheckVisibleTree(treeHeights [][]int, posX int, posY int) bool {

	TreeValue := treeHeights[posY][posX]
	MaxX := len(treeHeights) - 1
	MaxY := len(treeHeights[0]) - 1

	// north
	visibleNorth := true
	for i := 0; i < posY; i++ {
		if treeHeights[i][posX] >= TreeValue {
			visibleNorth = false
			break
		}
	}

	// south
	VisibleSouth := true
	for i := posY + 1; i <= MaxY; i++ {
		if treeHeights[i][posX] >= TreeValue {
			VisibleSouth = false
			break
		}
	}

	// east
	VisibleEast := true
	for i := posX + 1; i <= MaxX; i++ {
		if treeHeights[posY][i] >= TreeValue {
			VisibleEast = false
			break
		}
	}

	// west
	VisibleWest := true
	for i := 0; i < posX; i++ {
		if treeHeights[posY][i] >= TreeValue {
			VisibleWest = false
			break
		}
	}

	return visibleNorth || VisibleSouth || VisibleWest || VisibleEast
}

func DetermineScenicScore(treeHeights [][]int, posX int, posY int) int {
	TreeValue := treeHeights[posY][posX]
	MaxX := len(treeHeights) - 1
	MaxY := len(treeHeights[0]) - 1

	// north
	northScore := 0
	if posY != 0 {
		for i := posY - 1; i >= 0; i-- {
			if treeHeights[i][posX] >= TreeValue {
				northScore++
				break
			}
			northScore++
		}
	}

	// south
	southScore := 0
	if posY != MaxY {
		for i := posY + 1; i <= MaxY; i++ {
			if treeHeights[i][posX] >= TreeValue {
				southScore++
				break
			}
			southScore++
		}
	}

	// east
	eastScore := 0
	if posX != MaxX {
		for i := posX + 1; i <= MaxX; i++ {
			if treeHeights[posY][i] >= TreeValue {
				eastScore++
				break
			}
			eastScore++
		}
	}

	// west
	westScore := 0
	if posX != 0 {
		for i := posX - 1; i >= 0; i-- {
			if treeHeights[posY][i] >= TreeValue {
				westScore++
				break
			}
			westScore++
		}
	}

	return northScore * southScore * eastScore * westScore
}
