package ovens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasvass/Pizza-Factory/ovens"
	"github.com/vasvass/Pizza-Factory/pizza"
)

func TestBrickCook(t *testing.T) {
	t.Run("cooks the pizza", func(t *testing.T) {
		b := ovens.Brick{}
		p := pizza.Pizza{}

		cookedPizza := b.Cook(p)

		assert.True(t, cookedPizza.IsCooked)
	})
}
