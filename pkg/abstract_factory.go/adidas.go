package abstract_factory

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: Brand_Adidas,
			size: 41,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: Brand_Adidas,
			size: 12,
		},
	}
}
