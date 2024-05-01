package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
)

type OpenDeckResponse struct {
	ID        string         `json:"deck_id"`
	Shuffled  bool           `json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     []CardResponse `json:"cards"`
}

type CardResponse struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func fromDeckToCardResponse(deckCards []*deck.Card) []CardResponse {
	var cards []CardResponse
	for _, card := range deckCards {
		cards = append(cards, CardResponse{
			Value: card.Value(),
			Suit:  card.Suit(),
			Code:  card.Code(),
		})
	}
	return cards
}

func (s *Server) openDeck(c *gin.Context) {
	deckID := c.Param(DeckID)
	connectedDeck, err := s.deckManager.Get(deckID)
	if err == deck.ErrNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		s.logger.Error("could not search for deck", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		OpenDeckResponse{
			ID:        connectedDeck.ID.String(),
			Remaining: connectedDeck.Remaining(),
			Shuffled:  connectedDeck.Shuffled(),
			Cards:     fromDeckToCardResponse(connectedDeck.Open()),
		})
}
