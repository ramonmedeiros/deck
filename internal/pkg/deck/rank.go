package deck

type Rank int

const (
	rankSize int = 13

	ACE Rank = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var (
	codeToRank = map[string]Rank{
		"A":  ACE,
		"1":  ONE,
		"2":  TWO,
		"3":  THREE,
		"4":  FOUR,
		"5":  FIVE,
		"6":  SIX,
		"7":  SEVEN,
		"8":  EIGHT,
		"9":  NINE,
		"10": TEN,
		"J":  JACK,
		"Q":  QUEEN,
		"K":  KING,
	}
)

func (r *Rank) String() string {
	return map[Rank]string{
		ACE:   "ACE",
		ONE:   "1",
		TWO:   "2",
		THREE: "3",
		FOUR:  "4",
		FIVE:  "5",
		SIX:   "6",
		SEVEN: "7",
		EIGHT: "8",
		NINE:  "9",
		TEN:   "10",
		JACK:  "J",
		QUEEN: "QUEEN",
		KING:  "KING",
	}[*r]
}

func (r *Rank) Code() string {
	return map[Rank]string{
		ACE:   "A",
		ONE:   "1",
		TWO:   "2",
		THREE: "3",
		FOUR:  "4",
		FIVE:  "5",
		SIX:   "6",
		SEVEN: "7",
		EIGHT: "8",
		NINE:  "9",
		TEN:   "10",
		JACK:  "J",
		QUEEN: "Q",
		KING:  "K",
	}[*r]
}
