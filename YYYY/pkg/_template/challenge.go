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
