package ovens

import (
	"github.com/vasvass/Pizza-Factory/pizza"
)

type Ooni struct{}

func (o Ooni) Cook(p pizza.Pizza) pizza.Pizza {
	p.IsCooked = true
	return p
}
