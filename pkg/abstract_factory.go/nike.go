package abstract_factory

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: Brand_Nike,
			size: 41,
		},
	}
}

func (a *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: Brand_Nike,
			size: 12,
		},
	}
}
