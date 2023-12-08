package challenge

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var JokerOn = false

func Star1(inputFileName string) int {

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	fmt.Printf("\n")
	hands := []*HandData{}
	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)

		hand, err := GetHandData(dataLine)
		if err != nil {
			fmt.Printf("invalid data line: %s\n", dataLine)
			continue
		}

		hands = append(hands, hand)

	}

	OrderHands(hands)
	var retValue = 0
	for ranking := len(hands); ranking > 0; ranking-- {
		hand := hands[ranking-1]
		fmt.Printf("%v\n", hand)
		retValue = retValue + (hand.bid * (ranking))
	}
	return retValue
}

type HandData struct {
	hand     Hand
	HandType HandType
	bid      int
}

type Hand string
type CardType string
type HandType string

const (
	FiveOfAKind  HandType = "FiveOfAKind"
	FourOfAKind  HandType = "FourOfAKind"
	FullHouse    HandType = "FullHouse"
	ThreeOfAKind HandType = "ThreeOfAKind"
	TwoPair      HandType = "TwoPair"
	OnePair      HandType = "OnePair"
	HighCard     HandType = "HighCard"
)

func GetHighestHand(card1 Hand, card2 Hand) Hand {
	card1HandType := Borrowed_GetHandType(card1)
	card2HandType := Borrowed_GetHandType(card2)

	fmt.Printf("%s is %s - %s is %s\n", card1, card1HandType, card2, card2HandType)

	if card1HandType == card2HandType {
		return GetHighestOnTie(card1, card2)
	}

	if card1HandType == FiveOfAKind {
		return card1
	}
	if card2HandType == FiveOfAKind {
		return card2
	}

	if card1HandType == FourOfAKind {
		return card1
	}
	if card2HandType == FourOfAKind {
		return card2
	}

	if card1HandType == FullHouse {
		return card1
	}
	if card2HandType == FullHouse {
		return card2
	}

	if card1HandType == ThreeOfAKind {
		return card1
	}
	if card2HandType == ThreeOfAKind {
		return card2
	}

	if card1HandType == TwoPair {
		return card1
	}
	if card2HandType == TwoPair {
		return card2
	}

	if card1HandType == OnePair {
		return card1
	}
	if card2HandType == OnePair {
		return card2
	}

	return GetHighestOnTie(card1, card2)
}

func GetHighestOnTie(hand1 Hand, hand2 Hand) Hand {

	for i := 0; i < len(hand1); i++ {
		card1 := string(hand1[i])
		card2 := string(hand2[i])
		card1Num, isCharacterCard1 := strconv.Atoi(card1)
		card2Num, isCharacterCard2 := strconv.Atoi(card2)

		// -- both cards are character
		if isCharacterCard1 != nil && isCharacterCard2 != nil {

			if card1 == card2 {
				continue
			}

			if card1 == "J" || card2 == "J" {
				if JokerOn {
					if card1 == "J" {
						return hand2
					}
					if card2 == "J" {
						return hand1
					}
				}

				if card1 == "K" {
					return hand1
				}
				if card2 == "K" {
					return hand2
				}
				if card1 == "Q" {
					return hand1
				}
				if card2 == "Q" {
					return hand2
				}

			}

			if card1 > card2 {
				return hand2
			} else {
				return hand1
			}
		}

		// -- hand1 is character, hand2 is number
		if isCharacterCard1 != nil && isCharacterCard2 == nil {

			if JokerOn {
				if card1 == "J" {
					return hand2
				}
			}
			return hand1
		}

		// -- hand1 is number, hand2 is character
		if isCharacterCard2 != nil && isCharacterCard1 == nil {
			if JokerOn {
				if card2 == "J" {
					return hand1
				}
			}
			return hand2
		}

		// -- hand1 and 2 are numbers
		if card1 == card2 {
			continue
		}
		if card1Num > card2Num {
			return hand1
		} else {
			return hand2
		}
	}
	return hand1
}

func Borrowed_GetHandType(hand Hand) HandType {
	counts := make(map[rune]int, len(hand))
	hasThreeOfAKind := false
	pairs, jokers := 0, 0

	for _, label := range hand {
		if JokerOn && label == 'J' {
			jokers++
		} else {
			counts[label]++
		}
	}

	if JokerOn {
		maximumKey := '0'
		maximumValue := 0
		for label, count := range counts {
			if count > maximumValue {
				maximumKey = label
				maximumValue = count
			}
		}

		counts[maximumKey] += jokers
	}

	for _, count := range counts {
		switch count {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			hasThreeOfAKind = true
		case 2:
			pairs++
		}
	}

	if hasThreeOfAKind && pairs == 1 {
		return FullHouse
	}

	if hasThreeOfAKind {
		return ThreeOfAKind
	}

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func FailedHandType(hand Hand) HandType {

	orderedHand := OrderHandCards(hand)

	if len(strings.Trim(orderedHand, string(orderedHand[0]))) == 0 {
		return FiveOfAKind
	}

	if len(strings.Trim(orderedHand, string(orderedHand[0]))) == 1 ||
		len(strings.Trim(orderedHand, string(orderedHand[4]))) == 1 {
		return FourOfAKind
	}

	if len(strings.Trim(orderedHand, string(orderedHand[0]))) == 2 {
		if len(strings.Trim(orderedHand, string(orderedHand[4]))) == 3 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}
	if len(strings.Trim(orderedHand, string(orderedHand[4]))) == 2 {
		if len(strings.Trim(orderedHand, string(orderedHand[0]))) == 3 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}

	if len(strings.Trim(orderedHand, string(orderedHand[0]))) == 3 {
		if orderedHand[2] == orderedHand[3] || orderedHand[3] == orderedHand[4] {
			return TwoPair
		} else {
			return OnePair
		}
	}
	if len(strings.Trim(orderedHand[1:], string(orderedHand[1]))) == 2 {
		if orderedHand[3] == orderedHand[4] {
			return TwoPair
		} else {
			return OnePair
		}
	}
	if len(strings.Trim(orderedHand[2:], string(orderedHand[2]))) == 1 {
		if orderedHand[0] == orderedHand[1] {
			return TwoPair
		} else {
			return OnePair
		}
	}
	if len(strings.Trim(orderedHand[3:], string(orderedHand[3]))) == 0 {
		if orderedHand[0] == orderedHand[1] || orderedHand[1] == orderedHand[2] {
			return TwoPair
		} else {
			return OnePair
		}
	}
	return HighCard
}

func OrderHandCards(hand Hand) string {
	s := []rune(hand)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	orderedHand := string(s)
	return orderedHand
}

func OrderHands(handData []*HandData) []*HandData {
	sort.SliceStable(handData, func(i int, j int) bool {
		card1 := handData[i].hand
		card2 := handData[j].hand
		highestCard := GetHighestHand(card1, card2)
		fmt.Printf("%s wins | %s - %s | %s - %s\n", highestCard, card1, card2, OrderHandCards(card1), OrderHandCards(card2))
		return highestCard == card2
	})
	return handData
}

func GetHandData(dataLine string) (*HandData, error) {

	dataTokens := strings.Split(dataLine, " ")
	if len(dataTokens) != 2 {
		return nil, fmt.Errorf("invalid data line")
	}

	hand := dataTokens[0]
	if !IsValidHand(Hand(hand)) {
		return nil, fmt.Errorf("invalid hand")
	}

	bidTxt := dataTokens[1]
	bid, err := strconv.Atoi(bidTxt)
	if err != nil {
		return nil, fmt.Errorf("invalid bidTxt")
	}

	//fmt.Printf("%s %d\n", hand, bid)
	return &HandData{hand: Hand(hand), HandType: Borrowed_GetHandType(Hand(hand)), bid: bid}, nil
}

func IsValidHand(hand Hand) bool {
	if len(hand) != 5 {
		return false
	}

	for _, card := range hand {
		if card < '2' || card > '9' {
			if card != 'T' && card != 'J' && card != 'Q' && card != 'K' && card != 'A' {
				return false
			}
		}
	}
	return true
}
