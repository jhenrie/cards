package cards

type Suit int

const (
	Hearts Suit = iota
	Diamonds Suit = iota
	Clubs Suit = iota
	Spades Suit = iota
)

func (me Suit) String() string {
	return [...]string{"Hearts", "Diamonds", "Clubs", "Spades"}[me]
}

func (me Suit) Int() int {
	return me.Int()
}

func NewSuit(value int) Suit {
	var result Suit
	switch value {
	case 0:
		result = Hearts
	case 1:
		result = Diamonds
	case 2:
		result = Clubs
	case 3:
		result = Spades
	}
	return result
}