package main

import (
	"fmt"
)

const max = 140

func main() {
	fmt.Println(max)
	score := 80
	bonus := 4.5
	total := float64(score) + bonus

	var letter rune = 'G'
	var nextletter = letter + 1
	fmt.Println(nextletter)
	fmt.Println(string(nextletter))

	fmt.Println(total)

}
