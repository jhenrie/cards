package cards

import (
	"fmt"
)

type Card struct {
	value int
	symbol Suit
}

func NewCard(value int, suit Suit) *Card {
	if value < 2 || value > 14 {
		return nil
	}
	return &Card{
		value: value,
		symbol: suit,
	}
}

func (me Card) Info() (string, string) {
	var cardValue, cardSuit string

	switch me.value {
	case 11:
		cardValue = "J"
	case 12:
		cardValue = "Q"
	case 13:
		cardValue = "K"
	case 14:
		cardValue = "A"
	default:
		cardValue = fmt.Sprint(me.value)
	}

	switch me.symbol {
	case 0:
		cardSuit = "♦"
	case 1:
		cardSuit = "♣"
	case 2:
		cardSuit = "♥"
	case 3:
		cardSuit = "♠"
	default:
	}

	return cardValue, cardSuit
}

func (me Card) String() string {
	value, suit := me.Info()
	return fmt.Sprintf("%s %s", value, suit)
}

func (me Card) Equals(card Card) bool {
	isEquals := true

	if me.symbol != card.symbol || me.value != card.value {
		isEquals = false
	}

	return isEquals
}