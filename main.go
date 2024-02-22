package main

import (
	"fmt"
	"time"
)

const SIZE = 25

type Grid [SIZE][SIZE]int

func main() {
	var grid Grid
	seed(&grid)
	PrintGrid(CalcNextGrid(grid))
	fmt.Println()
	for {
		grid = CalcNextGrid(grid)
		PrintGrid(grid)
		time.Sleep(1 * time.Second / 4)
	}

}

func seed(grid *Grid) {
	start := len(grid) / 2
	// create a glider, looks like an L-
	grid[start][start+1] = 1
	grid[start+1][start+2] = 1
	grid[start+2][start] = 1
	grid[start+2][start+1] = 1
	grid[start+2][start+2] = 1
}

func CalcNextGrid(grid Grid) Grid {
	var newGrid Grid
	size := len(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			neighbours := NeighbourCount(i, j, grid)
			if neighbours < 2 || neighbours > 3 {
				newGrid[i][j] = 0
			} else if neighbours == 3 {
				newGrid[i][j] = 1
			} else {
				newGrid[i][j] = grid[i][j]
			}
		}
	}
	return newGrid
}

func NeighbourCount(x int, y int, grid Grid) int {
	var count int = 0
	size := len(grid)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			rowIndex := (x + i + size) % size
			colIndex := (y + j + size) % size

			count += grid[rowIndex][colIndex]
		}
	}
	return count
}

func PrintGrid(grid Grid) {
	for _, row := range grid {
		for _, item := range row {
			if item == 1 {
				fmt.Printf("0")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
