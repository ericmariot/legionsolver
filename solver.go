package main

func solver(selection [][]int, pieces []int) [4][4]int {
	g := emptyGrid()
	g = setPositions(g, selection)
	g = setPieces(g, pieces)

	return g
}

func emptyGrid() [4][4]int {
	return [4][4]int{}
}

func setPositions(g [4][4]int, positions [][]int) [4][4]int {
	for _, pos := range positions {
		x, y := pos[0], pos[1]
		g[x][y] = 1
	}

	return g
}
func setPieces(grid [4][4]int, sizes []int) [4][4]int {
	for _, size := range sizes {
		placed := false
		for i := range grid {
			for j := 0; j <= len(grid[i])-size; j++ {
				isSequence := true
				for k := 0; k < size; k++ {
					if grid[i][j+k] != 1 {
						isSequence = false
						break
					}
				}
				if isSequence {
					for k := 0; k < size; k++ {
						grid[i][j+k] = 2
					}
					placed = true
					break
				}
			}
			if placed {
				break
			}
		}
	}
	return grid
}
