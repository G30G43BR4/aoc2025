package main

import (
	"fmt"
)

func main() {
	dial := 50

	distance := 68

	test := (distance % 100)
	dial = (dial - (distance % 100))
	fmt.Println(dial)
	if dial < 0 {
		dial = 100 + dial
	}

	fmt.Println(test)
	fmt.Println(dial)
}
