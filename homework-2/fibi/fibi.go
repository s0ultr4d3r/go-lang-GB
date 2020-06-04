package fibi

import (
	"fmt"
	"math/big"
)

// Fibi calc Fibonacci with Math.Big
func Fibi(n int) big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	for limit := 1; limit < n; limit++ {
		fmt.Println(limit, a)
		a.Add(a, b)
		a, b = b, a
	}

	return *a
}
