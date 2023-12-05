package day11

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 */
func InputFileToMonkeys(inputFileName string) ([]*Monkey, error) {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("incorrect monkey syntax")
	}

	monkeys := []*Monkey{}
	monkey := NewMonkey()
	dataLines := strings.Split(string(data), "\n")

	for i := 0; i < len(dataLines)-1; i += 7 {
		//fmt.Printf("%s", dataLine)
		monkey, err = InputToMonkeys(dataLines[i : i+7])
		if err != nil {
			return nil, fmt.Errorf("incorrect monkey syntax")
		}
		monkeys = append(monkeys, monkey)
		monkey = NewMonkey()
	}
	return monkeys, nil
}

func PlayRound(monkeys []*Monkey, WorryLevelDivisor *big.Int) error {

	for monkeyIdx, monkey := range monkeys {
		for _, worryLevel := range monkey.StartingItemsWorryLevel {
			//fmt.Printf("\n%v \n", worryLevel)
			worryLevel := monkey.Operation(&worryLevel)
			//fmt.Printf("%v \n", worryLevel)

			//divisor
			//if WorryLevelDivisor > 1 {
			//	worryLevel = worryLevel.Div(worryLevel, big.NewInt(int64(WorryLevelDivisor)))
			//}
			//if WorryLevelDivisor > 1 {
			worryLevel = worryLevel.Mod(worryLevel, WorryLevelDivisor)
			//}
			//fmt.Printf("%v \n", worryLevel)

			//if worryLevel == 0 {
			//	return fmt.Errorf("no proper divisor")
			//}

			destinationMonkey := monkey.MonkeyFalse
			worryLevelCopy := big.NewInt(0)
			worryLevelCopy = worryLevelCopy.Set(worryLevel)
			if worryLevel.Mod(worryLevel, big.NewInt(int64(monkey.DivisibleValue))).Cmp(big.NewInt(int64(0))) == 0 {
				destinationMonkey = monkey.MonkeyTrue
			}
			//fmt.Printf("%v \n", destinationMonkey)

			monkeys[destinationMonkey].StartingItemsWorryLevel = append(monkeys[destinationMonkey].StartingItemsWorryLevel, *worryLevelCopy)

			if len(monkeys[monkeyIdx].StartingItemsWorryLevel) > 1 {
				monkeys[monkeyIdx].StartingItemsWorryLevel = monkeys[monkeyIdx].StartingItemsWorryLevel[1:]
			} else {
				monkeys[monkeyIdx].StartingItemsWorryLevel = []big.Int{}
			}
			//fmt.Printf("\n\n")
			monkey.ItemsInspected++
		}
	}

	return nil
}

func InputToMonkeys(dataLines []string) (*Monkey, error) {
	if len(dataLines) != 7 {
		return nil, fmt.Errorf("incorrect monkey syntax")
	}

	monkey := NewMonkey()
	monkeyNumber, err := strconv.Atoi(dataLines[0][7:8])
	if err != nil {
		return nil, fmt.Errorf("incorrect monkey syntax")
	}
	monkey.Number = monkeyNumber

	startingItemsTxt := dataLines[1][18:]
	for _, startingItemTxt := range strings.Split(startingItemsTxt, ",") {
		//startingItem, err := strconv.Atoi(strings.TrimSpace(startingItemTxt))
		startingItem, err := strconv.ParseUint(strings.TrimSpace(startingItemTxt), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("incorrect monkey syntax")
		}
		monkey.StartingItemsWorryLevel = append(monkey.StartingItemsWorryLevel, *big.NewInt(int64(startingItem)))
	}

	operationOperatorTxt := string(dataLines[2][23])
	operationValueTxt := string(dataLines[2][25:])
	switch operationOperatorTxt {
	case "+":
		operationValue, err := strconv.Atoi(operationValueTxt)
		if err != nil {
			if operationValueTxt == "old" {
				monkey.Operation = func(old *big.Int) *big.Int {
					return old.Add(old, old)
				}
			}
			return nil, fmt.Errorf("incorrect monkey syntax")
		} else {
			monkey.Operation = func(old *big.Int) *big.Int {
				return old.Add(old, big.NewInt(int64(operationValue)))
			}
		}
		break

	case "*":
		operationValue, err := strconv.Atoi(operationValueTxt)
		if err != nil {
			if operationValueTxt == "old" {
				monkey.Operation = func(old *big.Int) *big.Int {
					return old.Mul(old, old)
				}
			} else {
				return nil, fmt.Errorf("incorrect monkey syntax")
			}

		} else {
			monkey.Operation = func(old *big.Int) *big.Int {
				return old.Mul(old, big.NewInt(int64(operationValue)))
			}
		}

		break

	default:
		return nil, fmt.Errorf("incorrect monkey syntax")
	}

	divisibleValueTxt := dataLines[3][21:]
	divisibleValue, err := strconv.Atoi(divisibleValueTxt)
	if err != nil {
		return nil, fmt.Errorf("incorrect monkey syntax")

	}
	monkey.DivisibleValue = divisibleValue

	monkeyTrueTxt := dataLines[4][29:]
	monkeyTrue, err := strconv.Atoi(monkeyTrueTxt)
	if err != nil {
		return nil, fmt.Errorf("incorrect monkey syntax")

	}
	monkey.MonkeyTrue = monkeyTrue

	monkeyFalseTxt := dataLines[5][30:]
	monkeyFalse, err := strconv.Atoi(monkeyFalseTxt)
	if err != nil {
		return nil, fmt.Errorf("incorrect monkey syntax")
	}
	monkey.MonkeyFalse = monkeyFalse

	return monkey, nil
}

type Monkey struct {
	Number                  int
	StartingItemsWorryLevel []big.Int
	Operation               func(old *big.Int) *big.Int
	DivisibleValue          int
	MonkeyTrue              int
	MonkeyFalse             int

	ItemsInspected int
}

func NewMonkey() *Monkey {
	return new(Monkey)
}
