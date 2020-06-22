package main

import (
	"fmt"
	"os"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(50 * time.Millisecond)

	cancel := make(chan int)

	go func() {
		fmt.Print(" Hit enter for cancel")
		os.Stdin.Read(make([]byte, 1))
	}()

	naturals := make(chan int)
	squares := make(chan int)
	c := <-cancel
	if c > 0 {
		os.Exit(4)
	}
	// генерация
	go func() {
		for x := 0; x < 20; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// печать
	for x := range squares {
		fmt.Println(x)
	}
	<-time.After(time.Second * 15)

}
