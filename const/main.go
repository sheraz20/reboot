package main

import "fmt"

func main() {
	// constants
	const a = 100
	const b = 20.3
	const c = "Hello"
	fmt.Println("Constant a:", a)
	fmt.Println("Constant b:", b)
	fmt.Println("Constant c:", c)

	const PI float64 = 3.14
	fmt.Println("Constant PI:", PI)
	const lili = 6.022e23
	fmt.Println("Constant lili:", lili)

	const greeting = "Hello, earth!"
	const quote = "\"this is the quote!\" - a programmer"
	fmt.Println("Constant greeting:", greeting)
	fmt.Println("Constant quote:", quote)
}
