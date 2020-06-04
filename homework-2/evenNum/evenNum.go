package evenNum

import "fmt"

// EvenNum is for checking the number for a multiplicity of 2
func EvenNum(i int) {

	if i%2 == 0 {
		fmt.Println(i, " - является чётным.")
	} else {
		fmt.Println(i, " - является нечетным")
	}

}
