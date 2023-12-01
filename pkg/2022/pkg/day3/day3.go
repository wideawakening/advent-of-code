package day3

import (
	"fmt"
	"os"
	"strings"
)

/*
case sensitive
1 rucksack - two compartments
items of given type either on A or B

	CompA
		- itemTypeA-1
		- itemTypeA-2
	CompB
		- itemTypeB-1

	item priority
	- a-z: priority 1-26
	- A-Z: priority 27-52
*/
func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", dataLine)
		retValue += GetDuplicatedItemValue(dataLine)

	}
	return retValue
}

/*
- elf in groups of 3 - ELF1, ELF2, ELF3
- badge type 'B', all 3 elfs within a group carry it
- at most two elfs carry other item type
*/
func Star2(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	var retValue = 0

	var i = 0
	lines := strings.Split(string(data), "\n")
	var totalElves = len(lines)

	for i+3 < totalElves {
		fmt.Printf("%d\n", i)
		if i%3 != 0 {
			fmt.Println("no more pack of three elves")
			return retValue
		}
		retValue += GetItemPriorityValue(GetMatchingCharacter(lines[i], lines[i+1], lines[i+2]))
		i = i + 3
	}
	return retValue
}

func GetDuplicatedItemValue(ruckSackItems string) int {
	// 1. GetDuplicatedItem
	matchingCharacter := GetDuplicatedItem(ruckSackItems)

	// 2. GetItemPriorityValue
	if matchingCharacter != "" {
		return GetItemPriorityValue(matchingCharacter)
	}
	return 0
}

func GetDuplicatedItem(rucksackItems string) string {
	text1, text2 := GetSplitRuckSac(rucksackItems)
	return GetMatchingCharacter(text1, text2, "")
}

func GetSplitRuckSac(rucksackItems string) (string, string) {
	if len(rucksackItems)%2 != 0 {
		fmt.Errorf("Rucksack has an impair number of items")
		return "", ""
	}

	return rucksackItems[0 : len(rucksackItems)/2], rucksackItems[len(rucksackItems)/2:]
}

func GetMatchingCharacter(text1 string, text2 string, text3 string) string {
	for _, char1 := range strings.Split(text1, "") {
		for _, char2 := range strings.Split(text2, "") {
			if text3 != "" {
				for _, char3 := range strings.Split(text3, "") {
					if char1 == char2 && char1 == char3 {
						return char1
					}
				}
			} else {
				if char1 == char2 {
					return char1
				}
			}
		}
	}
	fmt.Errorf("No matching characters")
	return ""
}

func GetItemPriorityValue(matchingCharacter string) int {
	//fmt.Println("foooo")
	//fmt.Println(int('a'))
	//fmt.Println(int('b'))
	//fmt.Println(int('z')) // 122 - 96 = 26
	//fmt.Println(int('A')) // 65 - 38 = 27
	intValue := int(matchingCharacter[0])
	if intValue < 97 {
		return intValue - 38
	}
	return intValue - 96
}
