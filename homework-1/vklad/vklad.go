package vklad

import (
	"fmt"
)

// Vklad comment
func Vklad() {
	fmt.Println("Введите сумму вклада под 5%: ")
	var inputSum float64
	var outputSum float64
	fmt.Scanf("%f", &inputSum)
	for i := 1; i < 6; i++ {
		outputSum = (outputSum + outputSum*0.05)
		fmt.Println(outputSum)
	}

}
