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
