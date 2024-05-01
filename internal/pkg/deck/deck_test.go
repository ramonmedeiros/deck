package deck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeckGenerationSize(t *testing.T) {
	newDeck, err := NewDeck(false)
	require.NoError(t, err)

	cards := newDeck.Open()
	require.Len(t, cards, 52)
}

func TestDeckCheckUnshuffledOrder(t *testing.T) {
	newDeck, err := NewDeck(false)
	require.NoError(t, err)

	cards := newDeck.Open()
	require.Len(t, cards, 52)

	// expected order is by spades, diamonds, clubs, then hearts. Followed
	// by ranks A, 2 .. 10, J, Q, K
	currentCard := 0
	for suit := 1; suit <= 4; suit++ {
		for rank := 1; rank <= 13; rank++ {
			require.EqualValues(t, rank, cards[currentCard].rank)
			require.EqualValues(t, suit, cards[currentCard].suit)
			currentCard++
		}
	}
}

func TestDeckGenerationShuffle(t *testing.T) {
	newDeck, err := NewDeck(true)
	require.NoError(t, err)
	cards := newDeck.Open()

	require.Len(t, cards, 52)

	// expected order is by spades, diamonds, clubs, then hearts. Followed
	// by ranks A, 2 .. 10, J, Q, K
	// let's check how many cars are out of order
	currentCard := 0
	outOfOrder := 0
	for suit := 1; suit <= 4; suit++ {
		for rank := 1; rank <= 13; rank++ {
			notExpectedRank := rank != int(cards[currentCard].rank)
			notExpectedSuit := suit != int(cards[currentCard].suit)
			currentCard++

			if notExpectedRank || notExpectedSuit {
				outOfOrder++
			}
		}
	}

	t.Log("outOfOrder ", outOfOrder, "/", len(cards))
	require.NotZero(t, outOfOrder)
}
