package deck

import (
	"errors"
	"fmt"
	"strings"
)

type Suit int

const (
	suitSize int = 4

	SPADES Suit = iota
	DIAMONDS
	CLUBS
	HEARTS
)

var (
	codeToSuit = map[string]Suit{
		"S": SPADES,
		"D": DIAMONDS,
		"C": CLUBS,
		"H": HEARTS,
	}
)

type Rank int

const (
	rankSize int = 13

	ACE Rank = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var (
	codeToRank = map[string]Rank{
		"A":  ACE,
		"1":  ONE,
		"2":  TWO,
		"3":  THREE,
		"4":  FOUR,
		"5":  FIVE,
		"6":  SIX,
		"7":  SEVEN,
		"8":  EIGHT,
		"9":  NINE,
		"10": TEN,
		"J":  JACK,
		"Q":  QUEEN,
		"K":  KING,
	}
)

type Card struct {
	suit Suit
	rank Rank
}

// CodeToCard parses code and return correspondent card
func CodeToCard(code string) (*Card, error) {
	codeLength := len(code)
	if codeLength < 2 || codeLength > 3 {
		return nil, errors.New("code does not meet length requirements")
	}

	rankCode := code[:codeLength-1]
	rank, found := codeToRank[strings.ToUpper(rankCode)]
	if !found {
		return nil, fmt.Errorf("rank code not recognized: %s", rankCode)
	}

	suitCode := code[codeLength-1:]
	suit, found := codeToSuit[strings.ToUpper(suitCode)]
	if !found {
		return nil, fmt.Errorf("suit code not recognized: %s", suitCode)
	}

	return &Card{
		rank: rank,
		suit: suit,
	}, nil
}

// NewCard returns a new card given suit and rank
func NewCard(suit Suit, rank Rank) (*Card, error) {
	if int(suit) < 1 || int(suit) > suitSize {
		return nil, errors.New("suit does not exist")
	}

	if int(rank) < 1 || int(rank) > rankSize {
		return nil, errors.New("suit does not exist")
	}

	return &Card{
		suit: suit,
		rank: rank,
	}, nil
}
