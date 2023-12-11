package challenge

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func FindStart(inputFileName string) *MapPoint {

	inputData, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	for x, dataLine := range strings.Split(string(inputData), "\n") {
		for y, data := range dataLine {
			if string(data) == "S" {
				return &MapPoint{x, y, "S"}
			}
		}
	}

	return nil
}

func GetAdjacentValidMoves(inputFileName string, currentPointX int, currentPointY int) []MapPoint {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	dataRows := strings.Split(string(data), "\n")
	currentPoint := string(dataRows[currentPointX][currentPointY])
	up := "."
	if currentPointX-1 >= 0 {
		up = string(dataRows[currentPointX-1][currentPointY])
	}

	down := "."
	if currentPointX+1 <= len(dataRows)-1 {
		down = string(dataRows[currentPointX+1][currentPointY])
	}

	left := "."
	if currentPointY-1 >= 0 {
		left = string(dataRows[currentPointX][currentPointY-1])
	}

	right := "."
	if currentPointY+1 <= len(dataRows[currentPointX])-1 {
		right = string(dataRows[currentPointX][currentPointY+1])
	}

	var validMoves []MapPoint
	if IsValidAdjacent(currentPoint, up, "Up") {
		validMoves = append(validMoves, MapPoint{currentPointX - 1, currentPointY, up})
	}
	if IsValidAdjacent(currentPoint, down, "Down") {
		validMoves = append(validMoves, MapPoint{currentPointX + 1, currentPointY, down})
	}
	if IsValidAdjacent(currentPoint, left, "Left") {
		validMoves = append(validMoves, MapPoint{currentPointX, currentPointY - 1, left})
	}
	if IsValidAdjacent(currentPoint, right, "Right") {
		validMoves = append(validMoves, MapPoint{currentPointX, currentPointY + 1, right})
	}

	return validMoves
}

type Position string

type MapPoint struct {
	x      int
	y      int
	symbol string
}

func IsValidAdjacent(currentPoint string, adjacentPoint string, adjacentPosition Position) bool {

	switch currentPoint {
	case "S":
		switch adjacentPosition {
		case "Up":
			return adjacentPoint == "|" || adjacentPoint == "F" || adjacentPoint == "7"
		case "Down":
			return adjacentPoint == "|" || adjacentPoint == "J" || adjacentPoint == "L"
		case "Right":
			return adjacentPoint == "-" || adjacentPoint == "J" || adjacentPoint == "7"
		case "Left":
			return adjacentPoint == "-" || adjacentPoint == "L" || adjacentPoint == "F"
		}
		return false
	case "|":
		switch adjacentPosition {
		case "Up":
			return adjacentPoint == "|" || adjacentPoint == "F" || adjacentPoint == "7" || adjacentPoint == "S"
		case "Down":
			return adjacentPoint == "|" || adjacentPoint == "J" || adjacentPoint == "L" || adjacentPoint == "S"
		}
		return false
	case "-":
		switch adjacentPosition {
		case "Right":
			return adjacentPoint == "-" || adjacentPoint == "J" || adjacentPoint == "7" || adjacentPoint == "S"
		case "Left":
			return adjacentPoint == "-" || adjacentPoint == "L" || adjacentPoint == "F" || adjacentPoint == "S"
		}
		return false
	case "L":
		switch adjacentPosition {
		case "Up":
			return adjacentPoint == "|" || adjacentPoint == "F" || adjacentPoint == "7" || adjacentPoint == "S"
		case "Right":
			return adjacentPoint == "-" || adjacentPoint == "J" || adjacentPoint == "7" || adjacentPoint == "S"
		}
		return false
	case "J":
		switch adjacentPosition {
		case "Up":
			return adjacentPoint == "|" || adjacentPoint == "F" || adjacentPoint == "7" || adjacentPoint == "S"
		case "Left":
			return adjacentPoint == "-" || adjacentPoint == "L" || adjacentPoint == "F" || adjacentPoint == "S"
		}
		return false
	case "7":
		switch adjacentPosition {
		case "Down":
			return adjacentPoint == "|" || adjacentPoint == "J" || adjacentPoint == "L" || adjacentPoint == "S"
		case "Left":
			return adjacentPoint == "-" || adjacentPoint == "L" || adjacentPoint == "F" || adjacentPoint == "S"
		}
		return false
	case "F":
		switch adjacentPosition {
		case "Down":
			return adjacentPoint == "|" || adjacentPoint == "J" || adjacentPoint == "L" || adjacentPoint == "S"
		case "Right":
			return adjacentPoint == "-" || adjacentPoint == "J" || adjacentPoint == "7" || adjacentPoint == "S"
		}
		return false
	case ".":
		return false

	}
	return false
}

func GetLoop(inputFileName string) []MapPoint {
	//var loops []MapPoint
	startPoint := FindStart(inputFileName)
	validMoves := GetAdjacentValidMoves(inputFileName, startPoint.x, startPoint.y)
	for _, validMove := range validMoves {
		possibleLoop := ContinueDirection(inputFileName, []MapPoint{*startPoint, validMove})

		if possibleLoop == nil {
			continue
		}

		if IsStartingPoint(*startPoint, possibleLoop[len(possibleLoop)-1]) {
			fmt.Printf("Found loop of length %d\n", len(possibleLoop))
			for _, point := range possibleLoop {
				fmt.Printf("%s", point.symbol)
			}
			fmt.Println()
			fmt.Println()
			return possibleLoop
		}
	}
	return nil
}

func ContinueDirection(inputFileName string, currentPath []MapPoint) []MapPoint {
	startPoint := currentPath[0]
	currentPoint := currentPath[len(currentPath)-1]

	if IsStartingPoint(startPoint, currentPoint) {
		return currentPath
	}

	validMoves := GetAdjacentValidMoves(inputFileName, currentPoint.x, currentPoint.y)
	for _, validMove := range validMoves {

		// avoid going back track
		if IsPreviousPoint(currentPath, validMove) {
			continue
		}

		currentPath = append(currentPath, validMove)
		return ContinueDirection(inputFileName, currentPath)
	}

	return currentPath
}

func IsStartingPoint(startPoint MapPoint, currentPoint MapPoint) bool {
	return startPoint.x == currentPoint.x && startPoint.y == currentPoint.y
}

func IsPreviousPoint(currentPath []MapPoint, currentPoint MapPoint) bool {
	previousPoint := currentPath[len(currentPath)-2]
	return previousPoint.x == currentPoint.x && previousPoint.y == currentPoint.y
}

func FindInsideObjectsWithinLoop(inputFileName string) int {
	loopPoints := GetLoop(inputFileName)

	inputData, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	objectsWithinLoop := 0

	for x := 0; x < len(strings.Split(string(inputData), "\n"))-1; x++ {
		dataLine := strings.Split(string(inputData), "\n")[x]
		for y := 0; y < len(dataLine)-1; y++ {
			data := dataLine[y]
			firstPoint := GetNearestLoopPoint(MapPoint{x, y, string(data)}, loopPoints)
			if firstPoint == nil {
				break
			}

			possiblePipeEnd := GetEndOfPipe(*firstPoint, loopPoints)
			if possiblePipeEnd != nil {
				y = possiblePipeEnd.y + 1
				continue
			}

			secondPoint := GetNextLoopPoint(*firstPoint, loopPoints)
			if secondPoint == nil {
				break
			}

			// skip adjacent points
			if firstPoint.y+1 == secondPoint.y {
				y = secondPoint.y + 1
				continue
			}

			// TODO check if data is in between firstPoint and secondPoint
			addObjects := secondPoint.y - (firstPoint.y + 1)
			objectsWithinLoop += addObjects
			fmt.Printf("computing %d objects, from %v to %v\n", addObjects, firstPoint, secondPoint)

			y = secondPoint.y + 1
		}
	}

	return objectsWithinLoop
}

func GetNearestLoopPoint(currentPoint MapPoint, loopPoints []MapPoint) *MapPoint {
	for _, loopPoint := range loopPoints {
		if loopPoint.x != currentPoint.x {
			continue
		}
		if loopPoint.y > currentPoint.y {
			return &loopPoint
		}
	}
	return nil
}

func GetLastLoopPoint(currentPoint MapPoint, loopPoints []MapPoint) *MapPoint {

	farestPoint := MapPoint{currentPoint.x, currentPoint.y, currentPoint.symbol}
	for _, loopPoint := range loopPoints {
		if loopPoint.x != currentPoint.x {
			continue
		}
		if loopPoint.y > farestPoint.y {
			farestPoint = MapPoint{loopPoint.x, loopPoint.y, loopPoint.symbol}
		}
	}

	if farestPoint.x == currentPoint.x && farestPoint.y == currentPoint.y {
		return nil
	}
	return &farestPoint
}

func GetNextLoopPoint(point MapPoint, loopPoints []MapPoint) *MapPoint {

	nearestY := MapPoint{0, math.MaxInt, "0"}
	for _, possibleNearestPoint := range loopPoints {
		if possibleNearestPoint.x != point.x {
			continue
		}
		if possibleNearestPoint.y > point.y && possibleNearestPoint.y <= nearestY.y {
			nearestY = MapPoint{possibleNearestPoint.x, possibleNearestPoint.y, possibleNearestPoint.symbol}
		}
	}

	if nearestY.x == point.x && nearestY.y == point.y || nearestY.y == math.MaxInt {
		return nil
	}
	return &nearestY
}

func GetEndOfPipe(currentPoint MapPoint, loopPoints []MapPoint) *MapPoint {

	var lastPipePoint *MapPoint
	farestPoint := GetLastLoopPoint(currentPoint, loopPoints)
	if farestPoint == nil {
		return nil
	}

	nextLoopPoint := GetNextLoopPoint(currentPoint, loopPoints)
	if nextLoopPoint == nil {
		return nil
	}
	for nextLoopPoint.symbol == "-" {
		lastPipePoint = &MapPoint{nextLoopPoint.x, nextLoopPoint.y, nextLoopPoint.symbol}
		nextLoopPoint = GetNextLoopPoint(*lastPipePoint, loopPoints)
		if nextLoopPoint == nil {
			break
		}
	}

	if lastPipePoint != nil {
		lastPipePoint = GetNextLoopPoint(*lastPipePoint, loopPoints)
		fmt.Printf("Found pipe from %v to %v\n", currentPoint, lastPipePoint)
		return lastPipePoint
	}
	return nil
}
