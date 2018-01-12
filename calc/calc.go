package main

import (
	"fmt"
	"os"
	"strconv"
)

// Calculate a number by the following one, until there aren't remaining
func Calculate(init float64) {
	for i := 3; i < len(os.Args); i++ {
		number, _ := strconv.ParseFloat(os.Args[i], 32)
		switch os.Args[1] {
		case "*":
			init = init * number
		case "/":
			init = init / number
		case "+":
			init = init + number
		case "-":
			init = init - number
		}
	}
	fmt.Printf("%.2f\n", init)
}

func main() {
	init, _ := strconv.ParseFloat(os.Args[2], 32)

	Calculate(init)
}
