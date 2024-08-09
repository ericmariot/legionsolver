package main

import "fmt"

func main() {
	positions := [][]int{
		{1, 1},
	}
	pieces := []int{1, 2}

	grid := solver(positions, pieces)
	fmt.Println(grid)
}
