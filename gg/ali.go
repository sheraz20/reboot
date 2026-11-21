package main

import (
	"fmt"
)

func main() {
	var cn1 complex64 = complex(1+2, 0)
	var cn2 complex128 = complex(3*4, 0)
	fmt.Println("this is the complex numbers", cn1)
	fmt.Println("this is the complex numbers", cn2)
}
