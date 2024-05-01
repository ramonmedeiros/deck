package deck

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type DeckManager struct {
	decks map[uuid.UUID]*Deck
}

//go:generate moq -pkg decktest -skip-ensure -stub -out decktest/mock.go . Manager:ManagerMock
type Manager interface {
	Open(deckID uuid.UUID) (*Deck, error)
	NewDeck(shuffled bool, cards ...string) (*Deck, error)
}

// NewManager generates a new deck with all cars, or respecting the parameters
func NewManager() *DeckManager {
	return &DeckManager{
		decks: make(map[uuid.UUID]*Deck),
	}
}

func (m *DeckManager) Open(deckID uuid.UUID) (*Deck, error) {
	deck, found := m.decks[deckID]
	if !found {
		return nil, fmt.Errorf("could not found deck by id: %s", deckID.String())
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
