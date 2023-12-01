package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
input1: required calory number
input2 (file): per elf calories they're bringing

elf1- item1_calorie
elf1- item2_calorie

elf2- item1_calorie
elf2- item2_calorie

elf3- item1_calorie

output: how many calories is the elf with most calories carrying
*/
func Star1(elfCalorieFile string) int {
	data, err := os.ReadFile(elfCalorieFile)
	if err != nil {
		fmt.Errorf("error %s", err)
		return 0
	}

	var elf = 0
	var calCount = 0
	var maxCalCount = 0
	for _, line := range strings.Split(string(data), "\n") {
		fmt.Printf("current value %s\n", line)
		if line == "" {
			if calCount > maxCalCount {
				maxCalCount = calCount
				fmt.Printf("maxCalCount modified to %d\n", maxCalCount)
			}
			fmt.Printf("elf %d - %d calorie count finished |  maxCalCount: %d | going for next\n", elf, calCount, maxCalCount)
			calCount = 0
			elf++
			continue
		}
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			fmt.Errorf("error %s", err)
			return 0
		}
		calCount += lineValue
	}
	return maxCalCount
}

func Star2(elfCalorieFile string) int {
	data, err := os.ReadFile(elfCalorieFile)
	if err != nil {
		fmt.Errorf("error %s", err)
		return 0
	}

	var elf = 0
	var calCount = 0
	var maxCalCountTopThree = [3]int{}
	for _, line := range strings.Split(string(data), "\n") {
		fmt.Printf("current value %s\n", line)
		if line == "" {
			lowestIndex, lowestValue := GetLowestValueIndex(maxCalCountTopThree)
			if calCount > lowestValue {
				maxCalCountTopThree[lowestIndex] = calCount
				fmt.Printf("maxCalCount modified to %+v\n", maxCalCountTopThree)
			}
			fmt.Printf("elf %d - %d calorie count finished |  maxCalCount: %+v | going for next\n", elf, calCount, maxCalCountTopThree)
			calCount = 0
			elf++
			continue
		}
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			fmt.Errorf("error %s", err)
			return 0
		}
		calCount += lineValue
	}
	return maxCalCountTopThree[0] + maxCalCountTopThree[1] + maxCalCountTopThree[2]
}

func GetLowestValueIndex(top [3]int) (int, int) {
	var lowestIndex = 0
	var lowestValue = top[0]
	for index, value := range top {
		if value < lowestValue {
			lowestValue = value
			lowestIndex = index
		}
	}
	return lowestIndex, lowestValue
}
