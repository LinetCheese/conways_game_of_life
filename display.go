package main

import "fmt"

func cliDisplay() {
	fmt.Print("\033[2J") // Clears terminal
	mapToChar := map[bool]string{
		true:  "██",
		false: "░░",
	}

	output := ""

	for i := range gridSize {
		for j := range gridSize {
			output += mapToChar[grid[i][j]]
		}
		output += "\n"
	}
	fmt.Print(output)
}