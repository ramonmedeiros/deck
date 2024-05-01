package deck

import "errors"

type Suit int

const (
	suitSize int = 4

	SPADES Suit = iota
	DIAMONDS
	CLUBS
	HEARTS
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

type Card struct {
	suit Suit
	rank Rank
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
