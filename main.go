package main

import "autos"

func main() {
	vw := autos.Car{
		Model: &autos.Model{
			Label: "Golf",
			Year:  2010,
		},
		Racing: &autos.Racing{
			HorsePwr: 200,
			MaxSpeed: 250,
		},
		Comfort: &autos.Comfort{
			AirCond:   true,
			Autopilot: false,
		},
		Color:  "white",
		Volume: 500,
	}

	toyota := autos.Car{
		Model: &autos.Model{
			Label: "camry",
			Year:  2012,
		},
		Racing: &autos.Racing{
			HorsePwr: 180,
			MaxSpeed: 250,
		},
		Comfort: &autos.Comfort{
			AirCond:   false,
			Autopilot: false,
		},
		Color:  "black",
		Volume: 555,
	}

	vw.ShowInfo()
	toyota.ShowInfo()
}
