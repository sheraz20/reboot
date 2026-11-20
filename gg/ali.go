package main

import (
	"fmt"
	"math"
)

func main() {
	// Signed integer types and their zero values.
	// var i int
	// var i8 int8
	// var i16 int16
	// var i32 int32
	// var i64 int64
	// fmt.Println("zero values of signed integers:", i, i8, i16, i32, i64)
	fmt.Println("int   :", math.MinInt, "to", math.MaxInt)
	fmt.Println("int8  :", math.MinInt8, "to", math.MaxInt8)
	fmt.Println("int16 :", math.MinInt16, "to", math.MaxInt16)
	fmt.Println("int32 :", math.MinInt32, "to", math.MaxInt32)
	fmt.Println("int64 :", math.MinInt64, "to", math.MaxInt64)
}
