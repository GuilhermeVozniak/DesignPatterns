package main

import (
	"fmt"

	"github.com/GuilhermeVozniak/DesignPatterns/pkg/factory"
)

func main() {
	ak47 := factory.GetGun(factory.GunType_AK47)
	musket := factory.GetGun(factory.GunType_Musket)

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
