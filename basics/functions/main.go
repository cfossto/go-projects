package main

func yCoordinate(x int, k int, m int) int {

	var y = x*k+m;

	return y
}


func main(){

	var result int =yCoordinate(-5, 7,4)

	println(result)
}