package main

import (
	"fmt"
)

func main() {
	myPlant := NewPowerPlant(Wind)
	fmt.Println(myPlant.GetPlantType())
}

func NewPowerPlant(plantType PlantType) *powerPlant {
	return &powerPlant{plantType}
}

type powerPlant struct {
	plantType PlantType
}

func (pp *powerPlant) GetPlantType() PlantType {
	return pp.plantType
}

type PlantType int

const (
	Hydro PlantType = iota
	Wind
	Solar
)
