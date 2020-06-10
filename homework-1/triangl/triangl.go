package triangl

import (
	"fmt"
	"math"
)

const kattetA float64 = 10
const kattetB float64 = 10

//Square comment
func Square() {
	var trSquare float64
	trSquare = (kattetA * kattetB) / 2
	fmt.Println("Square is:", trSquare)
}

//Gipp comment
func Gipp() {
	var gippC float64
	gippC = math.Sqrt((math.Pow(kattetA, 2) + math.Pow(kattetB, 2)))
	fmt.Println("Гиппотенуза равна:", gippC)
	fmt.Println("Периметр равен:", (gippC + kattetB + kattetA))
}
