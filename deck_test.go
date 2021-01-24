package cards

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeck(t *testing.T) {
	t.Run("Init deck should return an array containing a complete deck", func(t *testing.T) {
		deck := InitDeck()

		assert.True(t, len(deck) == 52, fmt.Sprintf("Confiming the deck contains %d cards", len(deck)))

		var heartCards   int
		var spadeCards   int
		var clubCards    int
		var diamondCards int

		for _, card := range deck {
			switch card.symbol {
			case Hearts:
				heartCards = heartCards + 1
			case Spades:
				spadeCards = spadeCards + 1
			case Clubs:
				clubCards = clubCards + 1
			case Diamonds:
				diamondCards = diamondCards + 1
			}
		}
		assert.True(t, heartCards == 13, fmt.Sprintf("Confirming that their are only %d hearts", heartCards))
		assert.True(t, spadeCards == 13, fmt.Sprintf("Confirming that their are only %d spades", spadeCards))
		assert.True(t, clubCards == 13, fmt.Sprintf("Confirming that their are only %d clubs", clubCards))
		assert.True(t, diamondCards == 13, fmt.Sprintf("Confirming that their are only %d diamonds", diamondCards))

	})

	t.Run("draw card from deck", func(t *testing.T) {
		deck := InitDeck()
		assert.True(t, len(deck) == 52, fmt.Sprintf("Confiming the deck contains %d cards", len(deck)))

		deck.DrawOne()
		assert.True(t, len(deck) == 51, fmt.Sprintf("Confiming the deck contains %d cards", len(deck)))

	})
}
