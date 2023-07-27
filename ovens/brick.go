package ovens

import "github.com/vasvass/Pizza-Factory/pizza"

type Brick struct{}

func (b Brick) Cook(p pizza.Pizza) pizza.Pizza {
	p.IsCooked = true
	return p
}
