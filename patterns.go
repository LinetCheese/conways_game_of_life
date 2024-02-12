package main

func glider() {
	grid[2][1] = true
	grid[2][2] = true
	grid[2][3] = true
	grid[1][3] = true
	grid[0][2] = true
}

func gliderGun() {
	grid[6][0] = true
	grid[7][0] = true
	grid[6][1] = true
	grid[7][1] = true

	grid[4][13] = true
	grid[4][12] = true
	grid[5][11] = true
	grid[6][10] = true
	grid[7][10] = true
	grid[8][10] = true
	grid[9][11] = true
	grid[10][12] = true
	grid[10][13] = true
	grid[7][14] = true
	grid[5][15] = true
	grid[9][15] = true
	grid[6][16] = true
	grid[7][16] = true
	grid[8][16] = true
	grid[7][17] = true

	grid[4][20] = true
	grid[5][20] = true
	grid[6][20] = true
	grid[4][21] = true
	grid[5][21] = true
	grid[6][21] = true
	grid[3][22]= true
	grid[7][22] = true
	grid[2][24] = true
	grid[3][24] = true
	grid[7][24] = true
	grid[8][24] = true

	grid[4][34] = true
	grid[5][34] = true
	grid[4][35] = true
	grid[5][35] = true
}
