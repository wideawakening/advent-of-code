package challenge

import (
	"fmt"
	"os"
	"strings"
)

func Star2(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		queue := make([]int, 0, 2)
		fmt.Printf("%s\n", dataLine)
		for i, char := range dataLine {

			if char >= '0' && char <= '9' {
				fmt.Printf("%c\n", char)
				queue = append(queue, int(char-'0'))

			} else {
				numberWord := TryNumberWord(dataLine[i:])
				if numberWord != 0 {
					fmt.Printf("numberWord: %d\n", numberWord)
					queue = append(queue, numberWord)
				}
			}
		}

		fmt.Printf("queue %v\n", queue)
		if len(queue) < 1 {
			continue
		}
		first := queue[0]
		last := queue[len(queue)-1]
		fmt.Printf("%d%d\n", first, last)

		digit := first*10 + last
		fmt.Printf("calibration: %d\n", digit)
		retValue += digit
	}
	return retValue
}

func TryNumberWord(remainingTxt string) int {
	if strings.HasPrefix(remainingTxt, "one") {
		return 1
	}
	if strings.HasPrefix(remainingTxt, "two") {
		return 2
	}
	if strings.HasPrefix(remainingTxt, "three") {
		return 3
	}
	if strings.HasPrefix(remainingTxt, "four") {
		return 4
	}
	if strings.HasPrefix(remainingTxt, "five") {
		return 5
	}
	if strings.HasPrefix(remainingTxt, "six") {
		return 6
	}
	if strings.HasPrefix(remainingTxt, "seven") {
		return 7
	}
	if strings.HasPrefix(remainingTxt, "eight") {
		return 8
	}
	if strings.HasPrefix(remainingTxt, "nine") {
		return 9
	}
	return 0
}

func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		queue := make([]int, 0, 2)
		fmt.Printf("%s\n", dataLine)
		for _, char := range dataLine {
			// check if is digit
			if char >= '0' && char <= '9' {
				fmt.Printf("%c\n", char)
				queue = append(queue, int(char-'0'))
			}
		}
		fmt.Printf("queue %v\n", queue)
		first := queue[0]
		last := queue[len(queue)-1]
		fmt.Printf("%d%d\n", first, last)
		digit := first*10 + last
		fmt.Printf("calibration: %d\n", digit)
		// create a number from two digits
		retValue += digit
	}
	return retValue
}
