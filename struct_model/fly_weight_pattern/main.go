package main

import (
	"fly_weight_pattern/fly_weight"
	"fmt"
)

func main() {

	game := fly_weight.NewGame()

	//Add Terrorist
	game.AddTerrorist(fly_weight.TerroristDressType)
	game.AddTerrorist(fly_weight.TerroristDressType)
	game.AddTerrorist(fly_weight.TerroristDressType)
	game.AddTerrorist(fly_weight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(fly_weight.CounterTerrroristDressType)
	game.AddCounterTerrorist(fly_weight.CounterTerrroristDressType)
	game.AddCounterTerrorist(fly_weight.CounterTerrroristDressType)

	dressFactoryInstance := fly_weight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.GetDressMap() {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
