package ovens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasvass/Pizza-Factory/ovens"
	"github.com/vasvass/Pizza-Factory/pizza"
)

func TestOoniCook(t *testing.T) {
	t.Run("cooks the pizza", func(t *testing.T) {
		o := ovens.Ooni{}
		p := pizza.Pizza{}

		cookedPizza := o.Cook(p)

		assert.True(t, cookedPizza.IsCooked)
	})
}
