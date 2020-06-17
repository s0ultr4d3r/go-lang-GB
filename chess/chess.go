package chess

import (
	"fmt"
)

// KnightVars implements a miscalculation of the possible moves of a single-root chess piece
func KnightVars(x string, y int) {
	var mov1 = [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
	var mov2 = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	var buf = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var z int
	for n := 0; n < 7; n++ {
		if x == buf[n] {
			z = n + 1
		}
	}
	var a, b int
	var c string
	for i := 0; i <= 7; i++ {
		a = z + mov1[i]
		b = y + mov2[i]
		if a > 0 && b > 0 {
			c = buf[b-1]
			fmt.Println(x, y, "->", c, a)

		}
	}
}
