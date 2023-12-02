package challenge

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	gameID  int
	samples []Reveal
}

type Reveal struct {
	red   int
	green int
	blue  int
}

func Star2(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		if len(dataLine) == 0 {
			continue
		}
		reveal := GetMinPossibleFromGame(dataLine)
		power := reveal.blue * reveal.red * reveal.green
		retValue = retValue + power
	}
	return retValue
}

func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		if len(dataLine) == 0 {
			continue
		}
		game := GetGame(dataLine)
		fmt.Printf("Game %v\n", game)
		possible := true
		for _, sample := range game.samples {
			if !PossibleGame(game.gameID, sample.red, sample.green, sample.blue) {
				possible = false
				break
			}
		}
		if possible {
			retValue = retValue + game.gameID
		}
		fmt.Printf("SumGameIDs %d\n", retValue)

	}
	return retValue
}

func GetGame(gameTxt string) Game {
	game := Game{}
	gameTxtTokens := strings.Split(gameTxt, ":")
	game.gameID = GetGameID(gameTxtTokens[0])
	gameSamples := strings.Split(gameTxtTokens[1], ";")
	for _, gameSample := range gameSamples {
		game.samples = append(game.samples, GetReveal(gameSample))
	}
	return game
}

func GetGameID(gameTxt string) int {
	gameTxtTokens := strings.Split(gameTxt, " ")
	matched, err := strconv.Atoi(gameTxtTokens[1])
	if err != nil {

	}
	return matched
}

func GetReveal(gameSampleTxt string) Reveal {
	gameSample := Reveal{
		red:   0,
		green: 0,
		blue:  0,
	}
	cubesSampleTxt := strings.Split(gameSampleTxt, ",")
	for _, cubeSampleTxt := range cubesSampleTxt {
		cubeSampleTxt = strings.TrimSpace(cubeSampleTxt)
		cubeSample := strings.Split(cubeSampleTxt, " ")
		if len(cubeSample) != 2 {
			return Reveal{}
		}
		cubeNumberTxt := cubeSample[0]
		cubeColor := cubeSample[1]
		cubeNumber, err := strconv.Atoi(cubeNumberTxt)
		if err != nil {
			return Reveal{}
		}
		// fmt.Printf("%s is %d %s\n", cubeSampleTxt, cubeNumber, cubeColor)

		switch cubeColor {
		case "red":
			gameSample.red = cubeNumber
		case "green":
			gameSample.green = cubeNumber
		case "blue":
			gameSample.blue = cubeNumber
		}
	}
	return gameSample
}

func GetMinPossibleFromGame(gameTxt string) Reveal {
	game := GetGame(gameTxt)
	minPossible := Reveal{}
	for _, sample := range game.samples {
		minPossible = GetMinPossibleFromReveal(sample, minPossible)
	}
	return minPossible
}

func GetMinPossibleFromReveal(reveal Reveal, currentMinReveal Reveal) Reveal {
	minPossible := Reveal{}

	if reveal.red > currentMinReveal.red {
		minPossible.red = reveal.red
	} else {
		minPossible.red = currentMinReveal.red
	}

	if reveal.green > currentMinReveal.green {
		minPossible.green = reveal.green
	} else {
		minPossible.green = currentMinReveal.green
	}

	if reveal.blue > currentMinReveal.blue {
		minPossible.blue = reveal.blue
	} else {
		minPossible.blue = currentMinReveal.blue
	}
	return minPossible
}

func PossibleGame(gameID int, red int, green int, blue int) bool {
	const MinRed = 12
	const MinGreen = 13
	const MinBlue = 14

	if red > MinRed {
		fmt.Printf("game %d is not possible\n", gameID)
		return false
	}
	if green > MinGreen {
		fmt.Printf("game %d is not possible\n", gameID)
		return false
	}
	if blue > MinBlue {
		fmt.Printf("game %d is not possible\n", gameID)
		return false
	}
	return true
}
