package deck

type Deck struct {
	cards []*Card
}

func NewDeck(cards ...string) (*Deck, error) {
	d := Deck{}

	for s := 1; s <= suitSize; s++ {
		for r := 1; r <= rankSize; r++ {

			newCard, err := NewCard(Suit(s), Rank(r))
			if err != nil {
				return nil, err
			}
			d.cards = append(d.cards, newCard)
		}
	}
	return &d, nil
}

func (d *Deck) Open() {}

func (d *Deck) DrawCard() *Card {
	return &Card{}
}
