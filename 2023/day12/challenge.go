package challenge

import (
	"fmt"
	"os"
	"slices"
	"strconv"
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

type SpringStatus struct {
	rawSprings    string
	springs       []string
	springsStatus []int
}

func GetFixedLengthFromUnresolved(unresolved string) int {
	fixedValue := unresolved[strings.Index(unresolved, "#") : strings.LastIndex(unresolved, "#")+1]
	return len(fixedValue)
}

func GetPossibleArranges(input string) int {
	situation, err := GetSpringStatus(input)
	if err != nil {
		return 0
	}

	unresolvedSprings := situation.springs
	unresolvedStatus := situation.springsStatus
	for HasFixedStatus(unresolvedSprings) {

		for springGroupIdx, springsGroup := range unresolvedSprings {
			if springsGroup == "" {
				continue
			}

			// contains # and direct matching
			if strings.Contains(springsGroup, "#") {
				for statusGroupIdx, statusGroup := range unresolvedStatus {
					if GetFixedLengthFromUnresolved(springsGroup) <= statusGroup {
						fmt.Printf("there's a direct match with %s with groupLenght:%d\n", springsGroup, unresolvedStatus[statusGroupIdx])
						unresolvedSprings = slices.Delete(unresolvedSprings, springGroupIdx, springGroupIdx+1)
						unresolvedStatus = slices.Delete(unresolvedStatus, statusGroupIdx, statusGroupIdx+1)
						break
					}
				}
			}

			// contains # but NO direct matching

		}
	}

	fmt.Println("unresolvedSprings:", unresolvedSprings)
	fmt.Println("unresolvedStatus:", unresolvedStatus)

	// second scan without #
	options := 0
	for idx, springGroup := range unresolvedSprings {

		// same group number
		if len(unresolvedSprings) == len(unresolvedStatus) {
			options += ComputeCombinations(len(springGroup), unresolvedStatus[idx])
			continue
		}

		totalStatusCount := 0
		for _, statusCount := range unresolvedStatus {
			totalStatusCount += statusCount
		}

		if 3 == len(springGroup) && 3 == totalStatusCount+1 {
			options += options + 1
		}
	}

	return options
}

func ComputeCombinations(springGroupLength int, statusCount int) int {
	return springGroupLength - statusCount + 1
}

func HasFixedStatus(unresolvedSprings []string) bool {
	for _, springGroup := range unresolvedSprings {
		if strings.Contains(springGroup, "#") {
			return true
		}
	}
	return false
}

func GetSpringStatus(input string) (*SpringStatus, error) {
	var retValue SpringStatus
	statusSplitTxt := strings.Split(input, " ")
	if len(statusSplitTxt) != 2 {
		return nil, fmt.Errorf("input string is empty")
	}
	retValue.rawSprings = statusSplitTxt[0]
	for _, status := range strings.Split(statusSplitTxt[0], ".") {

		if status != "" {
			retValue.springs = append(retValue.springs, status)
		}
	}

	for _, springStatusCounter := range strings.Split(statusSplitTxt[1], ",") {
		springStatusCounter, err := strconv.Atoi(springStatusCounter)
		if err != nil {
			return nil, fmt.Errorf("counter is not numeric")
		}
		retValue.springsStatus = append(retValue.springsStatus, springStatusCounter)
	}
	return &retValue, nil
}
