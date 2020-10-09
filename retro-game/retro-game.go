package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type thing struct {
	fifth int
	sixth int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var firstInitInput int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &firstInitInput)

	var secondInitInput int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &secondInitInput)

	var thirdInitInput int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &thirdInitInput)
	turn := 0
	var currentThings, previourThings [5]thing

	for {

		turn++
		fmt.Fprintln(os.Stderr, "(", firstInitInput, ",", secondInitInput, ")", thirdInitInput)

		scanner.Scan()
		firstInput := scanner.Text()
		fmt.Fprint(os.Stderr, firstInput)

		scanner.Scan()
		secondInput := scanner.Text()
		fmt.Fprint(os.Stderr, secondInput)

		scanner.Scan()
		thirdInput := scanner.Text()
		fmt.Fprint(os.Stderr, thirdInput)

		scanner.Scan()
		fourthInput := scanner.Text()
		fmt.Fprint(os.Stderr, fourthInput)

		fmt.Fprintln(os.Stderr)

		for i := 0; i < thirdInitInput; i++ {
			var fifthInput, sixthInput int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &fifthInput, &sixthInput)
			fmt.Fprintln(os.Stderr, "(", fifthInput, ",", sixthInput, ")")

			currentThings[i] = thing{fifth: fifthInput, sixth: sixthInput}
		}
		fmt.Fprintln(os.Stderr)
		//do the diff
		if turn != 0 {
			for i := 0; i < len(currentThings); i++ {
				fmt.Fprintln(os.Stderr, "(", currentThings[i].fifth - previourThings[i].fifth, ",", currentThings[i].sixth - previourThings[i].sixth, ")")
			}
		}

		previourThings = currentThings

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("A") // Write action to stdout
	}
}
