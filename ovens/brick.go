package ovens

import (
	"math"
	"time"

	"github.com/vasvass/Pizza-Factory/pizza"
	timeUtils "github.com/vasvass/Pizza-Factory/utils"
)

// Cooking: 80 seconds under 400 degrees celsius
const kFactorBrick = 32000

type Brick struct {
	Temperature     int
	TemperatureUnit TemperatureUnit

	Sleeper Sleeper
}

func (b Brick) Cook(p pizza.Pizza) pizza.Pizza {
	p.IsCooked = true
	b.Sleeper.Sleep(time.Duration(b.CookingTime().ToNanoSeconds()))
	return p
}

func (b Brick) CookingTime() timeUtils.Seconds {
	if b.Temperature == 0 {
		return timeUtils.Seconds(math.MaxInt)
	}
	return timeUtils.Seconds(kFactorBrick / b.Temperature)
}
