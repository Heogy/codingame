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
const (
	ROCK     = "ROCK"
	PAPER    = "PAPER"
	SCISSORS = "SCISSORS"
	MOVE     = "MOVE"
	SWITCH   = "SWITCH"
	SPEED    = "SPEED"
	KILL     = "KILL"
	SURVIVE  = "SURVIVE"
	BIGMIAM  = "BIGMIAM"
	MIAM     = "MIAM"
	NOT_SURE = "NOT_SURE"
	NO_OBJ   = "NO_OBJ"
)

type pac struct {
	id              int
	rpc             string
	speedTurnsLeft  int
	abilityCooldown int
}

//type coord struct {
//	x int
//	y int
//}

type cel struct {
	floor bool
	value int

	my    bool
	myPac pac

	his    bool
	hisPac pac
}

type cmd struct {
	action string
	typ    string
	x      int
	y      int
	desc   string
	intent string
}

type coord struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	grid := initial(scanner)

	const MaxRound = 200

	var pacRandomTarget = make(map[int]coord, 5)

	for i := 0; i < 5; i++ {
		pacRandomTarget[i] = randCoord(grid)
	}
	//todo rand init

	for r := 0; r < MaxRound; r++ {

		myPacCount := rundInit(scanner, grid)
		displayGrid(grid)

		var cmds = make(map[int]cmd, len(myPacCount))

		for _, pacId := range myPacCount {

			cmds[pacId] = findCmd(pacId, grid)

			if cmds[pacId].intent == NO_OBJ {
				x, y, _ := findPacCoord(pacId, grid)
				if pacRandomTarget[pacId].y == y && pacRandomTarget[pacId].x == x {
					pacRandomTarget[pacId] = randCoord(grid)
				}
				cmds[pacId] = cmd{
					action: MOVE,
					typ:    "",
					x:      pacRandomTarget[pacId].x,
					y:      pacRandomTarget[pacId].y,
					desc:   "discover",
					intent: NO_OBJ,
				}
			}

		}

		strCmd := ""
		for pacId, cmd := range cmds {

			if cmd.action == MOVE {
				strCmd += "MOVE " + strconv.Itoa(pacId) + " " + strconv.Itoa(cmd.x) + " " + strconv.Itoa(cmd.y) + " " + cmd.desc + " | "
			} else if cmd.action == SWITCH {
				strCmd += "SWITCH " + strconv.Itoa(pacId) + " " + cmd.typ + " " + cmd.desc + " | "
			} else if cmd.action == SPEED {
				strCmd += "SPEED " + strconv.Itoa(pacId) + " " + cmd.desc + " | "
			}

		}
		fmt.Println(strCmd[:len(strCmd)-3])
	}

}

func randCoord(grid [][]cel) coord {
	height := len(grid)
	width := len(grid[0])
	var a = height * width
	var i = 0
	for i < a {
		x := rand.Intn(width)
		y := rand.Intn(height)
		if grid[y][x].floor {
			return coord{x, y}
		}
	}
	fmt.Fprintln(os.Stderr, "NOWHERE TO GO !!!!!!!")
	return coord{12, 12}
}

func findCmd(id int, grid [][]cel) cmd {

	x, y, pac := findPacCoord(id, grid)

	//if pac.abilityCooldown == 0 && pac.speedTurnsLeft ==0 {
	//	return cmd{
	//		action: SPEED,
	//		typ:    "",
	//		x:      0,
	//		y:      0,
	//		desc:   "SPEED",
	//		intent: "SPEED",
	//	}
	//}



	cmds := lookOne(x, y, grid, pac.rpc, id)

	i := 2
	for i < 5 {
		//fmt.Fprintln(os.Stderr, id, "look", i), i
		cmds = append(cmds, look(x, y, grid, pac.rpc, i, id)...)
		i++
	}

	return pickTheBest(cmds)

}

func look(x int, y int, grid [][]cel, rpc string, i int, id int) []cmd {
	//fmt.Fprint(os.Stderr, " > ", i)
	height := len(grid)
	width := len(grid[0])

	if i == 0 {
		fmt.Fprintln(os.Stderr, "look error")
		return nil
	} else if i == 1 {
		cmds := lookOne((x+1)%width, y, grid, rpc, id)
		cmds = lookOne((x-1+width)%width, y, grid, rpc, id)
		cmds = lookOne(x, (y+1)%height, grid, rpc, id)
		cmds = lookOne(x, (y-1+height)%height, grid, rpc, id)

		return cmds
	} else {
		cmds := look((x+1)%width, y, grid, rpc, i-1, id)
		cmds = look((x-1+width)%width, y, grid, rpc, i-1, id)
		cmds = look(x, (y+1)%height, grid, rpc, i-1, id)
		cmds = look(x, (y-1+height)%height, grid, rpc, i-1, id)
		return cmds
	}
}

func pickTheBest(cmds []cmd) cmd {

	var killCmds []cmd
	var surviveCmds []cmd
	var bigMiamCmds []cmd
	var miamCmds []cmd
	var notSureCmds []cmd

	for _, cmd := range cmds {
		if cmd.intent == KILL {
			fmt.Fprintln(os.Stderr, "KILL")
			killCmds = append(killCmds, cmd)
		} else if cmd.intent == SURVIVE {
			fmt.Fprintln(os.Stderr, "SURVIVE")
			surviveCmds = append(surviveCmds, cmd)
		} else if cmd.intent == NOT_SURE {
			fmt.Fprintln(os.Stderr, "NOT_SURE")
			notSureCmds = append(notSureCmds, cmd)
		} else if cmd.intent == BIGMIAM {
			bigMiamCmds = append(bigMiamCmds, cmd)
		} else if cmd.intent == MIAM {
			miamCmds = append(miamCmds, cmd)
		} else if cmd.intent == NOT_SURE {
			fmt.Fprintln(os.Stderr, "look error")
			notSureCmds = append(notSureCmds, cmd)
		}
	}

	if len(killCmds) > 0 {
		return killCmds[0]
	} else if len(surviveCmds) > 0 {
		return surviveCmds[0]
	} else if len(bigMiamCmds) > 0 {
		return bigMiamCmds[0]
	} else if len(miamCmds) > 0 {
		return miamCmds[0]
	} else if len(notSureCmds) > 0 {
		return notSureCmds[0]
	} else {
		return cmd{
			action: "",
			typ:    "",
			x:      0,
			y:      0,
			desc:   "",
			intent: NO_OBJ,
		}
	}

}

func lookOne(x int, y int, grid [][]cel, rpc string, id int) []cmd {
	var cmds []cmd
	height := len(grid)
	width := len(grid[0])

	good, celCmd := genCaseCmd((x+1)%width, y, grid, rpc, id)
	if good {
		cmds = append(cmds, celCmd)
	}

	good, celCmd = genCaseCmd((x-1+width)%width, y, grid, rpc, id)
	if good {
		cmds = append(cmds, celCmd)
	}

	good, celCmd = genCaseCmd(x, (y+1)%height, grid, rpc, id)
	if good {
		cmds = append(cmds, celCmd)
	}

	good, celCmd = genCaseCmd(x, (y-1+height)%height, grid, rpc, id)
	if good {
		cmds = append(cmds, celCmd)
	}

	return cmds

}

func genCaseCmd(x int, y int, grid [][]cel, rpc string, id int) (bool, cmd) {

	if grid[y][x].floor == false {
		return false, cmd{}
	} else if grid[y][x].value == 10 {
		return true, cmd{
			action: MOVE,
			typ:    "",
			x:      x,
			y:      y,
			desc:   "BIG BIG",
			intent: BIGMIAM,
		}
	} else if grid[y][x].value == 1 {
		return true, cmd{
			action: MOVE,
			typ:    "",
			x:      x,
			y:      y,
			desc:   "miam",
			intent: MIAM,
		}
	} else if grid[y][x].my {
		return false, cmd{} //todo stop collide for a pellet
	} else if grid[y][x].his {

		win := fight3(rpc, grid[y][x].hisPac.rpc)
		fmt.Fprintln(os.Stderr, id, "fight", rpc, "vs", grid[y][x].hisPac.rpc, "score", win)
		if win == 1 {
			return true, cmd{
				action: MOVE,
				typ:    "",
				x:      x,
				y:      y,
				desc:   "kill him",
				intent: KILL,
			}
		} else {
			nemesis := findNemesis(grid[y][x].hisPac.rpc)
			return true, cmd{
				action: SWITCH,
				typ:    nemesis,
				x:      x,
				y:      y,
				desc:   "iam batman",
				intent: SURVIVE,
			}
		}
	} else {
		return false, cmd{}
	}

}

func findNemesis(rpc string) string {

	if rpc == ROCK {
		return PAPER
	} else if rpc == PAPER {
		return SCISSORS
	} else {
		return ROCK
	}
}

func findPacCoord(id int, grid [][]cel) (int, int, pac) {
	for y, row := range grid {
		for x, cel := range row {
			if cel.my && cel.myPac.id == id {
				return x, y, cel.myPac
			}
		}
	}
	fmt.Fprint(os.Stderr, "pac not found")
	return 0, 0, pac{}

}

func rundInit(scanner *bufio.Scanner, grid [][]cel) []int {
	var myScore, opponentScore int

	var myPacIds []int

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
	var visiblePacCount int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &visiblePacCount)

	rinitGrid(grid)

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
			myPacIds = append(myPacIds, pacId)
			markAsMe(grid, y, x, pacId, typeId, speedTurnsLeft, abilityCooldown)
		} else {
			markAsHis(grid, y, x, pacId, typeId, speedTurnsLeft, abilityCooldown)
		}
	}

	var visiblePelletCount int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &visiblePelletCount)

	for i := 0; i < visiblePelletCount; i++ {
		var x, y, value int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &value)
		grid[y][x].value = value
	}

	return myPacIds
}

func rinitGrid(grid [][]cel) {

	for y, cels := range grid {
		for x := range cels {
			grid[y][x].my = false
			grid[y][x].myPac = pac{}
			grid[y][x].his = false
			grid[y][x].hisPac = pac{}
			grid[y][x].value = -1

		}
	}
}

func markAsHis(grid [][]cel, y int, x int, pacId int, typeId string, left int, cooldown int) {
	grid[y][x].his = true
	grid[y][x].hisPac = pac{pacId, typeId, left, cooldown}
}

func markAsMe(grid [][]cel, y int, x int, pacId int, typeId string, left int, cooldown int) {
	grid[y][x].my = true
	grid[y][x].myPac = pac{pacId, typeId, left,cooldown}
}

func initial(scanner *bufio.Scanner) [][]cel {
	var width, height int

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	var grid = make([][]cel, height)

	for y := 0; y < height; y++ {

		grid[y] = make([]cel, width)

		scanner.Scan()
		row := scanner.Text()
		for x := 0; x < width; x++ {

			if row[x:x+1] == " " {
				grid[y][x].floor = true
			} else {
				grid[y][x].floor = false
			}
			grid[y][x].my = false
			grid[y][x].his = false
		}
	}

	return grid
}

//1 win 0 nul -1 loose
func fight3(rpc1 string, rpc2 string) int {

	if rpc1 == ROCK && rpc2 == SCISSORS {
		return 1
	} else if rpc2 == ROCK && rpc1 == SCISSORS {
		return -1
	} else if rpc1 == SCISSORS && rpc2 == PAPER {
		return 1
	} else if rpc2 == SCISSORS && rpc1 == PAPER {
		return -1
	} else if rpc1 == PAPER && rpc2 == ROCK {
		return 1
	} else if rpc2 == PAPER && rpc1 == ROCK {
		return -1
	} else {
		return 0
	}
}

func displayGrid(grid [][]cel) {
	for _, row := range grid {
		for _, cel := range row {
			if cel.floor == false {
				fmt.Fprint(os.Stderr, "#")
			} else if cel.my {
				fmt.Fprint(os.Stderr, cel.myPac.id)
			} else if cel.his {
				if cel.hisPac.id == 0 {
					fmt.Fprint(os.Stderr, "A")
				} else if cel.hisPac.id == 1 {
					fmt.Fprint(os.Stderr, "B")
				} else if cel.hisPac.id == 2 {
					fmt.Fprint(os.Stderr, "C")
				} else if cel.hisPac.id == 3 {
					fmt.Fprint(os.Stderr, "D")
				} else if cel.hisPac.id == 4 {
					fmt.Fprint(os.Stderr, "E")
				} else {
					fmt.Fprint(os.Stderr, "?")
				}
			} else if cel.value == 1 {
				fmt.Fprint(os.Stderr, ".")
			} else if cel.value == 10 {
				fmt.Fprint(os.Stderr, "+")
			} else {
				fmt.Fprint(os.Stderr, " ")
			}
		}
		fmt.Fprintln(os.Stderr, "/")
	}
}

//fmt.Fprint(os.Stderr, "*")

//
//func randBullet(grid [][]cel) (int, int) {
//
//	for true {
//		x := rand.Intn(len(grid[0]))
//		y := rand.Intn(len(grid))
//		//fmt.Fprintln(os.Stderr, x, y, grid[y][x].value, grid[y][x].t)
//		if grid[y][x].value >= 1 && grid[y][x].t == false {
//			return x, y
//		}
//	}
//
//	return 0, 0
//}
