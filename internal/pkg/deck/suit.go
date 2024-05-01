package deck

type Suit int

const (
	suitSize int = 4

	SPADES Suit = iota
	DIAMONDS
	CLUBS
	HEARTS
)

var (
	codeToSuit = map[string]Suit{
		"S": SPADES,
		"D": DIAMONDS,
		"C": CLUBS,
		"H": HEARTS,
	}
)

func (s *Suit) String() string {
	return map[Suit]string{
		SPADES:   "SPADES",
		DIAMONDS: "DIAMONDS",
		CLUBS:    "CLUBS",
		HEARTS:   "HEARTS",
	}[*s]
}

func (s *Suit) Code() string {
	return map[Suit]string{
		SPADES:   "S",
		DIAMONDS: "D",
		CLUBS:    "C",
		HEARTS:   "H",
	}[*s]
}
