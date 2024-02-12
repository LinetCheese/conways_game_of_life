# Conway's Game of Life in Go

A stupid implementation of CGOL out of boredom. For more info on what is it:
https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

tl;dr:

Conway's Game of Life is a 2d cellular automaton, the evolution of which is
determined by only the initial input. Each cell can have only two states - 
alive or dead. On each cycle the state of the grid changes as follows:

- Any live cell with fewer than two live neighbors dies, as if by underpopulation.
- Any live cell with two or three live neighbors lives on to the next generation.
- Any live cell with more than three live neighbors dies, as if by overpopulation.
- Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Each tick these rules apply to all the cells at once and produce the new state
of the grid. Ideally the grid should be infinite but I couldn't be f-ed implementing
that, maybe later...

