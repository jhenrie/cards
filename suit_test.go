package cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuit(t *testing.T) {
	t.Run("should confirm that the const Hearts is 0", func(t *testing.T){
		assert.True(t, Hearts == 0, "Confirm hearts const = 0 [Hearts %d]", Hearts)
	})

	t.Run("should create a suit of heart", func(t *testing.T) {
		heart := NewSuit(0)
		assert.True(t, heart == Hearts, "Confirm the suite is a heart [heart(%d) == Hearts(%d)]", heart, Hearts)
	})

	t.Run("should confirm that the const Diamonds is 1", func(t *testing.T){
		assert.True(t, Diamonds == 1, "Confirm diamonds const = 1 [Diamonds %d]", Diamonds)
	})

	t.Run("should create a suit of diamonds", func(t *testing.T) {
		diamond := NewSuit(1)
		assert.True(t, diamond == Diamonds, "Confirm the suite is a diamonds [diamonds(%d) == Diamonds(%d)]", diamond, Diamonds)
	})

	t.Run("should confirm that the const Clubs is 2", func(t *testing.T){
		assert.True(t, Clubs == 2, "Confirm clubs const = 2 [Clubs %d]", Clubs)
	})

	t.Run("should create a suit of clubs", func(t *testing.T) {
		club := NewSuit(2)
		assert.True(t, club == Clubs, "Confirm the suite is a clubs [club(%d) == Clubs(%d)]", club, Clubs)
	})

	t.Run("should confirm that the const Spades is 3", func(t *testing.T){
		assert.True(t, Spades == 3, "Confirm spades const = 3 [Spades %d]", Diamonds)
	})

	t.Run("should create a suit of spades", func(t *testing.T) {
		spade := NewSuit(3)
		assert.True(t, spade == Spades, "Confirm the suite is a spades [spade(%d) == Spades(%d)]", spade, Spades)
	})
}
