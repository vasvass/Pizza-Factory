package ovens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasvass/Pizza-Factory/ovens"
	"github.com/vasvass/Pizza-Factory/pizza"
	timeUtils "github.com/vasvass/Pizza-Factory/utils"
)

func TestBrickCook(t *testing.T) {
	t.Run("cooks the pizza", func(t *testing.T) {
		b := ovens.Brick{}
		p := pizza.Pizza{}

		cookedPizza := b.Cook(p)

		assert.True(t, cookedPizza.IsCooked)
	})
}

func TestBrickCookingTime(t *testing.T) {
	testCases := []struct {
		desc            string
		temperature     int
		temperatureUnit ovens.TemperatureUnit
		expectedTime    timeUtils.Seconds
	}{
		{
			desc:            "450 degrees celcius should take 71 seconds",
			temperature:     450,
			temperatureUnit: ovens.TemperatureUnitCelsius,
			expectedTime:    71,
		},
		{
			desc:            "400 degrees celcius should take 80 seconds",
			temperature:     400,
			temperatureUnit: ovens.TemperatureUnitCelsius,
			expectedTime:    80,
		},
		{
			desc:            "300 degrees celcius should take 106 seconds",
			temperature:     300,
			temperatureUnit: ovens.TemperatureUnitCelsius,
			expectedTime:    106,
		},
		{
			desc:            "500 degrees celcius should take 64 seconds",
			temperature:     500,
			temperatureUnit: ovens.TemperatureUnitCelsius,
			expectedTime:    64,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b := ovens.Brick{
				Temperature:     tC.temperature,
				TemperatureUnit: tC.temperatureUnit,
			}

			cookingTime := b.CookingTime()

			assert.Equal(t, tC.expectedTime, cookingTime)
		})
	}
}
