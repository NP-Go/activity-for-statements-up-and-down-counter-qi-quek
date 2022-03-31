package main

import "fmt"

var numberBig, numberSmall int

func main() {
	fmt.Println("Please input two numbers")
	fmt.Println("Input first number")
	fmt.Scan(&numberBig)
	fmt.Println("Input second number")
	fmt.Scan(&numberSmall)

	if numberBig < numberSmall {
		numberBig, numberSmall = numberSmall, numberBig
	}

	for noAscend := numberSmall; noAscend <= numberBig; noAscend++ {
		fmt.Println(noAscend)
	}
	// status = false
	for noDescedng := numberBig; noDescedng >= numberSmall; noDescedng-- {
		fmt.Println(noDescedng)
	}

}
