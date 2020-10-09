package main

import (
	"fmt"
	"os"
	strconv "strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type opeT struct {
	v1  string
	v2  string
	typ string
}

const nan = 1111111

func main() {
	var N int
	const nan = 1111111
	fmt.Scan(&N)
	fmt.Fprintln(os.Stderr, N)

	values := make([]int, N)

	for i, _ := range values {
		values[i] = nan
	}

	ope := make([]opeT, N)

	for i := 0; i < N; i++ {
		var operation, arg1, arg2 string
		fmt.Scan(&operation, &arg1, &arg2)
		fmt.Fprintln(os.Stderr, "op ", operation, "a1 ", arg1, "a2 ", arg2)

		ope[i] = opeT{arg1, arg2, operation}

	}
	//for i := 0; i < N; i++ {
	//
	//	// fmt.Fprintln(os.Stderr, "Debug messages...")
	//	fmt.Println("1")// Write answer to stdout
	//}

	for !valuesOk(values) {
		for i, op := range ope {
			compute(op, values, i)
		}
	}

	fmt.Fprintln(os.Stderr, values)
	fmt.Fprintln(os.Stderr, ope)

	for _, value := range values {
		fmt.Println(value)
	}
}

func compute(op opeT, values []int, i int) {

	var v1, v2 int

	if op.v1[:1] == "$" {
		fmt.Fprintln(os.Stderr,"need ref", op.v1[:1])
		atoi, _ := strconv.Atoi(op.v1[1:])
		i := values[atoi]
		if i == nan {
			fmt.Fprintln(os.Stderr,"need ref", op.v1[:1], "not found")
			return
		} else {
			v1 = i
			fmt.Fprintln(os.Stderr,"need ref", op.v1[:1], "found", v1)

		}
	} else {
		v1, _ = strconv.Atoi(op.v1)
		fmt.Fprintln(os.Stderr,"value found", op.v1)

	}

	if op.typ == "VALUE" {
		values[i] = v1
		fmt.Fprintln(os.Stderr,"Set", i, "to", op.v1[1:])
		return
	}

	if op.v2[:1] == "$" {
		fmt.Fprintln(os.Stderr,"need ref", op.v2[:1])
		atoi, _ := strconv.Atoi(op.v2[1:])
		i := values[atoi]
		if i == nan {
			fmt.Fprintln(os.Stderr,"need ref", op.v2[:1], "not found")
			return
		} else {
			v2 = i
			fmt.Fprintln(os.Stderr,"need ref", op.v2[:1], "found", v2)
		}
	} else {
		v2, _ = strconv.Atoi(op.v2)
		fmt.Fprintln(os.Stderr,"value found", v2)

	}

	if op.typ == "MULT" {
		values[i] = v1 * v2
		fmt.Fprintln(os.Stderr,"MULT", i, v1, v2)

	} else if op.typ == "SUB" {
		values[i] = v1 - v2
		fmt.Fprintln(os.Stderr,"SUB", i, v1, v2)

	} else if op.typ == "ADD" {
		values[i] = v1 + v2
		fmt.Fprintln(os.Stderr,"SUB", i, v1, v2)

	} else {
		fmt.Fprintln(os.Stderr,"Unknown operation") // Write answer to stdout
	}

}

func valuesOk(values []int) bool {
	for _, value := range values {
		if value == nan {
			return false
		}
	}
	return true

}
