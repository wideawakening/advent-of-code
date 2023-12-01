package day2

import (
	"fmt"
	"os"
	"strings"
)

/*
							shape score
	  	A/X - rock			- 1 pts
		B/Y - paper			- 2 pts
		C/Z - scissors		- 3 pts

		round outcome score
		0 - lost
		3 - draw
		6  - win

		round score: shape score + round outcome score

input: strategy guide

	A Y
	B X
	C Z

output: total score following strategy
*/

const PIEDRA1 = "A"
const PAPEL1 = "B"
const TIJERA1 = "C"

const PIEDRA2 = "X"
const PAPEL2 = "Y"
const TIJERA2 = "Z"

func Star1(strategyGuide string, byShape bool) int {

	data, err := os.ReadFile(strategyGuide)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var totalPoints = 0
	for _, roundStrategy := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", roundStrategy)

		shape := strings.Split(roundStrategy, " ")
		if len(shape) != 2 {
			fmt.Print("error | shapes on the round were more than 2")
			continue
		}

		if byShape {
			totalPoints += EvaluatePointsByShape(shape[0], shape[1])
		} else {
			totalPoints += EvaluatePointsByAction(shape[0], shape[1])
		}

	}

	return totalPoints
}

func EvaluatePointsByShape(shape1 string, shape2 string) int {
	points := GetShapePoints(shape2) + EvaluateRound(shape1, shape2)
	return points
}

func EvaluatePointsByAction(shape string, action string) int {
	points := GetActionPoints(action) + DetermineShapePointsByAction(shape, action)
	return points
}

/*
X = loose
Y = draw
Z = win
*/
func GetActionPoints(action string) int {
	switch action {
	case PIEDRA2:
		return 0
	case PAPEL2:
		return 3
	case TIJERA2:
		return 6
	}
	return 0
}

/*
							shape score
	  	A/X - rock			- 1 pts
		B/Y - paper			- 2 pts
		C/Z - scissors		- 3 pts
*/
func GetShapePoints(shape string) int {
	switch shape {
	case PIEDRA1:
		return 1
	case PIEDRA2:
		return 1

	case PAPEL1:
		return 2
	case PAPEL2:
		return 2

	case TIJERA1:
		return 3
	case TIJERA2:
		return 3

	}
	return 0
}

/*
round outcome score
0 - lost
3 - draw
6  - win

			win/loose	- 6 pts
			C - X		piedra - tijera
			A - Y		papel	- piedra
			B - Z		tijera - papel

	draws	-	3 pts
	A - X
	B - Y
	C - Z
*/
func EvaluateRound(shape1 string, shape2 string) int {
	if shape1 == PIEDRA1 && shape2 == PIEDRA2 ||
		shape1 == PAPEL1 && shape2 == PAPEL2 ||
		shape1 == TIJERA1 && shape2 == TIJERA2 {
		return 3
	}

	if shape1 == TIJERA1 && shape2 == PIEDRA2 ||
		shape1 == PIEDRA1 && shape2 == PAPEL2 ||
		shape1 == PAPEL1 && shape2 == TIJERA2 {
		return 6
	}

	return 0
}

/*
X = loose
Y = draw
Z = win
*/
func DetermineShapePointsByAction(shape string, action string) int {
	switch action {
	case PAPEL2:
		return GetShapePoints(shape)
	case TIJERA2:
		return GetWinningShapePoints(shape)
	case PIEDRA2:
		return GetLoosingShapePoints(shape)
	}
	return 0
}

/*
A = piedra > Papel (Y)
B = papel >Tijeras (Z)
C = tijera > Piedra (X)
*/
func GetWinningShapePoints(shape string) int {
	switch shape {
	case PIEDRA1:
		return GetShapePoints(PAPEL2)
	case PAPEL1:
		return GetShapePoints(TIJERA2)
	case TIJERA1:
		return GetShapePoints(PIEDRA2)
	}
	return 0
}

/*
A = piedra > Tijeras(Z)
B = papel > Piedra(X)
C = tijera > Papel(Y)
*/
func GetLoosingShapePoints(shape string) int {
	switch shape {
	case PIEDRA1:
		return GetShapePoints(TIJERA2)
	case PAPEL1:
		return GetShapePoints(PIEDRA2)
	case TIJERA1:
		return GetShapePoints(PAPEL2)
	}
	return 0
}
