package main

import (
	"math"
)

func yCoordinate(x int, k int, m int) int {
	var y = x*k+m;
	return y
}



func power(x float64, y float64 ) float64 {
	result := math.Pow(x,y)
	return float64(result)
}


func main(){

	var result int =yCoordinate(-5, 7,4)
	println(result)

	var result2 float64 = power(4,2)
	println(int(result2))
}