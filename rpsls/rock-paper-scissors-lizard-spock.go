package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type pl struct {
	id int
	si string
}

func main() {
	var N int
	fmt.Scan(&N)

	var player = make([][]pl, N/2 + 1)
	player[0] = make([]pl, N)

	for i := 0; i < N; i++ {
		var NUMPLAYER int
		var SIGNPLAYER string
		fmt.Scan(&NUMPLAYER, &SIGNPLAYER)

		player[0][i] = pl{NUMPLAYER, SIGNPLAYER}
	}

	fmt.Fprintln(os.Stderr, player)

	var r = 0

	for len(player[r]) != 1 {
		r++
		player[r] = make([]pl, N/2)
		for i := 0; i < N/2; i++ {
			player[r][i] = fight(player[r-1][i*2], player[r-1][(i*2)+1])
		}
		N = N / 2
		fmt.Fprintln(os.Stderr, player)

	}

	fmt.Fprintln(os.Stderr, "last player ?", player)
	winner := player[r][0].id
	fmt.Println(winner) // Write answer to stdout
	fmt.Println(getOps(player, winner))
}

func getOps(players [][]pl, winner int) string {

	var loosers = ""
	for _, player := range players {
		for i := 0; i < len(player)/2; i++ {
			if player[i*2].id == winner {
				loosers += strconv.Itoa(player[i*2+1].id) + " "
			} else if player[i*2+1].id == winner {
				loosers += strconv.Itoa(player[i*2].id) + " "
			}
		}
	}

	return loosers[:len(loosers)-1]

}

func fight(p1 pl, p2 pl) pl {

	var sign1 = p1.si
	var sign2 = p2.si
	if sign1 == sign2 {
		if p1.id > p2.id {
			return p2
		} else {
			return p1
		}
	} else if sign1 == "C" && sign2 == "P" {
		return p1
	} else if sign2 == "C" && sign1 == "P" {
		return p2
	} else if sign1 == "P" && sign2 == "R" {
		return p1
	} else if sign2 == "P" && sign1 == "R" {
		return p2
	} else if sign1 == "R" && sign2 == "L" {
		return p1
	} else if sign2 == "R" && sign1 == "L" {
		return p2
	} else if sign1 == "L" && sign2 == "S" {
		return p1
	} else if sign2 == "L" && sign1 == "S" {
		return p2
	} else if sign1 == "S" && sign2 == "C" {
		return p1
	} else if sign2 == "S" && sign1 == "C" {
		return p2
	} else if sign1 == "C" && sign2 == "L" {
		return p1
	} else if sign2 == "C" && sign1 == "L" {
		return p2
	} else if sign1 == "L" && sign2 == "P" {
		return p1
	} else if sign2 == "L" && sign1 == "P" {
		return p2
	} else if sign1 == "P" && sign2 == "S" {
		return p1
	} else if sign2 == "P" && sign1 == "S" {
		return p2
	} else if sign1 == "S" && sign2 == "R" {
		return p1
	} else if sign2 == "S" && sign1 == "R" {
		return p2
	} else if sign1 == "R" && sign2 == "C" {
		return p1
	} else if sign2 == "R" && sign1 == "C" {
		return p2
	}
	return p1
}
