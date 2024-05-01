package deck

import (
	"math/rand"
)

type Deck struct {
	cards    []*Card
	shuffled bool
}

// NewDeck generates a new deck with all cars, or respecting the parameters
func NewDeck(shuffled bool, cards ...string) (*Deck, error) {
	d := Deck{shuffled: shuffled}

	for s := 1; s <= suitSize; s++ {
		for r := 1; r <= rankSize; r++ {

			newCard, err := NewCard(Suit(s), Rank(r))
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

// Open returns all cards from the deck
func (d *Deck) Open() []*Card {
	return d.cards
}

func (d *Deck) DrawCard() *Card {
	return &Card{}
}
