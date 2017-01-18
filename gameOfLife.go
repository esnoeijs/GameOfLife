package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"os"
	"time"
)

type cell struct {
	id int
	alive       bool
	connections [8]*cell
	x int
	y int
}

const xCells = 5

const yCells = 5

var liveList = []*cell{};

func CallClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func isSetAlready (id int, list []*cell) bool {
	for _, cell := range list {
		if (cell.id == id) {
			return true
		}
	}

	return false
}

func main() {
	rand.Seed(126)
	board := [xCells][yCells]cell{}

	var wrapCoordinate = func(x int, max int) int {
		if x < 0 {
			return max
		} else if x > max {
			return 0
		} else {
			return x
		}
	}


	var inList = func (x int, list []int) bool {
		for _, i := range list {
			if i == x {
				return true;
			}
		}

		return false;
	}
	var id = 0
	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			id += 1
			ints := []int{2, 8, 11, 12, 13}
			board[x][y] = cell{
				id: id,
				//alive:rand.Intn(2) == 1,
				alive: inList(id, ints),
				x: x,
				y: y,
			}

			xp := wrapCoordinate(x-1, xCells-1)
			xn := wrapCoordinate(x+1, xCells-1)
			yp := wrapCoordinate(y-1, yCells-1)
			yn := wrapCoordinate(y+1, yCells-1)

			var xyMap = [8][2]int{
				{xp, yp}, // NW
				{xp, y},  // N
				{xp, yn}, // NE
				{x, yp},  // W
				{x, yn},  // E
				{xn, yp}, // SW
				{xn, y},  // S
				{xn, yn}, // SE
			}

			for key, xy := range xyMap {
				board[x][y].connections[key] = &board[xy[0]][xy[1]]
			}

			if (board[x][y].alive) {
				liveList = append(liveList, &board[x][y])
				for _,c := range board[x][y].connections {
					if (false == isSetAlready(c.id, liveList)) {
						liveList = append(liveList, c)
					}
				}

			}
		}
	}


	var i = 0;
	for i < 1000 {
		i += 1

		printBoard(board)
		board = updateBoard(board)
		time.Sleep(1000 * time.Millisecond)
	}
}

func updateBoard(board [xCells][yCells]cell) [xCells][yCells]cell {
	aliveState := [xCells][yCells]bool{}



	for _, cell := range liveList {
		aliveScore := 0
		for n := 0; n < 8; n++ {
			if cell.connections[n].alive {
				aliveScore += 1
			}
		}

		if aliveScore < 2 || aliveScore > 3 {
			aliveState[cell.x][cell.y] = false
		} else if aliveScore == 3 {
			aliveState[cell.x][cell.y] = true
		} else {
			aliveState[cell.x][cell.y] = cell.alive
		}
	}
	//
	//for x := 0; x < xCells; x++ {
	//	for y := 0; y < yCells; y++ {
	//		aliveScore := 0
	//		for n := 0; n < 8; n++ {
	//			if board[x][y].connections[n].alive {
	//				aliveScore += 1
	//			}
	//		}
	//
	//		if aliveScore < 2 || aliveScore > 3 {
	//			aliveState[x][y] = false
	//		} else if aliveScore == 3 {
	//			aliveState[x][y] = true
	//		} else {
	//			aliveState[x][y] = board[x][y].alive
	//		}
	//	}
	//}

	liveList = []*cell{}
	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			if aliveState[x][y] {
				if (false == isSetAlready(board[x][y].id, liveList)) {
					liveList = append(liveList, &board[x][y])
				}
				for _,c := range board[x][y].connections {
					if (false == isSetAlready(c.id, liveList)) {
						liveList = append(liveList, c)
					}
				}
				board[x][y].alive = true
			} else {
				board[x][y].alive = false
			}
		}
	}

	return board
}

func printBoard(board [xCells][yCells]cell) {
	CallClear()

	fmt.Print("\n")
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
	fmt.Printf("liveList: %d\n", len(liveList))

	for _, x := range liveList {
		fmt.Printf("%d, ", x.id)
	}
	fmt.Println("\n");
}
