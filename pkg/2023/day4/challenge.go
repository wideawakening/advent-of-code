package challenge

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func Star2(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	retValue := 0
	nextMatchingCardsPerGame := make(map[int]int)

	for gameNumber, dataLine := range strings.Split(string(data), "\n") {
		gameNumber = gameNumber + 1
		// fmt.Printf("%s", dataLine)

		// compute current game
		matchingCardsPerGame := GetMatchingNumberPerGame(dataLine)
		retValue++

		// from current card, set next game cards
		for i := 1; i < matchingCardsPerGame+1; i++ {
			nextMatchingCardsPerGame[gameNumber+i] = nextMatchingCardsPerGame[gameNumber+i] + 1
		}

		// from copy-card, set next game cards
		currentCardCopies := nextMatchingCardsPerGame[gameNumber]
		for i := 0; i < currentCardCopies; i++ {
			for i := 1; i < matchingCardsPerGame+1; i++ {
				nextMatchingCardsPerGame[gameNumber+i] = nextMatchingCardsPerGame[gameNumber+i] + 1
			}
		}
	}

	for _, value := range nextMatchingCardsPerGame {
		retValue += value
	}
	return retValue
}

func GetMatchingNumberPerGame(scratch string) int {

	game := strings.Split(scratch, ":")
	if len(game) != 2 {
		fmt.Printf("Invalid scratchcard: %s\n", scratch)
		return 0
	}

	scratchNumbers := strings.Split(game[1], "|")
	if len(scratchNumbers) != 2 {
		fmt.Printf("Invalid scratchcard: %s\n", scratch)
		return 0
	}
	winingNumbers := strings.Split(scratchNumbers[0], " ")
	boardNumbers := strings.Split(scratchNumbers[1], " ")

	matchingNumbers := 0
	for _, myNumber := range boardNumbers {

		if myNumber == "" {
			continue
		}

		if slices.Contains(winingNumbers, myNumber) {
			matchingNumbers++
		}
	}

	fmt.Printf("game %s, matchingNumbers: %d\n", game[0], matchingNumbers)
	return matchingNumbers
}

func Star1(inputFileName string) float64 {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = float64(0)
	for _, dataLine := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", dataLine)
		retValue += EvaluateScratchStar1(dataLine)

	}
	return retValue
}

func EvaluateScratchStar1(scratch string) float64 {

	game := strings.Split(scratch, ":")
	if len(game) != 2 {
		fmt.Printf("Invalid scratchcard: %s\n", scratch)
		return float64(0)
	}

	scratchNumbers := strings.Split(game[1], "|")
	if len(scratchNumbers) != 2 {
		fmt.Printf("Invalid scratchcard: %s\n", scratch)
		return float64(0)
	}
	winingNumbers := strings.Split(scratchNumbers[0], " ")
	boardNumbers := strings.Split(scratchNumbers[1], " ")

	matchingNumbers := 0
	for _, myNumber := range boardNumbers {

		if myNumber == "" {
			continue
		}

		if slices.Contains(winingNumbers, myNumber) {
			matchingNumbers++
		}
	}

	fmt.Printf("matchingNumbers: %d\n", matchingNumbers)
	if matchingNumbers == 0 {
		return float64(0)
	}
	return math.Pow(float64(2), float64(matchingNumbers-1))
}
