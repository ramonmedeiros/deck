package deck

import (
	"errors"
	"fmt"
	"strings"
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

func (c *Card) Value() string {
	return c.rank.String()
}

func (c *Card) Suit() string {
	return c.suit.String()
}

func (c *Card) Code() string {
	return c.rank.Code() + c.suit.Code()
}
