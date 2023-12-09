package challenge

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Star(inputFileName string, isStar2 bool) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	sequences := make([][]int, 0)
	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)

		digitsTxt := strings.Split(dataLine, " ")
		currentSequenceValue := []int{}

		for _, digitTxt := range digitsTxt {
			digit, err := strconv.Atoi(digitTxt)
			if err != nil {
				fmt.Printf("Error parsing digit %s", digitTxt)
				continue
			}
			currentSequenceValue = append(currentSequenceValue, digit)
		}
		sequences = append(sequences, currentSequenceValue)
	}

	var retValue = 0
	for _, sequence := range sequences {
		treeSequence := GenerateNextSequenceTree([][]int{sequence})
		var interpolatedSequence [][]int
		if isStar2 {
			interpolatedSequence = InterpolatePreviousValues(treeSequence)
			retValue += interpolatedSequence[0][0]
		} else {
			interpolatedSequence = InterpolateNextValues(treeSequence)
			retValue += interpolatedSequence[0][len(interpolatedSequence[0])-1]
		}
	}
	return retValue
}

func GenerateNextSequenceTree(sequences [][]int) [][]int {

	currentSequence := sequences[len(sequences)-1]
	var nextSequence []int

	for i := 0; i < len(currentSequence)-1; i++ {
		currentDigit := currentSequence[i]
		nextDigit := currentSequence[i+1]
		nextSequence = append(nextSequence, nextDigit-currentDigit)
	}

	sequences = append(sequences, nextSequence)

	if IsEnd(nextSequence) {
		return sequences
	}
	return GenerateNextSequenceTree(sequences)
}

func IsEnd(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			return false
		}
	}
	return true
}

func InterpolateNextValues(sequences [][]int) [][]int {

	for i := len(sequences) - 1; i >= 0; i-- {
		currentSequence := sequences[i]
		if IsEnd(currentSequence) {
			sequences[i] = append(currentSequence, 0)
			continue
		}
		nextSequence := sequences[i+1]
		sequences[i] = append(currentSequence, ComputeNextValue(currentSequence[len(currentSequence)-1], nextSequence[len(nextSequence)-1]))
	}
	return sequences
}

func InterpolatePreviousValues(sequences [][]int) [][]int {

	for i := len(sequences) - 1; i >= 0; i-- {
		currentSequence := sequences[i]
		if IsEnd(currentSequence) {
			sequences[i] = append([]int{0}, currentSequence...)
			continue
		}
		nextSequence := sequences[i+1]
		sequences[i] = append([]int{ComputePreviousValue(currentSequence[0], nextSequence[0])}, currentSequence...)
	}
	return sequences
}

func ComputeNextValue(currentSequenceLastValue int, nextSequenceLastValue int) int {
	return nextSequenceLastValue + currentSequenceLastValue
}

func ComputePreviousValue(currentSequenceFirstValue int, nextSequenceFirstValue int) int {
	return currentSequenceFirstValue - nextSequenceFirstValue
}
