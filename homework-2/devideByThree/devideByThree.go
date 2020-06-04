package devideByThree

import "fmt"

// DevideByThree is for checking the number for a multiplicity of 3
func DevideByThree(i int) {

	if i%3 == 0 {
		fmt.Println(i, "кратно 3")
	} else {
		fmt.Println(i, "не кратно 3")
	}

}
