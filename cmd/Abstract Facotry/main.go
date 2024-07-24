package main

import (
	"fmt"

	"github.com/GuilhermeVozniak/DesignPatterns/pkg/abstract_factory.go"
)

func main() {
	var (
		adidasFactory = abstract_factory.GetSPortsFactory(abstract_factory.Brand_Adidas)
		nikeFactory   = abstract_factory.GetSPortsFactory(abstract_factory.Brand_Nike)

		// Shoes
		nikeShoe   = nikeFactory.MakeShoe()
		adidasShoe = adidasFactory.MakeShoe()

		// Shirt
		nikeShirt   = nikeFactory.MakeShirt()
		AdidasShirt = adidasFactory.MakeShirt()
	)

	printDetailsShirt(nikeShirt)
	printDetailsShirt(AdidasShirt)

	printDetailsShoe(nikeShoe)
	printDetailsShoe(adidasShoe)
}

func printDetailsShirt(s abstract_factory.IShirt) {
	fmt.Printf("Logo: %v", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printDetailsShoe(s abstract_factory.IShoe) {
	fmt.Printf("Logo: %v", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
