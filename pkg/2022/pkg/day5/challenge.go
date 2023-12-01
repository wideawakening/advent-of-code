package day5

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
crane moving cargo from one crate to another
*/
func Star1(blockMode bool, inputFileName string) string {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return ""
	}

	piles := [][]string{
		{"T", "R", "G", "W", "Q", "M", "F", "P"},
		{"R", "F", "H"},
		{"D", "S", "H", "G", "V", "R", "Z", "P"},
		{"G", "W", "F", "B", "P", "H", "Q"},
		{"H", "J", "M", "S", "P"},
		{"L", "P", "R", "S", "H", "T", "Z", "M"},
		{"L", "M", "N", "H", "T", "P"},
		{"R", "Q", "D", "F"},
		{"H", "P", "L", "N", "C", "S", "D"}}
	var moves [][]int

	var retValue = ""
	for _, dataLine := range strings.Split(string(data), "\n") {
		fmt.Printf("%s", dataLine)
		if dataLine == "" {
			fmt.Printf("empty line, continue")
			continue
		}

		moveData := strings.Split(dataLine, " ")
		quantity, _ := strconv.Atoi(moveData[1])
		pileFrom, _ := strconv.Atoi(moveData[3])
		pileTo, _ := strconv.Atoi(moveData[5])

		if quantity >= 0 && pileFrom >= 0 && pileTo <= 9 {
			moves = append(moves, []int{quantity, pileFrom, pileTo})
		}
	}

	pileList := CreateStacks(piles)
	for i := 0; i < len(moves); i++ {
		pileList = Move(blockMode, pileList, moves[i])
	}

	for i := 0; i < len(pileList); i++ {
		fmt.Println(pileList[i].Front().Value)
	}
	return retValue
}

func Move(blockMode bool, piles []*list.List, moves []int) []*list.List {
	quantity := moves[0]

	fromPile := piles[moves[1]-1]
	toPile := piles[moves[2]-1]

	if blockMode {
		moveBlock(quantity, fromPile, toPile)
	} else {
		move(quantity, fromPile, toPile)
	}

	return piles
}

func move(quantity int, fromPile *list.List, toPile *list.List) (*list.List, *list.List) {
	if quantity >= fromPile.Len() {
		quantity = fromPile.Len()
	}
	for i := quantity; i > 0; i-- {
		item := fromPile.Remove(fromPile.Front())
		toPile.PushFront(item)
	}
	return fromPile, toPile
}

func moveBlock(quantity int, fromPile *list.List, toPile *list.List) (*list.List, *list.List) {
	if quantity >= fromPile.Len() {
		quantity = fromPile.Len()
	}

	var tmpData []string

	for i := quantity; i > 0; i-- {
		item := fromPile.Remove(fromPile.Front())
		tmpData = append(tmpData, item.(string))
	}

	for j := len(tmpData); j > 0; j-- {
		toPile.PushFront(tmpData[j-1])
	}

	return fromPile, toPile
}

func CreateStacks(crateInitials [][]string) []*list.List {
	var lists []*list.List
	lists = make([]*list.List, len(crateInitials))

	for i := 0; i < len(crateInitials); i++ {
		lists[i] = createStack(crateInitials[i])
	}

	return lists
}

func createStack(crateInitial []string) *list.List {
	list := list.New()
	for i := 0; i < len(crateInitial); i++ {
		list.PushBack(crateInitial[i])
	}
	return list
}
