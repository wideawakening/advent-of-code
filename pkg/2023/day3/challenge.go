package challenge

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumberAndPositions struct {
	number    int
	positions []Position
}

type Position struct {
	row    int
	column int
}

func Star2(inputFileName string) int {
	retValue := 0

	numbersAndPositions := GetNumbersAndPositions(inputFileName)
	gearSymbol := GetGearSymbol(inputFileName)

	alreadyCalculated := map[Position]Position{}

	for _, numberAndPositions := range numbersAndPositions {
		if areAdjacent, ratePosition := AreAdjacent(gearSymbol, numberAndPositions.positions); areAdjacent {
			//fmt.Printf("number %d is adjacent to rateSymbol\n", numberAndPositions.number)

			for _, numberAndPositions2 := range numbersAndPositions {

				if _, exists := alreadyCalculated[Position{
					row:    numberAndPositions.positions[0].row,
					column: numberAndPositions.positions[0].column,
				}]; exists {
					continue
				}

				if _, exists := alreadyCalculated[Position{
					row:    numberAndPositions2.positions[0].row,
					column: numberAndPositions2.positions[0].column,
				}]; exists {
					continue
				}

				if
				// different number
				(numberAndPositions.positions[0].column != numberAndPositions2.positions[0].column ||
					numberAndPositions.positions[0].row != numberAndPositions2.positions[0].row) && (

				// num*num
				numberAndPositions.positions[0].row == ratePosition.row && ratePosition.row == numberAndPositions2.positions[0].row ||

					//     *
					// num  num
					numberAndPositions.positions[0].row-1 == ratePosition.row && ratePosition.row+1 == numberAndPositions2.positions[0].row ||

					// num  num
					//    *
					numberAndPositions.positions[0].row+1 == ratePosition.row && ratePosition.row-1 == numberAndPositions2.positions[0].row ||

					// num
					// *num
					numberAndPositions.positions[0].row+1 == ratePosition.row && ratePosition.row == numberAndPositions2.positions[0].row ||

					// *num
					// num
					numberAndPositions.positions[0].row == ratePosition.row && ratePosition.row+1 == numberAndPositions2.positions[0].row ||

					// num
					// *
					// num
					numberAndPositions.positions[0].row+1 == ratePosition.row && ratePosition.row+1 == numberAndPositions2.positions[0].row) {

					if areAdjacent2, _ := AreAdjacent([]Position{*ratePosition}, numberAndPositions2.positions); areAdjacent2 {
						fmt.Printf("%d and %d are adjacent to gearSymbol\n", numberAndPositions.number, numberAndPositions2.number)

						alreadyCalculated[Position{
							row:    numberAndPositions.positions[0].row,
							column: numberAndPositions.positions[0].column,
						}] = Position{
							row:    numberAndPositions2.positions[0].row,
							column: numberAndPositions2.positions[0].column,
						}

						alreadyCalculated[Position{
							row:    numberAndPositions2.positions[0].row,
							column: numberAndPositions2.positions[0].column,
						}] = Position{
							row:    numberAndPositions.positions[0].row,
							column: numberAndPositions.positions[0].column,
						}

						gearRation := numberAndPositions.number * numberAndPositions2.number
						//fmt.Printf("+%d =", gearRation)
						retValue += gearRation
						//fmt.Printf("%d\n", retValue)
					}
				}
			}
		}
	}

	return retValue
}

func Star1(inputFileName string) int {

	retValue := 0

	numbersAndPositions := GetNumbersAndPositions(inputFileName)
	symbolsPositions := GetSymbolsPosition(inputFileName)

	for _, numberAndPositions := range numbersAndPositions {
		if areAdjacent, _ := AreAdjacent(symbolsPositions, numberAndPositions.positions); areAdjacent {
			fmt.Printf("number %d is adjacent to symbol\n", numberAndPositions.number)
			retValue += numberAndPositions.number
		}
	}

	return retValue
}

func AreAdjacent(symbolsPositions []Position, numberPositions []Position) (bool, *Position) {
	for _, numberPosition := range numberPositions {
		if isAdjacent, position := IsPositionAdjacentToSymbol(symbolsPositions, numberPosition); isAdjacent {
			return true, position
		}
	}
	return false, nil
}

func IsPositionAdjacentToSymbol(symbolsPositions []Position, numberPosition Position) (bool, *Position) {
	for _, symbolPosition := range symbolsPositions {

		if (symbolPosition.row == numberPosition.row || symbolPosition.row-1 == numberPosition.row || symbolPosition.row+1 == numberPosition.row) &&
			(symbolPosition.column == numberPosition.column || symbolPosition.column-1 == numberPosition.column || symbolPosition.column+1 == numberPosition.column) {
			return true, &symbolPosition
		}
	}
	return false, nil
}

func GetNumbersAndPositions(inputFileName string) []NumberAndPositions {
	var numbersAndPositions []NumberAndPositions

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return numbersAndPositions
	}

	for row, dataLine := range strings.Split(string(data), "\n") {
		var possibleCombinedNumber string
		var possibleCombinedNumberPositions []Position

		column := 0
		for column < len(dataLine) {
			char := rune(dataLine[column])
			// fmt.Printf("[%d,%d] is %s\n", row, column, string(char))

			if IsRuneNumber(char) {
				possibleCombinedNumber += string(char)
				possibleCombinedNumberPositions = append(possibleCombinedNumberPositions, Position{row: row, column: column})

			} else {
				combinedNumber := GetNumber(possibleCombinedNumber)
				// we have a combinedNumber
				if combinedNumber != -1 {
					//fmt.Printf("[%d,%d] is combinedNumber %d\n", row, column, combinedNumber)
					numbersAndPositions = append(numbersAndPositions, NumberAndPositions{
						number:    combinedNumber,
						positions: possibleCombinedNumberPositions,
					})

					// reset
					possibleCombinedNumber = ""
					possibleCombinedNumberPositions = []Position{}
				}
			}
			column++
		}
	}
	return numbersAndPositions
}

func GetGearSymbol(inputFileName string) []Position {
	var positions []Position

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return positions
	}

	for row, dataLine := range strings.Split(string(data), "\n") {
		for column, char := range dataLine {
			if char == '*' {
				positions = append(positions, Position{row: row, column: column})
			}
		}
	}
	return positions
}

func GetSymbolsPosition(inputFileName string) []Position {
	var positions []Position

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return positions
	}

	for row, dataLine := range strings.Split(string(data), "\n") {
		for column, char := range dataLine {
			if IsSymbol(char) {
				positions = append(positions, Position{row: row, column: column})
			}
		}
	}
	return positions
}

func GetNumber(txt string) int {
	if number, error := strconv.Atoi(string(txt)); error == nil {
		return number
	}
	return -1
}

func IsRuneNumber(char rune) bool {
	if _, error := strconv.Atoi(string(char)); error == nil {
		return true
	}
	return false
}

func IsSymbol(char rune) bool {
	if char == '.' {
		return false
	}
	if IsRuneNumber(char) {
		return false
	}
	return true
}
