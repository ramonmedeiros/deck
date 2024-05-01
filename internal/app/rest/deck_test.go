package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
	"github.com/ramonmedeiros/deck/internal/pkg/deck/decktest"
	"github.com/stretchr/testify/require"
)

func TestNewDeck(t *testing.T) {
	managerMock := decktest.ManagerMock{
		NewDeckFunc: func(bool, ...string) (*deck.Deck, error) {
			return &deck.Deck{ID: uuid.New()}, nil
		},
	}
	server := New("8080", slog.Default(), &managerMock)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/deck", nil)
	require.NoError(t, err)

	server.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var response *NewDeckResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&response))
	require.NotNil(t, response)
	require.False(t, response.Shuffled)
	require.Zero(t, response.Remaining)
	require.NoError(t, uuid.Validate(response.ID))
}
