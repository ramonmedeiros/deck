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
