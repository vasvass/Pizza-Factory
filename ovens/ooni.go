package ovens

import (
	"math"

	"github.com/vasvass/Pizza-Factory/pizza"
	timeUtils "github.com/vasvass/Pizza-Factory/utils"
)

// kFactor is the constant that defines the relationship between temperature and cooking time
const kFactorOoni = 27000

type Ooni struct {
	Temperature     int
	TemperatureUnit TemperatureUnit
}

func (o Ooni) Cook(p pizza.Pizza) pizza.Pizza {
	p.IsCooked = true
	o.CookingTime()
	return p
}

// we assume that the relationship between temperature and cooking time is linear. For simplicity.
func (o Ooni) CookingTime() timeUtils.Seconds {
	if o.Temperature == 0 {
		return timeUtils.Seconds(math.MaxInt)
	}
	return timeUtils.Seconds(kFactorOoni / o.Temperature)
}
