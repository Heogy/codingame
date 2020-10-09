package main

import (
	"fmt"
	"math/rand"
	"strconv"
)
import "os"
import "bufio"

/**
 * Grab the pellets as fast as you can!
 **/

type pac struct {
	id int
	tx int
	ty int
}
type coord struct {
	x int
	y int
}

type cel struct {
	floor bool
	value int
	t     bool
	pacId int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: size of the grid
	// height: top left corner is (x=0, y=0)
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	var grid = make([][]cel, height)

	for i := 0; i < height; i++ {

		grid[i] = make([]cel, width)

		scanner.Scan()
		row := scanner.Text()

		for j := 0; j < width; j++ {
			if row[j:j+1] == " " {
				grid[i][j].floor = true
			} else {
				grid[i][j].floor = false
			}
			grid[i][j].value = 0
		}
	}

	for {

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				grid[i][j].value = 0
			}
		}

		var pacs []pac
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
		var visiblePacCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePacCount)

		for i := 0; i < visiblePacCount; i++ {
			var pacId int
			var mine bool
			var _mine int
			var x, y int
			var typeId string
			var speedTurnsLeft, abilityCooldown int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &pacId, &_mine, &x, &y, &typeId, &speedTurnsLeft, &abilityCooldown)
			mine = _mine != 0

			if mine {
				pacs = append(pacs, pac{pacId, x, y})
			}
		}
		fmt.Fprintln(os.Stderr, pacs)

		removePacTargetFromGrid(grid, pacs)

		var visiblePelletCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePelletCount)

		for i := 0; i < visiblePelletCount; i++ {
			var x, y, value int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &x, &y, &value)
			grid[y][x].value = value
		}
		//fmt.Fprintln(os.Stderr, "---------- round start -----------")
		//displayGrid(grid)
		//fmt.Fprintln(os.Stderr, "----------------------------------")

		var pacCels map[int]coord

		for _, pac := range pacs {
			fmt.Fprint(os.Stderr, pac)
			pacCels = getPacCels(grid, len(pacs))
			fmt.Fprintln(os.Stderr, "pacs len", len(pacCels))
			target := ihaveTarget(pacCels, pac)
			fmt.Fprint(os.Stderr, target)

			done := myTargetIsDone(pac, grid)
			fmt.Fprintln(os.Stderr, done)
			if !target || done {
				attributeAValidOne(grid, pac)
			}

			//fmt.Fprintln(os.Stderr, "---------- after", pac.id,"-----------")2
			//displayGrid(grid)
			//fmt.Fprintln(os.Stderr, "----------------------------------")

		}

		fmt.Fprintln(os.Stderr, "getPacCels")
		pacCels = getPacCels(grid, len(pacs))
		fmt.Fprintln(os.Stderr, "gen cmd")
		cmd := ""

		for i, c := range pacCels {
			cmd += "MOVE " + strconv.Itoa(i) + " " + strconv.Itoa(c.x) + " " + strconv.Itoa(c.y) + " to " + strconv.Itoa(c.x) + " " + strconv.Itoa(c.y) + " | "
		}

		//fmt.Fprintln(os.Stderr, cmd)
		fmt.Println(cmd[:len(cmd)-3]) // MOVE <pacId> <x> <y>
	}

}

func removePacTargetFromGrid(grid [][]cel, pacs []pac) {

	var pacIds = make(map[int]string, 6)
	for _, p := range pacs {
		pacIds[p.id] = ""
	}

	for y, row := range grid {
		for x, cel := range row {

			if cel.t {
				_, ok := pacIds[cel.pacId]
				if !ok {
					fmt.Fprintln(os.Stderr, "removing", cel.pacId, "x", x, "y", y)
					grid[y][x].t = false
					grid[y][x].pacId = -1
				}
			}
		}
	}

}

func displayGrid(grid [][]cel) {
	for _, rows := range grid {
		for _, c := range rows {
			if c.floor == false {
				fmt.Fprint(os.Stderr, "*")
			} else {
				fmt.Fprint(os.Stderr, "_")
			}

			fmt.Fprint(os.Stderr, c.value)
			if c.t == false {
				fmt.Fprint(os.Stderr, "t")
			} else {
				fmt.Fprint(os.Stderr, "o")
			}
			fmt.Fprint(os.Stderr, c.pacId)
		}
		fmt.Fprintln(os.Stderr, "")

	}
}

func attributeAValidOne(grid [][]cel, p pac) {
	fmt.Fprintln(os.Stderr, "attributing", p.id)

	x, y := randBullet(grid)
	fmt.Fprintln(os.Stderr, "rand done", p.id)

	rmOld(grid, p)
	attributeNew(grid, x, y, p)

	fmt.Fprintln(os.Stderr, "attributing", grid[y][x])

}

func attributeNew(grid [][]cel, x int, y int, p pac) {
	grid[y][x].t = true
	grid[y][x].pacId = p.id
}

func rmOld(grid [][]cel, p pac) {
	for y, row := range grid {
		for x, c := range row {
			if c.t == true && c.pacId == p.id {
				grid[y][x].t = false
			}
		}
	}
}

func myTargetIsDone(p pac, grid [][]cel) bool {
	for _, row := range grid {
		for _, c := range row {
			if c.t == true && c.pacId == p.id && c.value == 0 {
				return true
			}
		}
	}
	return false
}

func ihaveTarget(cels map[int]coord, p pac) bool {
	//fmt.Fprintln(os.Stderr, "target ", p.id)
	//fmt.Fprintln(os.Stderr, "map ", cels)
	//fmt.Fprintln(os.Stderr, "pac ", p)

	_, ok := cels[p.id]
	//fmt.Fprintln(os.Stderr, "return ", ok)
	return ok
}

func getPacCels(grid [][]cel, pacLen int) map[int]coord {

	var pacs = make(map[int]coord, pacLen)

	for y, cels := range grid {
		for x, cel := range cels {
			if cel.t {
				pacs[cel.pacId] = coord{x, y}
				fmt.Fprintln(os.Stderr, "found tcel", cel.pacId, "t", cel.t, "x", x, "y", y)
			}
		}
	}
	return pacs
}

func getBigBullet(grid [][]cel) []coord {

	var bbul []coord

	for y, cels := range grid {
		for x, cel := range cels {
			if cel.value == 10 {
				bbul = append(bbul, coord{x, y})
			}
		}
	}
	return bbul
}

func getBullet(grid [][]cel) []coord {

	var bbul []coord

	for y, cels := range grid {
		for x, cel := range cels {
			if cel.value > 1 {
				bbul = append(bbul, coord{x, y})
			}
		}
	}
	return bbul
}

func getABullet(grid [][]cel) (int, int) {
	for y, cels := range grid {
		for x, cel := range cels {
			if cel.value == 10 {
				return x, y
			}
		}
	}
	for y, cels := range grid {
		for x, cel := range cels {
			if cel.value == 1 {
				return x, y
			}
		}
	}
	return 0, 0
}

func randBullet(grid [][]cel) (int, int) {

	for true {
		x := rand.Intn(len(grid[0]))
		y := rand.Intn(len(grid))
		//fmt.Fprintln(os.Stderr, x, y, grid[y][x].value, grid[y][x].t)
		if grid[y][x].value >= 1 && grid[y][x].t == false {
			return x, y
		}
	}

	return 0, 0
}
