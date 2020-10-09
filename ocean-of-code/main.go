package main

import (
	"fmt"
	"math/rand"
)
import "os"
import "bufio"

type cell struct {
	island  bool
	visited bool
}

const MAX_CELLS = 15



func main() {
	rand.Seed(86)
	scanner, board := startInputs()

	x, y := searchStartingCell(board)
	fmt.Println(x, y)

	for {
		//displayBoard(board)
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)

		board[x][y].visited = true

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		opponentOrders := scanner.Text()
		fmt.Fprintln(os.Stderr, " op said ", opponentOrders)

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		direction := movables(x, board, y)

		size := len(direction)
		if size == 0 {
			fmt.Println("SURFACE")

			board = funcName(board)
		} else {
			fmt.Fprint(os.Stderr, "len  ", size)
			intn := rand.Intn(size)
			fmt.Fprint(os.Stderr, "rand ? ",intn)
			fmt.Println("MOVE", direction[intn], "TORPEDO", "| TORPEDO", x, y+3)
		}
	}
}


func movables(x int, board [15][15]cell, y int) []string {
	var direction []string

	if x+1 < MAX_CELLS && !board[x+1][y].visited && !board[x+1][y].island {
		//fmt.Fprint(os.Stderr, "going south")
		direction = append(direction, "E")
	}
	if x-1 >= 0 && !board[x-1][y].visited && !board[x-1][y].island {
		//fmt.Fprint(os.Stderr, "going north")
		direction = append(direction, "W")
	}
	if y+1 < MAX_CELLS && !board[x][y+1].visited && !board[x][y+1].island {
		//fmt.Fprint(os.Stderr, "going east")
		direction = append(direction, "S")
	}
	if y-1 >= 0 && !board[x][y-1].visited && !board[x][y-1].island {
		//fmt.Fprint(os.Stderr, "going west")
		direction = append(direction, "N")
	}
	return direction
}

func funcName(board [15][15]cell) [15][15]cell {
	for y := 0; y < MAX_CELLS; y++ {
		for x := 0; x < MAX_CELLS; x++ {
			board[x][y].visited = false
		}
	}
	return board
}



func startInputs() (*bufio.Scanner, [15][15]cell) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width, height, myId int
	var board [MAX_CELLS][MAX_CELLS]cell
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)

	for y := 0; y < height; y++ {
		scanner.Scan()
		line := scanner.Text()
		for x := 0; x < MAX_CELLS; x++ {
			if line[x] == 'x' {
				board[x][y] = cell{island: true}
			}
		}
	}
	return scanner, board
}

func searchStartingCell(board [15][15]cell) (int, int) {
	var x = 0
	var y = 0
	var startCell = board[x][y]
	for startCell.island == true {
		if x == MAX_CELLS {
			x = -1
			y++
		}
		x++
		startCell = board[x][y]
	}
	return x, y
}

func displayBoard(board [15][15]cell) {
	for y := 0; y < MAX_CELLS; y++ {
		for x := 0; x < MAX_CELLS; x++ {
			if board[x][y].island {

				fmt.Fprint(os.Stderr, "x")
			} else if board[x][y].visited {

				fmt.Fprint(os.Stderr, "o")
			} else {

				fmt.Fprint(os.Stderr, ".")
			}

		}
		fmt.Fprintln(os.Stderr, "")
	}
}
