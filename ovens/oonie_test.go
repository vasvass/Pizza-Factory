package ovens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasvass/Pizza-Factory/ovens"
	"github.com/vasvass/Pizza-Factory/pizza"
	timeUtils "github.com/vasvass/Pizza-Factory/utils"
)

func TestOoniCook(t *testing.T) {
	t.Run("cooks the pizza", func(t *testing.T) {
		o := ovens.Ooni{}
		p := pizza.Pizza{}

		cookedPizza := o.Cook(p)

		assert.True(t, cookedPizza.IsCooked)
	})
}

func TestOoniCookingTime(t *testing.T) {
	testCases := []struct {
		desc            string
		temperature     int
		temperatureUnit ovens.TemperatureUnit
		expectedTime    timeUtils.Seconds
	}{
		{
			desc:            "450 degrees celcius should take 60 seconds",
			temperature:     450,
			temperatureUnit: ovens.TemperatureUnitCelsius,
			expectedTime:    60,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			o := ovens.Ooni{
				Temperature:     tC.temperature,
				TemperatureUnit: tC.temperatureUnit,
			}

			cookingTime := o.CookingTime()

			assert.Equal(t, tC.expectedTime, cookingTime)
		})
	}
}
