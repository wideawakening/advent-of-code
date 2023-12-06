package challenge

import (
	"fmt"
	"os"
	"strings"
)

func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", dataLine)

	}
	return retValue
}

type RaceSpec struct {
	lastingTime    int
	recordDistance int
}

func GetDistanceOptions(lastingTime int) map[int]int {

	options := make(map[int]int)
	for option := 0; option <= lastingTime; option++ {
		options[option] = (lastingTime - option) * option
	}
	return options
}

func GetBeatingOptions(pressButtonOptions map[int]int, maxDistance int) int {

	beatingOptions := 0
	for _, value := range pressButtonOptions {
		if value > maxDistance {
			beatingOptions++
		}
	}
	return beatingOptions
}
