package converter

import "fmt"

const kurs float64 = 65

// Converter comment
func Converter() {
	fmt.Print("Enter your $: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := kurs * input
	fmt.Println("It's", output, "in rubles")
}
