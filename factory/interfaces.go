package factory

import "github.com/vasvass/Pizza-Factory/pizza"

type PizzaOven interface {
	Cook(pizza pizza.Pizza) pizza.Pizza
}
