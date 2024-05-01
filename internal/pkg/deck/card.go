package deck

import "errors"

type Suit int

const (
	suitSize int  = 4
	SPADES   Suit = iota + 1
	DIAMONDS
	CLUBS
	HEARTS
)

type Rank int

const (
	rankSize int = 13

	ACE Rank = iota + 1
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

type Card struct {
	suit Suit
	rank Rank
}

func NewCard(suit Suit, rank Rank) (*Card, error) {
	if int(suit) < 0 || int(suit) > suitSize {
		return nil, errors.New("suit does not exist")
	}

	if int(rank) < 0 || int(rank) > rankSize {
		return nil, errors.New("suit does not exist")
	}

	return &Card{
		suit: suit,
		rank: rank,
	}, nil
}
