package abstract_factory

type IsportsFacotry interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

type Brand string

const (
	Brand_Nike   = "Nike"
	Brand_Adidas = "Adidas"
)

func GetSPortsFactory(brand Brand) (products IsportsFacotry) {
	if brand == Brand_Nike {
		products = &Nike{}
	}

	if brand == Brand_Adidas {
		products = &Adidas{}
	}

	return
}
