package autos

import "fmt"

// Model contains label and year.
type Model struct {
	Label string
	Year  int
}

// Racing contains HP and MaxSpeed.
type Racing struct {
	HorsePwr int
	MaxSpeed int
}

// Comfort contains aircond and autopilot (bool types).
type Comfort struct {
	AirCond   bool
	Autopilot bool // Ругается, что это есть на pkg.go.dev хотя я же не импортирую оттуда ничего
}

// Car contains all and color+volume.
type Car struct {
	Model   *Model
	Racing  *Racing
	Comfort *Comfort
	Color   string
	Volume  int
}

// ShowInfo show all information about car.
func (Car Car) ShowInfo() { // Ругается, что это есть на pkg.go.dev хотя я же не импортирую оттуда ничего
	fmt.Printf(
		"It's %s %d year has %d HP and rich %d maxspeed. \n Aircond and autopilote in this model: %v %v \n It's %s color and has %d volume\n",
		Car.Model.Label,
		Car.Model.Year,
		Car.Racing.HorsePwr,
		Car.Racing.MaxSpeed,
		Car.Comfort.AirCond,
		Car.Comfort.Autopilot,
		Car.Color,
		Car.Volume)

}
