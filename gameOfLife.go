package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"os"
	"time"
	"runtime"
)

type cell struct {
	id          int
	alive       bool
	connections [8]*cell
	x           int
	y           int
}
const speedMiliseconds = 100
const xCells = 100

const yCells = 100

var liveList = []*cell{};

func CallClear() {
	clear := make(map[string]string)
	clear["linux"] = "clear"
	clear["darwin"] = "clear"
	clear["windows"] = "clr"

	value, ok := clear[runtime.GOOS]
	if ok {
		cmd := exec.Command(value)
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
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

	var inList = func(x int, list []int) bool {
		for _, i := range list {
			if i == x {
				return true;
			}
		}

		return false;
	}
	inList(1,[]int{0})

	var id = 0
	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			id += 1
			board[x][y] = cell{
				id: id,
				alive:rand.Intn(2) == 1,
				//alive: inList(id, []int{2, xCells + 3, (xCells * 2) + 1, (xCells * 2) + 2, (xCells * 2) + 3}),
				x:     x,
				y:     y,
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

		}
	}

	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			if (board[x][y].alive) {
				addToLiveList(board[x][y])
			}
		}
	}


	for range time.Tick(speedMiliseconds * time.Millisecond) {
			board = updateBoard(board)
			printBoard(board)

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

	liveList = []*cell{}
	for x := 0; x < xCells; x++ {
		for y := 0; y < yCells; y++ {
			if aliveState[x][y] {
				addToLiveList(board[x][y])
				board[x][y].alive = true
			} else {
				board[x][y].alive = false
			}
		}
	}

	return board
}

func isSetAlready(id int, list []*cell) bool {
	for _, cell := range list {
		if (cell.id == id) {
			return true
		}
	}

	return false
}

func addToLiveList(cell cell) {

	if (false == isSetAlready(cell.id, liveList)) {
		liveList = append(liveList, &cell)
	}
	for _, c := range cell.connections {
		if (false == isSetAlready(c.id, liveList)) {
			liveList = append(liveList, c)
		}
	}
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

	//for _, x := range liveList {
	//	fmt.Printf("%d, ", x.id)
	//}
	fmt.Println("\n");
}
