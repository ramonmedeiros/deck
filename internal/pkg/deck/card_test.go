package deck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCardCreation(t *testing.T) {
	newCard, err := NewCard(0, 0)
	require.Nil(t, newCard)
	require.Error(t, err)

	newCard, err = NewCard(0, 1)
	require.Nil(t, newCard)
	require.Error(t, err)

	newCard, err = NewCard(1, 1)
	require.NotNil(t, newCard)
	require.NoError(t, err)
	require.Equal(t, SPADES, newCard.suit)
	require.Equal(t, ACE, newCard.rank)
}

func TestCardCodeParse(t *testing.T) {
	aceSpades, err := CodeToCard("AS")
	require.NoError(t, err)
	require.NotNil(t, aceSpades)
	require.Equal(t, SPADES, aceSpades.suit)
	require.Equal(t, ACE, aceSpades.rank)

	kindDiamonds, err := CodeToCard("kd")
	require.NoError(t, err)
	require.NotNil(t, kindDiamonds)
	require.Equal(t, DIAMONDS, kindDiamonds.suit)
	require.Equal(t, KING, kindDiamonds.rank)

	tenClubs, err := CodeToCard("10C")
	require.NoError(t, err)
	require.NotNil(t, tenClubs)
	require.Equal(t, CLUBS, tenClubs.suit)
	require.Equal(t, TEN, tenClubs.rank)

	nineHearts, err := CodeToCard("9H")
	require.NoError(t, err)
	require.NotNil(t, nineHearts)
	require.Equal(t, HEARTS, nineHearts.suit)
	require.Equal(t, NINE, nineHearts.rank)
}
