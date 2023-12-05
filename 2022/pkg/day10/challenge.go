package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Star1
		  ProcessOperation
		 	- ProcessCycles: 2 for 'addx', 1 for 'nop'
				- operation outcome will not be ready until the cycle finishes
					noop; addx; addx; will be ready after 5 cycles, in cycle 6
		 	- ProcessX - starts with value 1
		  ConvertToOperation ("addx 15") > addx, 15

	      GetSignalStrength
			- X register * cycle number
			- only happens on cycles:20th, and every 40 cycles after that; 20, 60, 100, 140,...
			- The final result is the sum of each signal strength
*/

/*
Star2
  - CRT is 40 wide (0..39), 6 high, draws left>right, next row, ...
  - CRT draws 1 pixel per cycle with a '#'
  - X register, is the horizontal position of the middle - 3 pixels wide
  - CRT produces '#' if the regX position, matches of the 3 pixel beign drawn
    otherwise a '.'
*/
func Star1(inputFileName string) (int, int, int) {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0, 0, 0
	}

	registerX := new(int)
	*registerX = 1

	counterCycle := new(int)
	*counterCycle = 1

	signalStrength := 0
	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)
		signalStrength += ProcessOperationStar1(dataLine, counterCycle, registerX)
	}
	return signalStrength, *counterCycle, *registerX
}

func ProcessOperationStar1(dataline string, counterCycle *int, registerX *int) int {
	data := strings.Split(dataline, " ")
	if len(data) < 1 {
		fmt.Errorf("ProcessOperation, not correct data syntax")
		return 0
	}

	op := data[0]

	var preRegisterX int
	preRegisterX = *registerX

	var preCounterCycle int
	preCounterCycle = *counterCycle

	switch op {
	case "noop":
		*counterCycle += 1
		break

	case "addx":
		*counterCycle += 2
		if len(data) != 2 {
			fmt.Errorf("ProcessOperation, not correct data syntax")
			return 0
		}
		opArg, err := strconv.Atoi(data[1])
		if err != nil {
			fmt.Errorf("ProcessOperation, op argument not a number")
			return 0
		}
		*registerX += opArg
		break

	default:
		break
	}

	calculate, checkpoint := IsCalculateSignalStrength(preCounterCycle, *counterCycle)
	if calculate {
		if *counterCycle == checkpoint {
			cycleStrength := *registerX * checkpoint
			fmt.Printf("%d * %d = %d \n", *registerX, checkpoint, cycleStrength)
			return cycleStrength

		} else {
			cycleStrength := preRegisterX * checkpoint
			fmt.Printf("%d * %d = %d \n", preRegisterX, checkpoint, cycleStrength)
			return cycleStrength
		}
	}
	return 0
}

func Star2(inputFileName string) {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
	}

	registerX := new(int)
	*registerX = 1

	counterCycle := new(int)
	*counterCycle = 1

	ProcessCRT(1, 1)
	for _, dataLine := range strings.Split(string(data), "\n") {
		ProcessOperationStar2(dataLine, counterCycle, registerX)
	}
}

func ProcessOperationStar2(dataline string, counterCycle *int, registerX *int) {
	data := strings.Split(dataline, " ")
	if len(data) < 1 {
		fmt.Errorf("ProcessOperation, not correct data syntax")
	}

	op := data[0]

	var preRegisterX int
	preRegisterX = *registerX

	var preCounterCycle int
	preCounterCycle = *counterCycle

	switch op {
	case "noop":
		*counterCycle += 1
		ProcessCRT(*counterCycle, *registerX)
		break

	case "addx":
		*counterCycle += 2
		if len(data) != 2 {
			fmt.Errorf("ProcessOperation, not correct data syntax")
		}
		opArg, err := strconv.Atoi(data[1])
		if err != nil {
			fmt.Errorf("ProcessOperation, op argument not a number")
		}
		*registerX += opArg

		ProcessCRT(preCounterCycle+1, preRegisterX)
		ProcessCRT(preCounterCycle+2, *registerX)
		break

	default:
		break
	}

}

func ProcessCRT(cycle int, regx int) (bool, bool) {

	crtPosition := ConvertCycleToCRTPosition(cycle)
	litPixel := crtPosition >= regx && crtPosition <= regx+2

	if litPixel {
		fmt.Printf("█")
	} else {
		fmt.Printf("░")
	}
	if cycle%40 == 0 {
		fmt.Println()
		return litPixel, true
	}

	return litPixel, false
}

func ConvertCycleToCRTPosition(cycle int) int {
	if cycle < 40 {
		return cycle
	}

	position := cycle - ((cycle / 40) * 40)
	if position == 0 {
		return 40
	}
	return position
}

func IsCalculateSignalStrength(previousCycle int, postCycle int) (bool, int) {
	checkpoints := []int{20, 60, 100, 140, 180, 220, 260, 300, 340, 380, 420}

	for _, checkpoint := range checkpoints {
		if postCycle < checkpoint {
			return false, 0
		}

		if previousCycle == checkpoint {
			return false, 0
		}

		if previousCycle < checkpoint && postCycle >= checkpoint {
			return true, checkpoint
		}
	}
	return false, 0
}
