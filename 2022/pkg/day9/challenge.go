package day9

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	headPos := Position{x: 0, y: 0}
	headPosition := &headPos

	tailPos := Position{0, 0}
	tailPosition := &tailPos

	var tailPositions []Position
	tailPositions = append(tailPositions, tailPos)
	fmt.Printf("movement: - - head: %v - tail: %v \n", headPos, tailPos)

	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)
		movement, err := getMovement(dataLine)
		if err != nil {
			continue
		}

		move(headPosition, movement)
		tailToHeadPositions := GetTailToHeadMovements(*headPosition, *tailPosition)
		*tailPosition = tailToHeadPositions[len(tailToHeadPositions)-1]
		fmt.Printf("movement: %v  - head: %v - tail: %v \n", *movement, *headPosition, tailToHeadPositions)

		tailPositions = AddTailPositions(tailPositions, tailToHeadPositions)
	}

	//for _, tailPosition := range tailPositions {
	//	fmt.Printf("%d\n", tailPosition)
	//}
	return GetUniquePosition(tailPositions)
}

func AddTailPositions(tailPositions []Position, tailToHeadPositions []Position) []Position {
	for _, tailToHeadMov := range tailToHeadPositions {
		tailPositions = append(tailPositions, tailToHeadMov)
	}
	return tailPositions
}

func GetTailToHeadMovements(headPosition Position, tailPosition Position) []Position {
	if headPosition == tailPosition {
		return []Position{headPosition}
	}

	var tailMovements []Position

	// 1 distance, no need to move
	if (math.Abs(float64(headPosition.x)-float64(tailPosition.x)) <= 1 && math.Abs(float64(headPosition.y)-float64(tailPosition.y)) <= 1) ||
		(math.Abs(float64(tailPosition.x)-float64(headPosition.x)) <= 1 && math.Abs(float64(tailPosition.y)-float64(headPosition.y)) <= -1) {
		return []Position{tailPosition}

		// diagonal
	} else if headPosition.x != tailPosition.x && headPosition.y != tailPosition.y {
		if headPosition.x > tailPosition.x {
			// UpRight
			if headPosition.y > tailPosition.y {
				tailPosition.x += 1
				tailPosition.y += 1
				// DownRight
			} else {
				tailPosition.x += 1
				tailPosition.y -= 1
			}
		} else {
			// UpLeft
			if headPosition.y > tailPosition.y {
				tailPosition.x -= 1
				tailPosition.y += 1
				// DownLeft
			} else {
				tailPosition.x -= 1
				tailPosition.y -= 1
			}
		}
		if (math.Abs(float64(headPosition.x)-float64(tailPosition.x)) <= 1 && math.Abs(float64(headPosition.y)-float64(tailPosition.y)) <= 1) ||
			(math.Abs(float64(tailPosition.x)-float64(headPosition.x)) <= 1 && math.Abs(float64(tailPosition.y)-float64(headPosition.y)) <= -1) {
			return []Position{tailPosition}
		}
		tailMovements = append(tailMovements, tailPosition)
	}

	if headPosition.x == tailPosition.x {
		// up
		if headPosition.y > tailPosition.y {
			counter := 1
			for i := tailPosition.y; i < headPosition.y-1; i++ {
				tailMovements = append(tailMovements, Position{tailPosition.x, tailPosition.y + counter})
				counter++
			}

			// down
		} else if headPosition.y < tailPosition.y {
			counter := 1
			if headPosition.y <= 0 {
				for i := tailPosition.y; i > headPosition.y+1; i-- {
					tailMovements = append(tailMovements, Position{tailPosition.x, tailPosition.y - counter})
					counter++
				}
			} else {
				for i := tailPosition.y; i > headPosition.y-1; i-- {
					tailMovements = append(tailMovements, Position{tailPosition.x, tailPosition.y - counter})
					counter++
				}
			}
		}
	}

	if headPosition.y == tailPosition.y {
		// left
		if headPosition.x < tailPosition.x {
			counter := 1
			if headPosition.x <= 0 {
				for i := tailPosition.x; i > headPosition.x+1; i-- {
					tailMovements = append(tailMovements, Position{tailPosition.x - counter, tailPosition.y})
					counter++
				}

			} else {
				for i := tailPosition.x; i > headPosition.x+1; i-- {
					tailMovements = append(tailMovements, Position{tailPosition.x - counter, tailPosition.y})
					counter++
				}
			}

			// right
		} else if headPosition.x > tailPosition.x {
			counter := 1
			for i := tailPosition.x; i < headPosition.x-1; i++ {
				tailMovements = append(tailMovements, Position{tailPosition.x + counter, tailPosition.y})
				counter++
			}
		}
	}

	if tailMovements == nil {
		return []Position{tailPosition}
	}
	return tailMovements
}

func getMovement(dataLine string) (*Movement, error) {
	movementTxt := strings.Split(dataLine, " ")
	if len(movementTxt) != 2 {
		return nil, fmt.Errorf("movement has no correct syntax %s", dataLine)
	}

	action := movementTxt[0]
	steps, _ := strconv.Atoi(movementTxt[1])

	if (action == "U" || action == "D" || action == "L" || action == "R") && steps > 0 {
		return NewMovement(action, steps), nil
	}
	return nil, fmt.Errorf("movement has no correct syntax %s", dataLine)
}

func Move(position Position, movements []*Movement) Position {
	var newPosition = &position
	for _, movement := range movements {
		move(newPosition, movement)
	}
	return position
}

func move(currentPosition *Position, movement *Movement) {
	switch movement.direction {
	case "R":
		currentPosition.x += movement.steps
		break

	case "D":
		currentPosition.y -= movement.steps
		break

	case "U":
		currentPosition.y += movement.steps
		break

	case "L":
		currentPosition.x -= movement.steps
		break
	}
}

func GetUniquePosition(positions []Position) int {

	uniquePositions := make(map[Position]bool)

	//uniquePositions[Position{0, 0}] = true
	//exists := uniquePositions[Position{0, 0}]

	for _, position := range positions {
		if !uniquePositions[position] {
			uniquePositions[position] = true
		}
	}

	//counter := 0
	//for tailPosition, _ := range uniquePositions {
	//	fmt.Printf("%d\n", tailPosition)
	//	counter++
	//}
	//fmt.Printf("\n\n%d\n\n", counter)
	return len(uniquePositions)
}

type Position struct {
	x int
	y int
}

type Movement struct {
	direction string
	steps     int
}

func NewMovement(direction string, steps int) *Movement {
	movement := new(Movement)
	movement.direction = direction
	movement.steps = steps
	return movement
}
