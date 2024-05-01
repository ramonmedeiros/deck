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

func TestNewDeckCustom(t *testing.T) {
	manager := deck.NewManager()
	server := New("8080", slog.Default(), manager)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/deck?cards=AS,2S", nil)
	require.NoError(t, err)

	server.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var response *NewDeckResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&response))
	require.NotNil(t, response)

	require.False(t, response.Shuffled)
	require.Equal(t, 2, response.Remaining)
	require.NoError(t, uuid.Validate(response.ID))

	cards, err := manager.Get(response.ID)
	require.NoError(t, err)

	require.Equal(t, "AS", cards.Open()[0].Code())
	require.Equal(t, "2S", cards.Open()[1].Code())

}

func TestOpenDeck(t *testing.T) {
	manager := deck.NewManager()
	server := New("8080", slog.Default(), manager)

	newDeck, err := manager.NewDeck(false)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/deck/"+newDeck.ID.String(), nil)
	require.NoError(t, err)

	server.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var response *OpenDeckResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&response))

	require.NotNil(t, response)
	require.False(t, response.Shuffled)
	require.Equal(t, 52, response.Remaining)
	require.Equal(t, newDeck.ID.String(), response.ID)

	for index, card := range newDeck.Open() {
		require.Equal(t, card.Code(), response.Cards[index].Code)
		require.Equal(t, card.Value(), response.Cards[index].Value)
		require.Equal(t, card.Suit(), response.Cards[index].Suit)
	}
}

func TestDrawCard(t *testing.T) {
	manager := deck.NewManager()
	server := New("8080", slog.Default(), manager)

	newDeck, err := manager.NewDeck(false)
	require.NoError(t, err)

	cards := newDeck.Open()
	drawCards := []*deck.Card{cards[0], cards[1]}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/deck/"+newDeck.ID.String()+"/draw?count=2", nil)
	require.NoError(t, err)

	server.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var response *OpenDeckResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&response))

	require.NotNil(t, response)

	for index, card := range drawCards {
		require.Equal(t, card.Code(), response.Cards[index].Code)
		require.Equal(t, card.Value(), response.Cards[index].Value)
		require.Equal(t, card.Suit(), response.Cards[index].Suit)
	}
}
