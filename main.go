package main

import (
	"time"
)

const gridSize = 50
const cycleLength = 50 * time.Millisecond

var grid [gridSize][gridSize]bool

func init() {
	// Input into CGL
	gliderGun()
}

func newCell(i int, j int) bool {
	// - Any live cell with fewer than two live neighbors dies, as if by underpopulation.
	// - Any live cell with two or three live neighbors lives on to the next generation.
	// - Any live cell with more than three live neighbors dies, as if by overpopulation.
	// - Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	aliveCount := 0

	for nI := i - 1; nI <= i+1; nI++ {
		for nJ := j - 1; nJ <= j+1; nJ++ {
			if nI == i && nJ == j {
				continue
			}
			if nI < 0 || nJ < 0 || nI >= gridSize || nJ >= gridSize {
				continue
			}
			if grid[nI][nJ] {
				aliveCount += 1
			}
		}
	}

	if grid[i][j] {
		if aliveCount == 2 || aliveCount == 3 {
			return true
		}
		return false
	}
	if aliveCount == 3 {
		return true
	}
	return false
}

func main() {
	for {
		var newCanvas [gridSize][gridSize]bool

		for i := range gridSize {
			for j := range gridSize {
				newCanvas[i][j] = newCell(i, j)
			}
		}
		grid = newCanvas
		cliDisplay()
		time.Sleep(cycleLength)
	}
}
