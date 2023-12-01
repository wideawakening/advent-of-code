package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
every section uniqueID
elf assgned to A range of section IDs

get full assignment overlap

NO full overlap
2-4
6-8

YES full overlpa
6-6
4-6
*/
func Star1(inputFileName string, fullOverlap bool) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		roundData := strings.Split(dataLine, ",")

		if len(roundData) == 2 {
			if fullOverlap {
				if GetFullyCovers(roundData[0], roundData[1]) {
					retValue += 1
				}
			} else {
				if GetCovers(roundData[0], roundData[1]) {
					retValue += 1
				}
			}

		}
	}
	return retValue
}

func GetPairNumbers(pair string) (int, int) {
	pairMinMax := strings.Split(pair, "-")
	pairMin, _ := strconv.Atoi(pairMinMax[0])
	pairMax, _ := strconv.Atoi(pairMinMax[1])
	return pairMin, pairMax
}

func GetFullyCovers(pair1 string, pair2 string) bool {
	pair1Min, pair1Max := GetPairNumbers(pair1)
	pair2Min, pair2Max := GetPairNumbers(pair2)

	if pair1Min <= pair2Min && pair1Max >= pair2Max {
		return true
	} else if pair2Min <= pair1Min && pair2Max >= pair1Max {
		return true
	}
	return false
}

func GetCovers(pair1 string, pair2 string) bool {
	pair1Min, pair1Max := GetPairNumbers(pair1)
	pair2Min, pair2Max := GetPairNumbers(pair2)

	if pair1Min >= pair2Min {
		return pair1Min <= pair2Max
	}

	return pair1Max >= pair2Min
}
