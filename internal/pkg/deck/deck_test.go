package deck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeckGeneration(t *testing.T) {
	newDeck, err := NewDeck()
	require.NoError(t, err)

	require.Len(t, newDeck.cards, 52)

}
