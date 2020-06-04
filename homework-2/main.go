package main

import (
	"devideByThree"
	"evenNum"
	"fibi"
	"fmt"
)

func main() {
	var i int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &i)
	evenNum.EvenNum(i)
	devideByThree.DevideByThree(i)
	fibi.Fibi(100)
}
