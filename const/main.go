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
	// break line
	const Bell = "Bell is \a"
	fmt.Println(Bell)
	const LB = "I\nam\ngood\nguy\n!"
	fmt.Println(LB)

	const bio = "this is sheraz " + "he is trying to get into rebootcamp" + "and thats's way he is preparing so hard"
	fmt.Println(bio)
	const sherry = "age is 22" + " and he is a good guy" + " and he is a good programmer"
	fmt.Println(sherry)
	// break line
	const active = true
	const ready = false
	fmt.Println("Constant active:", active)
	fmt.Println("Constant ready:", ready)
}
