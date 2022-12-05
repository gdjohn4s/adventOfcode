package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "input.prod"

func main() {
	puzzle1()
}

/* Day 4 Puzzle 1 */
func puzzle1() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var stronzi int = 0
	for sc.Scan() {
		var first, lastFirst, second, lastSecond int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &first, &lastFirst, &second, &lastSecond)
		if second >= first && lastSecond <= lastFirst || first >= second && lastFirst <= lastSecond {
			stronzi += 1
		}

		fmt.Println(stronzi)
	}
}

/* Day 4 Puzzle 2 */
func puzzle2() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var stronzi int = 0
	for sc.Scan() {
		var first, lastFirst, second, lastSecond int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &first, &lastFirst, &second, &lastSecond)
		if second <= lastFirst && lastSecond >= first || first <= second && lastFirst >= lastSecond {
			stronzi += 1
		}
	}
	fmt.Println(stronzi)
}

