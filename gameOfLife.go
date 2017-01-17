package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"os"
	"time"
)

type cell struct {
	alive       bool
	connections [8]*cell
}

const xCells = 10

const yCells = 20

func CallClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	rand.Seed(125)
	board := [xCells][yCells]cell{}


	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			board[x][y] = cell{alive:rand.Intn(2) == 1}

			var wrapCoordinate = func(x int, max int) int {
				if x < 0 {
					return max
				} else if x > max {
					return 0
				} else {
					return x
				}
			}

			xp := wrapCoordinate(x-1, xCells-1)
			xn := wrapCoordinate(x+1, xCells-1)
			yp := wrapCoordinate(y-1, yCells-1)
			yn := wrapCoordinate(y+1, yCells-1)

			board[x][y].connections[0] = &board[x][yp]
			board[x][y].connections[1] = &board[xn][yp]
			board[x][y].connections[2] = &board[xn][y]
			board[x][y].connections[3] = &board[xn][yn]
			board[x][y].connections[4] = &board[x][yn]
			board[x][y].connections[5] = &board[xp][yn]
			board[x][y].connections[6] = &board[xp][y]
			board[x][y].connections[7] = &board[xp][yp]
		}
	}

	for {
		time.Sleep(100 * time.Millisecond)
		printBoard(board)
		board = updateBoard(board)
	}
}

func updateBoard(board [xCells][yCells]cell) [xCells][yCells]cell {
	tmpBoard := [xCells][yCells]cell{}

	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			aliveScore := 0
			for n := 0; n < 8; n++ {
				if board[x][y].connections[n].alive {
					aliveScore += 1
				}
			}

			if aliveScore < 2 || aliveScore > 3 {
				tmpBoard[x][y].alive = false
			} else if aliveScore == 3 {
				tmpBoard[x][y].alive = true
			} else {
				tmpBoard[x][y].alive = board[x][y].alive
			}
		}
	}

	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			board[x][y].alive = tmpBoard[x][y].alive
		}
	}

	return board
}

func printBoard(board [xCells][yCells]cell) {
	CallClear()

	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			if board[x][y].alive {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
