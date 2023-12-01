package day6

import (
	"fmt"
	"os"
	"strings"
)

/*
TUNING TROUBLE
- start-of-packet (SOP) marker: 4/14 (sopSize) characters all different
- return the index after the SOP
*/
func Star1(inputFileName string, sopSize int) (int, string) {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0, ""
	}

	for _, dataLine := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", dataLine)
		return FindSOPIndex(dataLine, sopSize)
	}
	return 0, ""
}

func FindSOPIndex(transmission string, sopSize int) (int, string) {
	for i := 0; i+sopSize < len(transmission); i++ {
		isSOP, SOP := CheckIsSOP(i, transmission, sopSize)
		if isSOP {
			return strings.Index(transmission, SOP) + sopSize, SOP
		}
	}
	return 0, ""
}

func CheckIsSOP(startIndex int, transmission string, sopSize int) (bool, string) {
	if startIndex+sopSize > len(transmission) {
		return false, ""
	}
	sopPackage := ""
	sopPackage += string(transmission[startIndex])

	var nextCharacter string

	for i := startIndex + 1; i < startIndex+sopSize; i++ {
		nextCharacter = string(transmission[i])
		if strings.Contains(sopPackage, string(nextCharacter)) {
			return false, ""
		}
		sopPackage += nextCharacter
	}
	fmt.Println(sopPackage)
	return true, transmission[startIndex : startIndex+sopSize]
}
