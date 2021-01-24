package cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard(t *testing.T) {
	t.Run("Test for card equality will return true when equal", func(t *testing.T){
		fourHearts := NewCard(4, Hearts)
		test := NewCard( 4, Hearts)

		assert.True(t, fourHearts.Equals(*test), "Equals works when comparing [%s == %s]", fourHearts, test)
	})
}
