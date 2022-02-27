package cards

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard(t *testing.T) {
	t.Run("Test for card equality will return true when equal", func(t *testing.T) {
		fourHearts := NewCard(4, Hearts)
		test := NewCard(4, Hearts)

		assert.True(t, fourHearts.Equals(*test), "Equals works when comparing [%s == %s]", fourHearts, test)
	})

	t.Run("Card returns proper values", func(t *testing.T) {
		cardValues := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		for _, test := range cardValues {
			testHeart := NewCard(test, Hearts)
			assert.True(t, testHeart != nil, "Card is not nil")
			assert.True(t, testHeart.GetValue() == test, "Test that card returns proper value [value = %d]", test)
		}
	})

	t.Run("Card returns proper formatted value for joker", func(t *testing.T) {
		joker := NewCard(11, Diamonds)
		assert.True(t, joker != nil, "Card is not nil")
		assert.True(t, joker.GetValueFormatted() == "J", "Test that joker returns formatted value of 'J'")
	})

	t.Run("Card returns proper formatted value for queen", func(t *testing.T) {
		queen := NewCard(12, Diamonds)
		assert.True(t, queen != nil, "Card is not nil")
		assert.True(t, queen.GetValueFormatted() == "Q", "Test that queen returns formatted value of 'Q'")
	})

	t.Run("Card returns proper formatted value for king", func(t *testing.T) {
		king := NewCard(13, Diamonds)
		assert.True(t, king != nil, "Card is not nil")
		assert.True(t, king.GetValueFormatted() == "K", "Test that king returns formatted value of 'K'")
	})

	t.Run("Card returns proper formatted value for ace", func(t *testing.T) {
		joker := NewCard(14, Diamonds)
		assert.True(t, joker != nil, "Card is not nil")
		assert.True(t, joker.GetValueFormatted() == "A", "Test that joker returns formatted value of 'A'")
	})

	t.Run("Card returns proper formatted value for hearts", func(t *testing.T) {
		heart := NewCard(2, Hearts)
		assert.True(t, heart != nil, "Card is not nil")
		fmt.Printf("\n***** %s *****\n", heart.GetSuitFormatted())
		assert.True(t, heart.GetSuitFormatted() == "♥", "Test that hearts returns formatted value of '♥'")
	})

	t.Run("Card returns proper formatted value for spades", func(t *testing.T) {
		spade := NewCard(2, Spades)
		assert.True(t, spade != nil, "Card is not nil")
		assert.True(t, spade.GetSuitFormatted() == "♠", "Test that hearts returns formatted value of '♠'")
	})

	t.Run("Card returns proper formatted value for diamonds", func(t *testing.T) {
		diamond := NewCard(2, Diamonds)
		assert.True(t, diamond != nil, "Card is not nil")
		assert.True(t, diamond.GetSuitFormatted() == "♦", "Test that hearts returns formatted value of '♦'")
	})

	t.Run("Card returns proper formatted value for clubs", func(t *testing.T) {
		club := NewCard(2, Clubs)
		assert.True(t, club != nil, "Card is not nil")
		assert.True(t, club.GetSuitFormatted() == "♣", "Test that hearts returns formatted value of '♣'")
	})
}
