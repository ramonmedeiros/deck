package deck

import (
	"github.com/google/uuid"
)

type Deck struct {
	ID       uuid.UUID
	cards    []*Card
	shuffled bool
}

type API interface {
	Open() []*Card
	DrawCard(int) []*Card
	Remaining() int
	Shuffled() bool
}

// Open returns all cards from the deck
func (d *Deck) Open() []*Card {
	// avoid to pass the pointer to the internal array
	newCardsArray := d.cards
	return newCardsArray
}

// Remaining returns remaining amount of cards
func (d *Deck) Remaining() int {
	return len(d.cards)
}

// Shuffled returns if deck was shuffled during creation
func (d *Deck) Shuffled() bool {
	return d.shuffled
}

// DrawCard return cards according to requested quantity
func (d *Deck) DrawCard(quantity int) []*Card {
	var cards []*Card
	for i := 0; i < quantity; i++ {
		var card *Card
		card, d.cards = d.cards[0], d.cards[1:]
		cards = append(cards, card)
	}

	return cards
}
