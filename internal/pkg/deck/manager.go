package deck

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type DeckManager struct {
	decks map[uuid.UUID]*Deck
}

var ErrNotFound = errors.New("could not find deck")

//go:generate moq -pkg decktest -skip-ensure -stub -out decktest/mock.go . Manager:ManagerMock
type Manager interface {
	Get(deckID string) (*Deck, error)
	NewDeck(shuffled bool, cards ...string) (*Deck, error)
}

// NewManager generates a new deck with all cars, or respecting the parameters
func NewManager() *DeckManager {
	return &DeckManager{
		decks: make(map[uuid.UUID]*Deck),
	}
}

// Get returns deck for a given ID
func (m *DeckManager) Get(deckID string) (*Deck, error) {
	deckUUID, err := uuid.Parse(deckID)
	if err != nil {
		return nil, fmt.Errorf("could not parse id %s: %w", deckUUID.String(), err)
	}

	deck, found := m.decks[deckUUID]
	if !found {
		return nil, ErrNotFound
	}
	return deck, nil
}

// NewDeck generates a new deck with all cars, or respecting the parameters
func (m *DeckManager) NewDeck(shuffled bool, cards ...string) (*Deck, error) {
	d := Deck{
		ID:       uuid.New(),
		shuffled: shuffled,
	}
	m.decks[d.ID] = &d

	if len(cards) == 0 {
		for s := 1; s <= suitSize; s++ {
			for r := 1; r <= rankSize; r++ {

				newCard, err := NewCard(Suit(s), Rank(r))
				if err != nil {
					return nil, err
				}
				d.cards = append(d.cards, newCard)
			}
		}
	} else {
		for _, card := range cards {
			newCard, err := CodeToCard(card)
			if err != nil {
				return nil, err
			}
			d.cards = append(d.cards, newCard)
		}
	}

	if shuffled {
		rand.Shuffle(len(d.cards), func(first, second int) {
			d.cards[first], d.cards[second] = d.cards[second], d.cards[first]
		})
	}

	return &d, nil
}
