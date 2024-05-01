package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
)

const (
	Quantity = "count"
)

type DrawCardResponse struct {
	Cards []CardResponse `json:"cards"`
}

func (s *Server) drawCard(c *gin.Context) {
	countString := c.Request.URL.Query().Get(Quantity)
	count := 0
	if countString != "" {
		var err error
		count, err = strconv.Atoi(countString)
		if err != nil {
			s.logger.Error("could not parse count", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

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
		DrawCardResponse{
			Cards: fromDeckToCardResponse(connectedDeck.DrawCard(count)),
		})
}
