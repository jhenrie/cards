package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

func InitDeck() Deck {
	var deck = make([]Card, 52)
	for i := 0; i <= len(deck)-1; i++ {
		deck[i].value = (i % 13) + 2
		deck[i].symbol = NewSuit(i / 13)
	}

	return deck
}

func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func (deck Deck) String() string {
	var msg string

	for i := 0; i <= len(deck)-1; i++ {
		var cardValue, cardSuit = deck[i].Info()
		msg = fmt.Sprintf("%s[%s %s]", msg, cardValue, cardSuit)

		if i < len(deck)-1 {
			msg = fmt.Sprintf("%s, ", msg)
		}
	}
	msg = fmt.Sprintf("%s\n", msg)

	return msg
}

func (deck *Deck) Draw(numberOfCards int) Deck {
	draw := (*deck)[len(*deck)-numberOfCards : len(*deck)]
	*deck = (*deck)[:len(*deck)-numberOfCards]

	return draw
}

func (deck *Deck) DrawOne() Card {
	draw := deck.Draw(1)
	return draw[0]
}


func (deck *Deck) AddCard(card Card) {
	*deck = append(Deck{card}, *deck...)
}

func (deck *Deck) AddCards(cards Deck) {
	for _, card := range cards {
		deck.AddCard(card)
	}
}
