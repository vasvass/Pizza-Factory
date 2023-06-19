package pizza

type ToppingCategory string

const (
	MeatTopping      ToppingCategory = "meat"
	VegetableTopping ToppingCategory = "vegetable"
	CheeseTopping    ToppingCategory = "cheese"
)

type Pizza struct {
	Dough    Dough
	Sauce    Sauce
	Toppings Toppings
}

type Sauce struct {
	Name   string
	Amount int
}

type Toppings []Topping

type Topping struct {
	Name     string
	Amount   int
	Category ToppingCategory
}

type Dough struct {
	FlourAmount     int
	WaterAmount     int
	SaltAmount      int
	YeastAmount     int
	ExtraIngredient ExtraDoughIngredients
}

type ExtraDoughIngredients []ExtraDoughIngredient

type ExtraDoughIngredient struct {
	Name   string
	Amount int
}
