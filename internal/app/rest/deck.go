package rest

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type NewDeckResponse struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

const (
	// url params
	DeckID = "id"

	// query params
	Shuffle = "shuffle"
	Cards   = "cards"
)

func (s *Server) setupEndpoint() {
	s.router.POST("/deck", s.newDeck)

	deckEndpoint := s.router.Group("/deck")

	deckEndpoint.GET("/:"+DeckID, s.openDeck)
	deckEndpoint.GET("/:"+DeckID+"/draw", s.drawCard)
}

func (s *Server) newDeck(c *gin.Context) {
	params := c.Request.URL.Query()

	shuffle := false
	shuffleArgument := params.Get(Shuffle)
	if shuffleArgument == "true" {
		shuffle = true
	}

	var cards []string
	codesArgument := params.Get(Cards)
	if codesArgument != "" {
		cards = strings.Split(codesArgument, ",")
	}

	newDeck, err := s.deckManager.NewDeck(shuffle, cards...)
	if err != nil {
		s.logger.Error("could not generate", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK,
		NewDeckResponse{
			ID:        newDeck.ID.String(),
			Remaining: newDeck.Remaining(),
			Shuffled:  newDeck.Shuffled(),
		})
}

func (s *Server) openDeck(c *gin.Context) {
}

func (s *Server) drawCard(c *gin.Context) {
}
