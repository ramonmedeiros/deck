package deck

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeckTestSuite struct {
	suite.Suite
	deckManager Manager
}

func (suite *DeckTestSuite) SetupTest() {
	suite.deckManager = NewManager()
}

func (suite *DeckTestSuite) TestDeckGenerationSize(t *testing.T) {
	newDeck, err := suite.deckManager.NewDeck(false)
	require.NoError(t, err)
	require.NotNil(t, newDeck.ID)
	require.NoError(t, uuid.Validate(newDeck.ID.String()))

	cards := newDeck.Open()
	require.Len(t, cards, 52)
}

func (suite *DeckTestSuite) TestDeckCheckUnshuffledOrder(t *testing.T) {
	newDeck, err := suite.deckManager.NewDeck(false)
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

func (suite *DeckTestSuite) TestDeckGenerationShuffle(t *testing.T) {
	newDeck, err := suite.deckManager.NewDeck(false)
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

func (suite *DeckTestSuite) TestDeckCustomGeneration(t *testing.T) {
	newDeck, err := suite.deckManager.NewDeck(false, "AS")
	require.NoError(t, err)

	cards := newDeck.Open()
	require.Len(t, cards, 1)
	require.Equal(t, SPADES, cards[0].suit)
	require.Equal(t, ACE, cards[0].rank)

	newDeck, err = suite.deckManager.NewDeck(false, "AS", "10H")
	require.NoError(t, err)

	cards = newDeck.Open()
	require.Len(t, cards, 2)
	require.Equal(t, SPADES, cards[0].suit)
	require.Equal(t, ACE, cards[0].rank)
	require.Equal(t, HEARTS, cards[1].suit)
	require.Equal(t, TEN, cards[1].rank)
}

func (suite *DeckTestSuite) TestDeckDrawCards(t *testing.T) {
	newDeck, err := suite.deckManager.NewDeck(false)
	require.NoError(t, err)

	cards := newDeck.Open()
	require.Len(t, cards, 52)

	cards = newDeck.DrawCard(1)
	require.Len(t, cards, 1)
	require.Len(t, newDeck.Open(), 51)

	cards = newDeck.DrawCard(10)
	require.Len(t, cards, 10)
	require.Len(t, newDeck.Open(), 41)
}
